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

type NFTOwnerListController struct{}

// 创建一条新记录
func (NFTLC *NFTOwnerListController) CreateNFTInf(c *gin.Context) {
	var nftOwnerList models.NFTOwnerList // 实例化空对象
	repositories.MakeDbModel[models.NFTOwnerList](&nftOwnerList, c)
	if err := c.BindJSON(&nftOwnerList); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	fmt.Println(&nftOwnerList)
	db := repositories.GetDb(c)
	result := db.Create(&nftOwnerList)
	if result.Error != nil {
		utils.SendResponse(c.Writer, http.StatusInternalServerError, result.Error)
		panic(result.Error)
	}
	utils.SendResponse(c.Writer, http.StatusOK, "success")
}

// 查询属于个人的上架的NFT
func (NFTLC *NFTOwnerListController) GetOwnerUpSaleNFTs(c *gin.Context) {
	var nftOwnerList models.NFTOwnerList
	if err := c.BindJSON(&nftOwnerList); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	fmt.Println(nftOwnerList)
	var results dto.NFTOwnerListDTO
	db := repositories.GetDb(c)
	query := `
	SELECT 
	nol.id,
	nol.seriesName,
	nol.nftName,
	nol.isActive,
	nol.price,
	nol.symbol,
	nol.tokenId,
	nol.ownerAddress,
	nol.nftAddress,
	nol.ipfsPath,
	nol.maxmums,
	nol.description,
	s.sale_id
	FROM nft_owner_lists nol
	INNER JOIN sales s ON s.nft_owner_list_Id = nol.id 
	WHERE s.nft_owner_list_Id = ? `
	db.Raw(query, nftOwnerList.Id).Scan(&results)
	utils.SendResponse(c.Writer, http.StatusOK, results)
}

// query := `
//
//	SELECT nol.*, s.*
//	FROM nft_owner_lists nol
//	INNER JOIN sales s ON s.sale_id = nol.id
//	WHERE nol.ownerAddress = ?`
//
// 查询属于个人的NFT
func (NFTLC *NFTOwnerListController) GetOwnerNFTs(c *gin.Context) {
	var nftOwnerList models.NFTOwnerList
	if err := c.BindJSON(&nftOwnerList); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	var results []models.NFTOwnerList
	db := repositories.GetDb(c)
	query := `
		SELECT nol.*
		FROM nft_owner_lists nol
		WHERE nol.currentOwner = ?`
	db.Raw(query, nftOwnerList.OwnerAddress).Scan(&results)
	for _, result := range results {
		fmt.Println(result)
	}
	utils.SendResponse(c.Writer, http.StatusOK, results)
}

// 查询属于个人的NFT by nftAddress
func (NFTLC *NFTOwnerListController) GetOwnerNFTsByAddress(c *gin.Context) {
	var nftOwnerList models.NFTOwnerList
	if err := c.BindJSON(&nftOwnerList); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	var results []models.NFTOwnerList
	db := repositories.GetDb(c)
	query := `
	SELECT nol.*
	FROM nft_owner_lists nol
	WHERE nol.nftAddress = ? AND nol.ownerAddress = ?`
	db.Raw(query, nftOwnerList.NFTAddress, nftOwnerList.OwnerAddress).Scan(&results)
	utils.SendResponse(c.Writer, http.StatusOK, results)
}

// 修改一条上架情况
func (NFTLC *NFTOwnerListController) UpdateNFTOwnerList(c *gin.Context) {
	var newNFTOwnerList models.NFTOwnerList

	if err := c.ShouldBindJSON(&newNFTOwnerList); err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}

	db := repositories.GetDb(c)

	// 执行更新操作
	result := db.Model(&models.NFTOwnerList{}).Where("id = ?", newNFTOwnerList.Id).Updates(map[string]interface{}{
		"isActive": newNFTOwnerList.IsActive, "price": newNFTOwnerList.Price,
	})
	fmt.Println(result)

	if result.Error != nil {
		utils.SendResponse(c.Writer, http.StatusInternalServerError, result.Error)
		panic(result.Error)
	}

	utils.SendResponse(c.Writer, http.StatusOK, newNFTOwnerList)
}

// 购买后修改
func (NFTLC *NFTOwnerListController) UpdateNFTOwnerListAfterBuy(c *gin.Context) {
	var newNFTOwnerList models.NFTOwnerList

	if err := c.ShouldBindJSON(&newNFTOwnerList); err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}

	db := repositories.GetDb(c)

	// 执行更新操作
	result := db.Model(&models.NFTOwnerList{}).Where(map[string]interface{}{"tokenId": newNFTOwnerList.TokenId, "nftAddress": newNFTOwnerList.NFTAddress}).Updates(map[string]interface{}{
		"isActive": newNFTOwnerList.IsActive, "price": newNFTOwnerList.Price, "currentOwner": newNFTOwnerList.CurrentOwner,
	})
	fmt.Println(result)

	if result.Error != nil {
		utils.SendResponse(c.Writer, http.StatusInternalServerError, result.Error)
		panic(result.Error)
	}

	utils.SendResponse(c.Writer, http.StatusOK, newNFTOwnerList)
}

type NFTSearchCriteria struct {
	Key string `json:"key"`
	// IsActive   bool   `json:"isActive"`
	// MinPrice   string `json:"minPrice"`
	// MaxPrice   string `json:"maxPrice"`
	// MinMaxmums string `json:"minMaxmums"`
	// MaxMaxmums string `json:"maxMaxmums"`
}

func (NFTLC *NFTOwnerListController) Search(c *gin.Context) {
	var results []dto.NFTOwnerListDTO
	var criteria NFTSearchCriteria
	if err := c.ShouldBindJSON(&criteria); err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}
	keyword := "%" + criteria.Key + "%"
	//精确
	repositories.GetDb(c).Model(&models.NFTOwnerList{}).
		Select("*").
		Joins("INNER JOIN sales s ON s.nft_owner_list_Id = nft_owner_lists.id").
		Where("seriesName = ? OR nftName = ? OR symbol = ? OR description = ?", keyword, keyword, keyword, keyword).
		Find(&results)
	if len(results) == 0 {
		//模糊
		repositories.GetDb(c).Model(&models.NFTOwnerList{}).
			Select("*").
			Joins("INNER JOIN sales s ON s.nft_owner_list_Id = nft_owner_lists.id").
			Where("seriesName LIKE ? OR nftName LIKE ? OR symbol LIKE ? OR description LIKE ?", keyword, keyword, keyword, keyword).
			Find(&results)
	}

	utils.SendResponse(c.Writer, http.StatusOK, results)
}
