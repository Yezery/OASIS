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
	nol.maximums,
	nol.description,
	s.sale_id
	FROM nft_owner_lists nol
	INNER JOIN sales s ON s.nft_owner_list_Id = nol.id 
	WHERE s.nft_owner_list_Id = ? `
	db.Raw(query, nftOwnerList.Id).Scan(&results)
	utils.SendResponse(c.Writer, http.StatusOK, results)
}

// 查询属于个人的NFT
func (NFTLC *NFTOwnerListController) GetOwnerNFTs(c *gin.Context) {
	var nftOwnerList models.NFTOwnerList
	if err := c.BindJSON(&nftOwnerList); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}

	fmt.Println(nftOwnerList)
	var results []models.NFTOwnerList
	db := repositories.GetDb(c)
	query := `
		SELECT nol.*
		FROM nft_owner_lists nol
		WHERE nol.currentOwner = ?`
	db.Raw(query, nftOwnerList.OwnerAddress).Scan(&results)
	utils.SendResponse(c.Writer, http.StatusOK, results)
}

// 查询属于个人的NFT by nftAddress and current_nftaddress
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
	WHERE nol.nftAddress = ? AND nol.currentOwner = ?`
	db.Raw(query, nftOwnerList.NFTAddress, nftOwnerList.CurrentOwner).Scan(&results)
	utils.SendResponse(c.Writer, http.StatusOK, results)
}

// 查询by nftAddress
func (NFTLC *NFTOwnerListController) GetSeriesByNFTAddress(c *gin.Context) {
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
	WHERE nol.nftAddress = ? `
	db.Raw(query, nftOwnerList.NFTAddress).Scan(&results)
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

// 查询
type NFTSearchCriteria struct {
	Key string `json:"key"`
}
type SearchStruct struct {
	Value string              `json:"value"`
	NFT   dto.NFTOwnerListDTO `json:"NFT"`
}

func (NFTLC *NFTOwnerListController) Search(c *gin.Context) {
	var results []dto.NFTOwnerListDTO
	var criteria NFTSearchCriteria
	if err := c.ShouldBindJSON(&criteria); err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}
	keyword := "%" + criteria.Key + "%"

	//模糊
	repositories.GetDb(c).Model(&models.NFTOwnerList{}).
		Select("*").
		Joins("LEFT JOIN sales s ON s.nft_owner_list_Id = nft_owner_lists.id").
		Where("seriesName LIKE ? OR nftName LIKE ? OR symbol LIKE ?", keyword, keyword, keyword).
		Find(&results)

	var searchStructList []SearchStruct
	for _, value := range results {
		var searchStruct = SearchStruct{
			value.SeriesName + " · " + value.NFTName + " · (" + value.Symbol + ")",
			value,
		}
		searchStructList = append(searchStructList, searchStruct)
	}
	fmt.Println(searchStructList)
	utils.SendResponse(c.Writer, http.StatusOK, searchStructList)
}

func (NFTLC *NFTOwnerListController) KeyDownSearch(c *gin.Context) {
	var results []dto.NFTOwnerListDTO
	var criteria NFTSearchCriteria
	if err := c.ShouldBindJSON(&criteria); err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}
	keyword := "%" + criteria.Key + "%"

	//模糊
	repositories.GetDb(c).Model(&models.NFTOwnerList{}).
		Select("*").
		Joins("LEFT JOIN sales s ON s.nft_owner_list_Id = nft_owner_lists.id").
		Where("seriesName LIKE ? OR nftName LIKE ? OR symbol LIKE ?", keyword, keyword, keyword).
		Find(&results)

	var searchStructList []SearchStruct
	for _, value := range results {
		var searchStruct = SearchStruct{
			value.SeriesName + " · " + value.NFTName + " · (" + value.Symbol + ")",
			value,
		}
		fmt.Println(searchStruct)
		searchStructList = append(searchStructList, searchStruct)
	}
	fmt.Println(searchStructList)
	utils.SendResponse(c.Writer, http.StatusOK, searchStructList)
}

type MainCriteria struct {
	Key      string      `json:"key"`
	Active   bool        `json:"isActive"`
	Type     string      `json:"type" `
	MaxPrice string      `json:"maxPrice"`
	MinPrice string      `json:"minPrice"`
	Maxmums  string      `json:"maxmums"`
	Minmums  string      `json:"minmums"`
	PageDto  dto.PageDto `json:"pageDto"`
}

type MainSearchResult struct {
	Data  []dto.NFTOwnerListDTO `json:"data"`
	Total int64                 `json:"total"`
}

func (NFTLC *NFTOwnerListController) MainSearch(c *gin.Context) {
	var (
		criteria         MainCriteria
		mainSearchResult MainSearchResult
	)
	if err := c.ShouldBindJSON(&criteria); err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}
	fmt.Println(criteria)
	query := repositories.GetDb(c).Model(&models.NFTOwnerList{}).Select("*").
		Joins("LEFT JOIN sales s ON s.nft_owner_list_Id = nft_owner_lists.id").
		Where("nft_owner_lists.seriesName LIKE ? OR nft_owner_lists.nftName LIKE ? OR nft_owner_lists.symbol LIKE ? ",
			"%"+criteria.Key+"%", "%"+criteria.Key+"%", "%"+criteria.Key+"%")
	query.Where("nft_owner_lists.isActive = ? ", criteria.Active)

	if criteria.Type == "1" {
		query = query.Where("nft_owner_lists.description = '3D'")
	} else if criteria.Type == "2" {
		query = query.Where("nft_owner_lists.description != '3D'")
	}

	if criteria.MinPrice != "" && criteria.MaxPrice != "" {
		query = query.Where("(nft_owner_lists.price BETWEEN ? AND ? )", criteria.MinPrice, criteria.MaxPrice)
	}

	if criteria.Minmums != "" && criteria.Maxmums != "" {
		query = query.Where("(nft_owner_lists.maxmums BETWEEN ? AND ? )", criteria.Minmums, criteria.Maxmums)
	}
	pageReq := &dto.PageDto{
		Page:     criteria.PageDto.Page,
		PageSize: criteria.PageDto.PageSize,
	}

	err := query.Scopes(utils.Paginate(pageReq)).Find(&mainSearchResult.Data).Count(&mainSearchResult.Total).Error
	if err != nil {
		fmt.Println(err)
		utils.SendResponse(c.Writer, http.StatusBadGateway, mainSearchResult)
	}
	utils.SendResponse(c.Writer, http.StatusOK, mainSearchResult)
}
