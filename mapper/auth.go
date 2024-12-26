package mapper

import (
	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
)

func EmployeeDTOToEmployeeDomain(data *dto.Employee) domain.Employee {
	return domain.Employee{
		Name: data.Name,
		Email: data.Email,
		Whatsapp: data.Whatsapp,
		Password: data.Password,
		DateOfBirth: data.DateOfBirth,
		Gender: data.Gender,
		Image: data.Image,
		Address: data.Address,
		Position: data.Position,
		EmployementType: data.EmployementType,
	}
}