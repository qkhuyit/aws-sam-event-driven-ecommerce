package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/errors"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/repositories"
	"github.com/sirupsen/logrus"
	"time"
)

type OrderService interface {
	Create(ctx context.Context, order types.Order, details []types.OrderDetail, discountCode *string) (*types.Order, error)
}

func NewOrderService(orderRepository repositories.OrderRepository,
	orderDetailRepository repositories.OrderDetailRepository,
	productRepository repositories.ProductRepository,
	logger *logrus.Logger) OrderService {
	return &orderServiceImpl{
		orderRepository:       orderRepository,
		orderDetailRepository: orderDetailRepository,
		productRepository:     productRepository,
		logger:                logger,
	}
}

type orderServiceImpl struct {
	orderRepository       repositories.OrderRepository
	orderDetailRepository repositories.OrderDetailRepository
	productRepository     repositories.ProductRepository
	logger                *logrus.Logger
}

func (orderService orderServiceImpl) Create(ctx context.Context, order types.Order, details []types.OrderDetail, discountCode *string) (*types.Order, error) {
	orderService.logger.Infoln("[orderServiceImpl#Create] BEGIN create order")
	defer orderService.logger.Infoln("[orderServiceImpl#Create] END create order")

	order.Id = uuid.New().String()
	order.CreatedOn = time.Now().Format("2006-01-02 15:04:05")
	order.Status = string(types.ORDER_CREATED)
	order.TotalPrice = 0

	//TODO validate product list
	for idx, item := range details {
		product, err := orderService.productRepository.FindById(ctx, item.ProductId)
		if err != nil {
			orderService.logger.Errorln("[orderServiceImpl#Create] failed get product id ", item.ProductId)
			return nil, err
		}
		if product == nil {
			orderService.logger.Errorln("[orderServiceImpl#Create] failed get product id ", item.ProductId)
			return nil, errors.NewResourceNotfoundError("Product", "id", item.ProductId)
		}

		details[idx].OrderId = order.Id
		details[idx].Price = product.Price
		order.TotalPrice += details[idx].Price
	}

	_, err := orderService.orderRepository.Create(ctx, order)
	if err != nil {
		orderService.logger.Errorln("[orderServiceImpl#Create] fail save order. err: ", err)
		return nil, err
	}

	err = orderService.orderDetailRepository.CreateAll(ctx, details)
	if err != nil {
		orderService.logger.Errorln("[orderServiceImpl#Create] fail save order. err: ", err)
		return nil, err
	}

	return &order, err
}
