package controllers

import (
	"fmt"
	"net/http"

	"example.com/m/models"
	"example.com/m/models/dto"
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

// 条件查询
func (SC *SaleController) GetSaleListByContractAddress(c *gin.Context) {
	var vo models.NFTOwnerList
	if err := c.BindJSON(&vo); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	var results []dto.NFTOwnerListDTO
	repositories.GetDb(c).Model(&models.NFTOwnerList{}).
		Select("*").
		Joins("INNER JOIN sales s ON s.nft_owner_list_Id = nft_owner_lists.id").
		Where("nftAddress = ?", vo.NFTAddress).
		Find(&results)
	fmt.Println(results)
	utils.SendResponse(c.Writer, http.StatusOK, &results)
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
	var NFT dto.NFTOwnerListDTO
	if err := c.BindJSON(&NFT); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	db := repositories.GetDb(c)
	update := db.Model(&models.NFTOwnerList{}).Where(map[string]interface{}{"tokenId": NFT.TokenId, "nftAddress": NFT.NFTAddress}).Updates(map[string]interface{}{
		"isActive": false,
	})
	delete := db.Where("sales.sale_id = ?", NFT.SaleId).Delete(&models.Sale{})
	if delete.Error != nil || update.Error != nil {
		panic(delete.Error)
	}
	utils.SendResponse(c.Writer, http.StatusOK, "")
}
