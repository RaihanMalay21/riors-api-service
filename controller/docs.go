package controller

// response documentasi

type ResponseErrorBadRequest struct {
	ErrorFields []map[string]string
}

type ResponseErrorInternalServer struct {
	Message string `json:"message"`
}

type ResponseErrorNotFound struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Success string `json:"Success"`
}

type CategoryInput struct {
	CategoryName string `json:"category"`
}

type ProductInput struct {
	CategoryId  uint    `json:"categoryId"`
	ProductName string  `json:"namaProduct"`
	HargaBarang float64 `json:"hargaProduct"`
	Type        string  `json:"type"`
	Image       string  `json:"image"`
}
