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
	Email string `json:"email"`
	Password string `json:"password"`
}

type Verification struct {
	Code int `json:"code"`
}
