package mapper

import (
	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
)

func CartDTOTODomain(data *dto.Cart) domain.Cart {
	return domain.Cart{
		UserId:           data.UserId,
		ProductVariantId: data.ProductVariantId,
		AmountItem:       data.AmountItem,
		AmountPrice:      data.AmountPrice,
	}
}

func CartDomainTODTO(data *domain.Cart) dto.Cart {
	return dto.Cart{
		ProductId:        data.ProductVariant.ProductId,
		UserId:           data.UserId,
		Color:            data.ProductVariant.Color,
		Size:             data.ProductVariant.Size,
		AmountItem:       data.AmountItem,
		AmountPrice:      data.AmountPrice,
		ProductVariantId: data.ProductVariantId,
		Image:            data.ProductVariant.Image,
		ProductName:      data.ProductVariant.Product.ProductName,
		Price:            data.ProductVariant.Product.Price,
	}
}

func ArrayCartDomainTODTO(carts *[]domain.Cart) []dto.Cart {
	var cartsDTO = make([]dto.Cart, 0, len(*carts))

	for _, data := range *carts {
		cartsDTO = append(cartsDTO, CartDomainTODTO(&data))
	}

	return cartsDTO
}
