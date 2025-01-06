package products

import "time"

// response documentasi

type ResponseErrorBadRequest struct {
	ErrorFields []map[string]string
}

type ResponseErrorInternalServer struct {
	Error string `json:"error"`
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
	CategoryId     uint    `json:"categoryId"`
	ProductName    string  `json:"productName"`
	Price          float64 `json:"price"`
	Desc           string  `json:"desc"`
	DateRelase     string  `json:"dateRelase"`
	Type           string  `json:"typeProduct"`
	Image          string  `json:"image"`
	CategoryGender string  `json:"categoryGender"`
}

type ResponseProduct struct {
	Id             uint      `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CategoryId     uint      `json:"categoryId"`
	ProductName    string    `json:"namaProduct"`
	HargaBarang    float64   `json:"hargaProduct"`
	Type           string    `json:"type"`
	Image          string    `json:"image"`
	CategoryGender string    `json:"categoryGender"`
}
