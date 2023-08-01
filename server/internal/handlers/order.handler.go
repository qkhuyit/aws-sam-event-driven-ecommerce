package handlers

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator/v10"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/converters"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/errors"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/models"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/transform"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/utils"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/services"
	"github.com/sirupsen/logrus"
)

type OrderHandler interface {
	Create(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	ChangeStatus(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func NewOrderHandler(logger *logrus.Logger, orderService services.OrderService) OrderHandler {
	return &orderHandlerImpl{
		logger:         logger,
		orderService:   orderService,
		validate:       validator.New(),
		orderConverter: converters.OrderConverter{},
	}
}

type orderHandlerImpl struct {
	logger         *logrus.Logger
	orderService   services.OrderService
	validate       *validator.Validate
	orderConverter converters.OrderConverter
}

func (orderHandler orderHandlerImpl) ChangeStatus(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	orderHandler.logger.Infoln("[orderHandlerImpl#ChangeStatus] BEGIN change order status")
	defer orderHandler.logger.Infoln("[orderHandlerImpl#ChangeStatus] END change order status")

	id, ok := req.PathParameters["id"]
	if !ok {
		return transform.SendAppError(errors.NewModelInvalidError(fmt.Errorf("id is require")))
	}

	model, err := utils.JsonDeserialize[models.ChangeStatusRequestModel](req.Body)
	if err != nil {
		return transform.SendAppError(errors.NewModelInvalidError(err))
	}

	status := types.NewOrderStatus(model.Status)

	err = orderHandler.orderService.ChangeStatus(ctx, id, status)
	if err != nil {
		return transform.SendError(err)
	}

	return transform.SendSuccessWithData(nil)
}

func (orderHandler orderHandlerImpl) Create(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	orderHandler.logger.Infoln("[orderHandlerImpl#Create] BEGIN create new order")
	defer orderHandler.logger.Infoln("[orderHandlerImpl#Create] END create new order")
	model, err := utils.JsonDeserialize[models.CreateOrderRequestModel](req.Body)
	if err != nil {
		orderHandler.logger.Errorln("[orderHandlerImpl#Create] failed deserialize request body. err: ", err)
		return transform.SendAppError(errors.NewModelInvalidError(err))
	}

	err = orderHandler.validate.Struct(model)
	if err != nil {
		orderHandler.logger.Errorln("[orderHandlerImpl#Create] failed validate request body schema. err: ", err)
		return transform.SendAppError(errors.NewModelInvalidError(err))
	}

	order := orderHandler.orderConverter.ConvertCreateOrderModelToOrder(*model)
	items := orderHandler.orderConverter.ConvertCreateOrderModelToOrderDetails(*model)
	newOrder, err := orderHandler.orderService.Create(ctx, order, items, nil)
	if err != nil {
		orderHandler.logger.Errorln("[orderHandlerImpl#Create] failed save order. err: ", err)
		return transform.SendError(err)
	}

	return transform.SendSuccessWithData(newOrder)
}
