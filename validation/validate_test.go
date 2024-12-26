package validation

import (
	"fmt"
	"testing"
)

func TestValidationFunc(t *testing.T) {
	// boll := isMinUppercase("vbfhbvkF")
	// boll := isMinNumber("cjdbckdj5")
	// boll := isMinUniCharacter("732739hdscvf")
	// boll := isDateFormat("2005-12-03")
	// boll := isWhatsapp("089524474969")
	boll := isUniqueWhatsappEmployee("089524474969")
	fmt.Println(boll)
}
