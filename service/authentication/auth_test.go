package authentication

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	response := make(map[string]interface{})
	authentication := &AuthenticationService{}

	cookie, status := authentication.LoginUser("rcndonvpifvf", "jgufvoufh", response)

	fmt.Println(cookie, status)
}
