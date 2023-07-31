package models

type CreateOrderRequestModel struct {
	FullName string             `json:"full_name" validate:"required,max=125"`
	Email    string             `json:"email" validate:"required,max=125"`
	Address  string             `json:"address" validate:"required,max=256"`
	Phone    string             `json:"phone" validate:"required,max=50"`
	Items    []OrderDetailModel `json:"items" validate:"required"`
}

type OrderDetailModel struct {
	ProductId string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,gt=1"`
}
