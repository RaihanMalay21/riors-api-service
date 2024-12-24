package validation

import (
	"fmt"
	"testing"
)

func TestValidationFunc(t *testing.T) {
	// boll := isMinUppercase("vbfhbvkF")
	// boll := isMinNumber("cjdbckdj5")
	boll := isMinUniCharacter("732739hdscvf")
	fmt.Println(boll)
}