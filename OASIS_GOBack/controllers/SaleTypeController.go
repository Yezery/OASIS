package controllers

import (
	"net/http"

	"example.com/m/models"
	"example.com/m/repositories"
	"example.com/m/utils"
	"github.com/gin-gonic/gin"
)

type SaleTypeController struct{}

var TypeModel models.SaleType

// 查询上架表中的全部记录
func (STC *SaleTypeController) GetTypeList(c *gin.Context) {
	var Type []models.SaleType
	repositories.MakeDbModel[models.SaleType](&TypeModel, c)
	db := repositories.GetDb(c)
	result := db.Find(&Type)
	if result.Error != nil {
		utils.SendResponse(c.Writer, http.StatusBadGateway, result.Error)
		panic(result.Error)
	}
	utils.SendResponse(c.Writer, http.StatusOK, Type)
}
