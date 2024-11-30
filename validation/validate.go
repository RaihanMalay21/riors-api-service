package validation

import (
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

	validate.RegisterTranslation("number", trans, func(ut ut.Translator) error {
		return ut.Add("number", "Harus Berupa Angka", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		Param := fe.Field()
		t, _ := ut.T("number", Param)
		return t
	})

}

func isMaxSizeFile(Size uint) bool {
	return Size <= 2000000
}

func isUniqueProduct(productName string) bool {
	var product domain.Product
	if err := config.DB.Where("product_name = ?", productName).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true
		}
		return false
	}
	return false
}

// func isAllowedExtention(ext string) bool {
// 	if ext == ".jpeg" || ext == ".jpg" || ext == ".png" || ext == ".gif" {
// 		return true
// 	}

// 	return false
// }
