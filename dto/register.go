package dto

type RegisterUser struct {
	Email    string `validate:"required,email,uniqueEmailUser" json:"email"`
	Password string `validate:"required,minUppercase,minNumber,minUniCharacter,min=8" json:"password"`
	Code     int    `validate:"-"`
	Try      int
}


