package models

type NFTOwnerList struct {
	Id           int64  `gorm:"primaryKey;autoIncrement;"`
	SeriesName   string `gorm:"column:seriesName" json:"seriesName"`
	NFTName      string `gorm:"column:nftName" json:"nftName"`
	IsActive     bool   `gorm:"column:isActive" json:"isActive"`
	Price        string `gorm:"column:price" json:"price"`
	Symbol       string `gorm:"column:symbol" json:"symbol"`
	TokenId      uint64 `gorm:"column:tokenId" json:"tokenId"`
	CurrentOwner string `gorm:"column:currentOwner" json:"currentowner"`
	OwnerAddress string `gorm:"column:ownerAddress" json:"ownerAddress"`
	NFTAddress   string `gorm:"column:nftAddress" json:"nftAddress"`
	IpfsPath     string `gorm:"column:ipfsPath" json:"ipfsPath"`
	Maxmums      uint64 `gorm:"column:maxmums" json:"maxmums"`
	Description  string `gorm:"column:description" json:"description"`
}
