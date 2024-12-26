package dto

type Employee struct {
	Id              uint   `json:"id"`
	Name            string `validate:"min=6,max=50,required" json:"name"`
	Email           string `validate:"uniqueEmailEmployee,required,email,max=100" json:"email"`
	Whatsapp        string `validate:"required,uniqueWAEmployee,whatsapp" json:"whatsapp"`
	Password        string `validate:"required"`
	Position        string `validate:"required,positionEmployee" json:"position"`
	EmployementType string `validate:"required,employementType" json:"employeeType"`
	DateOfBirth     string `validate:"required,date_format" json:"dateOfBirth"`
	Gender          string `validate:"required,genderEmployee" json:"gender"`
	Address         string `validate:"required" json:"address"`
	Image           string `validate:"required" json:"image"`
	FileSize        uint   `validate:"required,maxSizeFile" json:"-"`
	Ext             string `validate:"required" json:"-"`
	ImageType       string `validate:"required" json:"-"`
}

type SetEmployeePassword struct {
	Password string `validate:"required"`
}
