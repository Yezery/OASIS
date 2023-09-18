package dto

import "example.com/m/models"

type NFTOwnerListDTO struct {
	*models.NFTOwnerList
	*models.Sale
}
