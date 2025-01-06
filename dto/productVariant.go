package dto

type ProductVariant struct {
	Id        uint    `json:"id"`
	ProductId uint    `json:"productId"`
	Color     string  `json:"color" validate:"required,max=50"`
	Size      string  `json:"size" validate:"required,max=50"`
	Stock     uint64  `json:"stock" validate:"required,number"`
	Image     string  `gorm:"varchar(200);not null"`
	Product   Product `gorm:"foreignKey:productId;references:Id"`
}
