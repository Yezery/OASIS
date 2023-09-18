package controllers

import (
	"fmt"
	"net/http"

	"example.com/m/models"
	"example.com/m/repositories"
	"example.com/m/utils"
	"github.com/gin-gonic/gin"
)

type SaleController struct{}

var UpdateParam repositories.UpdateParam

// 查询上架表中的全部记录
func (SC *SaleController) GetSaleList(c *gin.Context) {
	var Sales []models.Sale
	db := repositories.GetDb(c)
	result := db.Find(&Sales)
	fmt.Println(result)
	if result.Error != nil {
		utils.SendResponse(c.Writer, http.StatusInternalServerError, result.Error)
		panic(result.Error)
	}
	utils.SendResponse(c.Writer, http.StatusOK, Sales)
}

// 创建一条新记录
func (SC *SaleController) CreateSale(c *gin.Context) {
	var sale models.Sale
	repositories.MakeDbModel[models.Sale](&models.Sale{}, c)
	if err := c.BindJSON(&sale); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	fmt.Println(&sale)
	db := repositories.GetDb(c)
	// 向 Sale 表插入记录
	result := db.Create(&sale)
	if result.Error != nil {
		utils.SendResponse(c.Writer, http.StatusInternalServerError, result.Error)
		panic(result.Error)
	}

	utils.SendResponse(c.Writer, http.StatusOK, result)
}

// 购买后删除相对应纪律
func (SC *SaleController) DeleteSale(c *gin.Context) {
	var NFT models.NFTOwnerList
	if err := c.BindJSON(&NFT); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	db := repositories.GetDb(c)
	sale := models.Sale{
		SaleId:         0,
		NFTOwnerListId: NFT.Id,
	}
	// 执行删除操作
	// 执行更新操作
	update := db.Model(&models.NFTOwnerList{}).Where("id = ?", sale.NFTOwnerListId).Updates(map[string]interface{}{
		"isActive": false,
	})
	delete := db.Delete(&sale, "nft_owner_list_id = ?", sale.NFTOwnerListId)
	if delete.Error != nil || update.Error != nil {
		panic(delete.Error)
	}
	utils.SendResponse(c.Writer, http.StatusOK, delete)
}
