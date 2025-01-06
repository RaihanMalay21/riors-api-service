package service

import (
	"fmt"
	"testing"

	"github.com/RaihanMalay21/api-service-riors/service/helper"
)

func TestService(t *testing.T) {
	var helper helper.HelperService
	value := helper.GenerateRandomNumber()
	fmt.Println(value)
}
