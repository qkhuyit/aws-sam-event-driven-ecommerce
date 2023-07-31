package converters

import (
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/models"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
)

type OrderConverter struct {
}

func (c OrderConverter) ConvertCreateOrderModelToOrder(m models.CreateOrderRequestModel) types.Order {
	return types.Order{
		Email:    m.Email,
		Address:  m.Address,
		FullName: m.FullName,
		Phone:    m.Phone,
	}
}

func (c OrderConverter) ConvertCreateOrderModelToOrderDetails(m models.CreateOrderRequestModel) []types.OrderDetail {
	items := make([]types.OrderDetail, len(m.Items))
	for i, item := range m.Items {
		items[i] = types.OrderDetail{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
	}
	return items
}
