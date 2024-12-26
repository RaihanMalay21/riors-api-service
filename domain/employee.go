package domain

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Id              uint   `gorm:"primaryKey" json:"id"`
	Name            string `gorm:"varchar(100);not null" json:"name"`
	Email           string `gorm:"varchar(100);not null;unique" json:"email"`
	Whatsapp        string `gorm:"varchar(20);not null;unique" json:"whatsapp"`
	Password        string `gorm:"varchar(200);not null" json:"-"`
	DateOfBirth     string `gorm:"varchar(20);not null" json:"dateOfBirth"`
	Gender          string `gorm:"type:new_employee_gender;not null" json:"gender"`
	Image           string `gorm:"varchar(200);not null" json:"image"`
	Address         string `gorm:"vacrhar(200);not null" json:"address"`
	Position        string `gorm:"type:new_employee_position;not null" json:"position"`
	EmployementType string `gorm:"type:new_employee_type;not null" json:"employementType"`
}
