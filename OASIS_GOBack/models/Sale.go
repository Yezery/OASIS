package models

type Sale struct {
	SaleId         int64         `gorm:"column:sale_id" json:"saleId"`
	NFTOwnerListId int64         `gorm:"column:nft_owner_list_Id"`
	NFTOwnerList   *NFTOwnerList `gorm:"foreignKey:NFTOwnerListId;references:Id" json:"NFTOwnerList"`
}
