package types

type OrderStatus string

const (
	ORDER_CREATED OrderStatus = "CREATED"
	CANCELED      OrderStatus = "CANCELED"
	CONFIRMED     OrderStatus = "CONFIRMED"
	DELIVERING    OrderStatus = "DELIVERING"
	DELIVERED     OrderStatus = "DELIVERED"
	COMPLETED     OrderStatus = "COMPLETED"
)

type Order struct {
	Id         string `json:"id"  dynamodbav:"id"`
	FullName   string `json:"full_name"  dynamodbav:"full_name"`
	Email      string `json:"email"  dynamodbav:"email"`
	Address    string `json:"address"  dynamodbav:"address"`
	Phone      string `json:"phone"  dynamodbav:"phone"`
	TotalPrice int    `json:"price"  dynamodbav:"price"`
	Status     string `json:"status"  dynamodbav:"status"`
	CreatedOn  string `json:"created_on"  dynamodbav:"created_on"`
}

type OrderDetail struct {
	ProductId string `json:"product_id" dynamodbav:"product_id"`
	OrderId   string `json:"order_id" dynamodbav:"order_id"`
	Quantity  int    `json:"quantity" dynamodbav:"quantity"`
	Price     int    `json:"price" dynamodbav:"price"`
}

type OrderHistory struct {
	OrderId        string `json:"order_id"`
	HistoryType    string `json:"history_type"`
	HistoryMessage string `json:"history_message"`
	CreatedOn      string `json:"created_on"`
	CreatedBy      string `json:"created_by"`
}
