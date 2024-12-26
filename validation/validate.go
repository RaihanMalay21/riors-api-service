package validation

import (
	"regexp"
	"time"
	"unicode"

	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/RaihanMalay21/api-service-riors/domain"

	"errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterCustomValidationsProduct(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterValidation("uniqueProduct", func(fl validator.FieldLevel) bool {
		product := fl.Field().String()
		return isUniqueProduct(product)
	})

	validate.RegisterTranslation("uniqueProduct", trans, func(ut ut.Translator) error {
		return ut.Add("uniqueProduct", "Product dengan nama tersebut sudah tersedia", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("uniqueProduct", fe.Field())
		return t
	})

	validate.RegisterValidation("maxSizeFile", func(fl validator.FieldLevel) bool {
		sizeFile := fl.Field().Uint()
		return isMaxSizeFile(uint(sizeFile))
	})

	validate.RegisterTranslation("maxSizeFile", trans, func(ut ut.Translator) error {
		return ut.Add("maxSizeFile", "Ukuran Image Terlalu Besar Maksimal Size 2mb", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("maxSizeFile", fe.Field())
		return t
	})

	validate.RegisterValidation("minUppercase", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isMinUppercase(content)
	})

	validate.RegisterTranslation("minUppercase", trans, func(ut ut.Translator) error {
		return ut.Add("minUppercase", "harus mengandung minimal satu huruf kapital.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("minUppercase", fe.Field())
		return t
	})

	validate.RegisterValidation("minNumber", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isMinNumber(content)
	})

	validate.RegisterTranslation("minNumber", trans, func(ut ut.Translator) error {
		return ut.Add("minNumber", "harus mengandung minimal satu number.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("minNumber", fe.Field())
		return t
	})

	validate.RegisterValidation("minUniCharacter", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isMinUniCharacter(content)
	})

	validate.RegisterTranslation("minUniCharacter", trans, func(ut ut.Translator) error {
		return ut.Add("minUniCharacter", "harus mengandung minimal satu spesial karakter.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("minUniCharacter", fe.Field())
		return t
	})

	validate.RegisterValidation("uniqueEmailUser", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isUniqueEmailUser(content)
	})

	validate.RegisterTranslation("uniqueEmailUser", trans, func(ut ut.Translator) error {
		return ut.Add("uniqueEmailUser", "Email sudah terdaftar", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("uniqueEmailUser", fe.Field())
		return t
	})

	validate.RegisterValidation("uniqueEmailEmployee", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isUniqueEmailEmployee(content)
	})

	validate.RegisterTranslation("uniqueEmailEmployee", trans, func(ut ut.Translator) error {
		return ut.Add("uniqueEmailEmployee", "Email sudah terdaftar", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("uniqueEmailEmployee", fe.Field())
		return t
	})

	validate.RegisterValidation("uniqueWAEmployee", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isUniqueWhatsappEmployee(content)
	})

	validate.RegisterTranslation("uniqueWAEmployee", trans, func(ut ut.Translator) error {
		return ut.Add("uniqueWAEmployee", "Whatsapp sudah terdaftar", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("uniqueWAEmployee", fe.Field())
		return t
	})

	validate.RegisterValidation("uniqueWAUser", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isUniqueWhatshappUser(content)
	})

	validate.RegisterTranslation("uniqueWAUser", trans, func(ut ut.Translator) error {
		return ut.Add("uniqueWAUser", "whatsapp sudah terdaftar", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("uniqueWAUser", fe.Field())
		return t
	})

	validate.RegisterValidation("whatsapp", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isWhatsapp(content)
	})

	validate.RegisterTranslation("whatsapp", trans, func(ut ut.Translator) error {
		return ut.Add("whatsapp", "Format harus berupa nomor whatsapp", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("whatsapp", fe.Field())
		return t
	})

	validate.RegisterValidation("date_format", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isDateFormat(content)
	})

	validate.RegisterTranslation("date_format", trans, func(ut ut.Translator) error {
		return ut.Add("date_format", "Format harus berupa tahun bulan tanggal", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("date_format", fe.Field())
		return t
	})

	validate.RegisterValidation("genderEmployee", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isGenderEmployee(content)
	})

	validate.RegisterTranslation("genderEmployee", trans, func(ut ut.Translator) error {
		return ut.Add("genderEmployee", "Gender harus berupa man atau woman", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("genderEmployee", fe.Field())
		return t
	})

	validate.RegisterValidation("employementType", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isEmployementType(content)
	})

	validate.RegisterTranslation("employementType", trans, func(ut ut.Translator) error {
		return ut.Add("employementType", "Jenis employee harus berupa tetap, kontrak, freelance", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("employementType", fe.Field())
		return t
	})

	validate.RegisterValidation("positionEmployee", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		return isPositionEmployee(content)
	})

	validate.RegisterTranslation("positionEmployee", trans, func(ut ut.Translator) error {
		return ut.Add("positionEmployee", "posisi employee harus berupa tetap, kontrak, freelance", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("positionEmployee", fe.Field())
		return t
	})

	// validate.RegisterValidation("minUppercase", func(fl validator.FieldLevel) bool {
	// 	sizeFile := fl.Field.string()
	// 	return
	// })

	// validate.RegisterValidation("typeExt", func(fl validator.FieldLevel) bool {
	// 	ext := fl.Field().String()
	// 	return isAllowedExtention(ext)
	// })

	// validate.RegisterTranslation("typeExt", trans, func(ut ut.Translator) error {
	// 	return ut.Add("typeExt", "Ektensi Image Harus Berupa .jpeg, .jpg, .png, .gif", true)
	// }, func(ut ut.Translator, fe validator.FieldError) string {
	// 	t, _ := ut.T("typeExt", fe.Field())
	// 	return t
	// })

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "Harus Di Isi", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		field := fe.StructField()
		t, _ := ut.T("required", field)
		return t
	})

	validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "Maksimal Panjang Hanya {0} Karakter", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		Param := fe.Param()
		t, _ := ut.T("max", Param)
		return t
	})

	validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "Minimal Panjang Hanya {0} Karakter", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		Param := fe.Param()
		t, _ := ut.T("min", Param)
		return t
	})

	validate.RegisterTranslation("number", trans, func(ut ut.Translator) error {
		return ut.Add("number", "Harus Berupa Angka", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		Param := fe.Field()
		t, _ := ut.T("number", Param)
		return t
	})

	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "harus berupa email", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		Param := fe.Field()
		t, _ := ut.T("email", Param)
		return t
	})

}

func isMaxSizeFile(Size uint) bool {
	return Size <= 2000000
}

func isUniqueProduct(productName string) bool {
	var product domain.Product
	if err := config.DB.Where("product_name = ?", productName).First(&product).Error; err != nil {
		return errors.Is(err, gorm.ErrRecordNotFound)
	}
	return false
}

func isMinUppercase(content string) bool {
	for _, char := range content {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func isMinNumber(content string) bool {
	for _, char := range content {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func isMinUniCharacter(content string) bool {
	for _, char := range content {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			return true
		}
	}
	return false
}

func isUniqueEmailUser(email string) bool {
	var user domain.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return errors.Is(err, gorm.ErrRecordNotFound)
	}
	return false
}

func isUniqueEmailEmployee(email string) bool {
	var user domain.Employee
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return errors.Is(err, gorm.ErrRecordNotFound)
	}
	return false
}

func isUniqueWhatshappUser(nomor string) bool {
	var user domain.User
	if err := config.DB.Where("whatsapp = ?", nomor).First(&user).Error; err != nil {
		return errors.Is(err, gorm.ErrRecordNotFound)
	}
	return false
}

func isUniqueWhatsappEmployee(nomor string) bool {
	var user domain.Employee
	if err := config.DB.Where("whatsapp = ?", nomor).First(&user).Error; err != nil {
		return errors.Is(err, gorm.ErrRecordNotFound)
	}
	return false
}


func isWhatsapp(nomor string) bool {
	regex := `^\+?[0-9]{8,15}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(nomor)
}

func isDateFormat(date string) bool {
	dateRegex := `^\d{4}-\d{2}-\d{2}$` // yyyy-mm-dd
	re := regexp.MustCompile(dateRegex)

	if !re.MatchString(date) {
		return false
	}

	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func isGenderEmployee(gender string) bool {
	if gender == "Man" || gender == "Woman" {
		return true
	}
	return false
}

func isEmployementType(content string) bool {
	if content == "Tetap" || content == "Kontrak" || content == "Freelance" {
		return true
	}
	return false
}

func isPositionEmployee(content string) bool {
	if content == "Staff" || content == "Owner" {
		return true
	}
	return false
}

// func isAllowedExtention(ext string) bool {
// 	if ext == ".jpeg" || ext == ".jpg" || ext == ".png" || ext == ".gif" {
// 		return true
// 	}

// 	return false
// }
