package dto

type ChangePassword struct {
	Id             uint   `validate:"required"`
	PasswordBefore string `validate:"required"`
	Password       string `validate:"required,minUppercase,minNumber,minUniCharacter,min=8"`
}

type ResetPassword struct {
	Token          string `validate:"requiredToken"`
	Password       string `validate:"required,minUppercase,minNumber,minUniCharacter,min=8"`
	PasswordRepeat string `validat:"required"`
}
