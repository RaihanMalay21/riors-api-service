package mapper

import (
	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
)

func CategoryDTOTODomain(data *dto.Category) domain.Category {
	return domain.Category{
		CategoryName: data.CategoryName,
	}
}

func CategoryDomainTODTO(data *domain.Category) dto.Category {
	return dto.Category{
		Id:           data.Id,
		CategoryName: data.CategoryName,
		CreatedAt:    data.CreatedAt,
	}
}

func GetAllCategoryDomainTODTO(datas *[]domain.Category) *[]dto.Category {
	var dataDto []dto.Category

	for _, data := range *datas {
		dto := CategoryDomainTODTO(&data)
		dataDto = append(dataDto, dto)
	}

	return &dataDto
}
