package authentication

type ResponseSuccess struct {
	Success string `json:"success"`
}

type ResponseErrorBadRequest struct {
	ErrorFields []map[string]string `json:"ErrorField"`
}

type ResponseErrorInternalServer struct {
	Error string `json:"error"`
}

type ResponsAuthorization struct {
	Message string `json:"message"`
}

type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Verification struct {
	Code int `json:"code"`
}

type SignupEmploye struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Whatsapp        string `json:"whatsapp"`
	Position        string `json:"position"`
	EmployementType string `json:"employementType"`
	Image           string `json:"image"`
	DateOfBirth     string `json:"dateOfBirth"`
	Gender          string `json:"gender"`
	Address         string `json:"address"`
}

type ChangePassword struct {
	PasswordBefore string `json:"passwordBefore"`
	Password       string `json:"password"`
}
