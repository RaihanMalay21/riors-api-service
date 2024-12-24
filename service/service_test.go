package service

import (
	"fmt"
	"testing"
)

func TestService(t *testing.T) {
	value := GenerateRandomNumber()
	fmt.Println(value)
}