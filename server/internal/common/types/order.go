package types

type OrderStatus string

const (
	ORDER_STATUS_CREATED                             OrderStatus = "CREATED"
	ORDER_STATUS_CANCELED                            OrderStatus = "CANCELED"
	ORDER_STATUS_CONFIRMED                           OrderStatus = "CONFIRMED"
	ORDER_STATUS_SUPPLIER_CONFIRMED                  OrderStatus = "SUPPLIER_CONFIRMED"
	ORDER_STATUS_SUPPLIER_CANCELED                   OrderStatus = "SUPPLIER_CANCELED"
	ORDER_STATUS_SUPPLIER_PACKING                    OrderStatus = "PACKING"
	ORDER_STATUS_SUPPLIER_PACKED                     OrderStatus = "PACKED"
	ORDER_STATUS_WAITING_DELIVER_TO_TRANSPORT_VENDOR OrderStatus = "WAITING_DELIVER_TO_TRANSPORT_VENDOR"
	ORDER_STATUS_DELIVERED_TO_TRANSPORT_VENDOR       OrderStatus = "DELIVERED_TO_TRANSPORT_VENDOR"
	ORDER_STATUS_DELIVER_TO_TRANSPORT_VENDOR_FAILED  OrderStatus = "DELIVER_TO_TRANSPORT_VENDOR_FAILED"
	ORDER_STATUS_DELEVERING                          OrderStatus = "DELEVERING"
	ORDER_STATUS_DELVEVER_FALIED                     OrderStatus = "DELVEVER_FALIED"
	ORDER_STATUS_DELVEVER_REJECTED                   OrderStatus = "DELVEVER_REJECTED"
	ORDER_STATUS_ROLLBACK_PACK                       OrderStatus = "ROLLBACK_PACK"
	ORDER_STATUS_COMPLETED                           OrderStatus = "COMPLETED"
	ORDER_STATUS_FAILED                              OrderStatus = "ORDER_STATUS_FAILED"
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
