package controllers

import (
	"fmt"
	"net/http"
	"time"

	"example.com/m/models"
	"example.com/m/repositories"
	"example.com/m/utils"
	"github.com/gin-gonic/gin"
)

type TransactionController struct{}

func (TC *TransactionController) ScheduleDailySummary(c *gin.Context) {
	db := repositories.GetDb(c)
	type SaleSummary struct {
		Date          time.Time
		TotalTurnover float64
	}

	var summaries []SaleSummary
	db.
		Table("transactions").
		Select("DATE(created_at) AS date, SUM(turnover) AS total_turnover").
		Where("created_at >= CURDATE() - INTERVAL WEEKDAY(CURDATE()) DAY AND created_at < CURDATE() + INTERVAL 1 DAY").
		Group("DATE(created_at)").
		Order("DATE(created_at)").
		Scan(&summaries)
	utils.SendResponse(c.Writer, http.StatusOK, summaries)
}
func (TC *TransactionController) MakeNewTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		fmt.Println(transaction)
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}
	repositories.MakeDbModel[models.Transaction](&models.Transaction{}, c)
	db := repositories.GetDb(c)
	db.Create(&transaction)
	utils.SendResponse(c.Writer, http.StatusOK, transaction)
}
