package mapper

import (
	"sync"

	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
)

func ProductDTOTODomain(data *dto.Product) domain.Product {
	return domain.Product{
		CategoryId: data.CategoryId,
		ProductName: data.ProductName,
		HargaBarang: data.HargaBarang,
		Type: data.Type,
		Image: data.Image,
	}
}

func ProductDomainTODTO(data *domain.Product) dto.Product {
	return dto.Product{
		Id: data.Id,
		CategoryId: data.CategoryId,
		ProductName: data.ProductName,
		HargaBarang: data.HargaBarang,
		Type: data.Type,
		Image: data.Image,
	}
}

func GetAllProductDomainTODTO(datas *[]domain.Product) *[]dto.Product {
	var dataDto = make([]dto.Product, len(*datas))
	group := &sync.WaitGroup{}

	for i, data := range *datas {
		group.Add(1)
		go func(i int, data domain.Product) {
			defer group.Done()
			dataDto[i] = ProductDomainTODTO(&data)
		}(i, data)
	} 

	group.Wait()
	return &dataDto
}