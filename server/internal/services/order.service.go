package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/errors"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/repositories"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type OrderService interface {
	Create(ctx context.Context, order types.Order, details []types.OrderDetail, discountCode *string) (*types.Order, error)
	Cancel(ctx context.Context, id string) error
	Confirm(ctx context.Context, id string) error
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

func (orderService orderServiceImpl) Confirm(ctx context.Context, id string) error {
	orderService.logger.Infoln("[orderServiceImpl#Confirm] BEGIN confirm order")
	defer orderService.logger.Infoln("[orderServiceImpl#Confirm] END confirm order")

	order, err := orderService.orderRepository.FindById(ctx, id)
	if err != nil {
		orderService.logger.Errorln("[orderServiceImpl#Confirm] fail get order detail")
		return err
	}

	if strings.Compare(order.Status, string(types.ORDER_STATUS_CREATED)) != 0 {
		orderService.logger.Errorln("[orderServiceImpl#Confirm] can't confirm order processed")
		return errors.NewModelInvalidError(fmt.Errorf(""))
	}

	err = orderService.orderRepository.Patch(ctx, id, map[string]interface{}{
		"status": types.ORDER_STATUS_CONFIRMED.ToString(),
	})
	if err != nil {
		orderService.logger.Errorln("[orderServiceImpl#Confirm] fail save order.")
		return err
	}

	return nil
}

func (orderService orderServiceImpl) Cancel(ctx context.Context, id string) error {
	orderService.logger.Infoln("[orderServiceImpl#Cancel] BEGIN cancel order")
	defer orderService.logger.Infoln("[orderServiceImpl#Cancel] END cancel order")

	order, err := orderService.orderRepository.FindById(ctx, id)
	if err != nil {
		orderService.logger.Errorln("[orderServiceImpl#Cancel] fail get order detail")
		return err
	}

	if strings.Compare(order.Status, string(types.ORDER_STATUS_CREATED)) != 0 {
		orderService.logger.Errorln("[orderServiceImpl#Cancel] can't cancel order processed")
		return errors.NewModelInvalidError(fmt.Errorf(""))
	}

	err = orderService.orderRepository.Patch(ctx, id, map[string]interface{}{
		"status": types.ORDER_STATUS_CANCELED.ToString(),
	})
	if err != nil {
		orderService.logger.Errorln("[orderServiceImpl#Cancel] fail save order.")
		return err
	}

	return nil
}

func (orderService orderServiceImpl) Create(ctx context.Context, order types.Order, details []types.OrderDetail, discountCode *string) (*types.Order, error) {
	orderService.logger.Infoln("[orderServiceImpl#Create] BEGIN create order")
	defer orderService.logger.Infoln("[orderServiceImpl#Create] END create order")

	order.Id = uuid.New().String()
	order.CreatedOn = time.Now().Format("2006-01-02 15:04:05")
	order.Status = string(types.ORDER_STATUS_CREATED)
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
