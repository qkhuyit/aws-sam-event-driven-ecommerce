package types

type OrderStatus string

const (
	ORDER_STATUS_CREATED   OrderStatus = "CREATED"
	ORDER_STATUS_CANCELED  OrderStatus = "CANCELED"
	ORDER_STATUS_CONFIRMED OrderStatus = "ONFIRMED"
	DELIVERING             OrderStatus = "DELIVERING"
	DELIVERED              OrderStatus = "DELIVERED"
	COMPLETED              OrderStatus = "COMPLETED"
)

func (o OrderStatus) ToString() string {
	return string(o)
}

func NewOrderStatus(str string) OrderStatus {
	switch str {
	case "CREATED":
		return ORDER_STATUS_CREATED
	case "CANCELED":
		return ORDER_STATUS_CANCELED
	}

	return ""
}

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
