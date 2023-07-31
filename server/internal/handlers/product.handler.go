package handlers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator/v10"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/converters"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/errors"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/models"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/transform"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/utils"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/services"
	"github.com/sirupsen/logrus"
)

type ProductHandler interface {
	GetList(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	GetDetail(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Create(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Update(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Delete(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func NewProductHandler(logger *logrus.Logger, productService services.ProductService) ProductHandler {
	return &productHandlerImpl{
		logger:           logger,
		productService:   productService,
		validate:         validator.New(),
		productConverter: converters.ProductConverter{},
	}
}

type productHandlerImpl struct {
	logger           *logrus.Logger
	validate         *validator.Validate
	productService   services.ProductService
	productConverter converters.ProductConverter
}

func (handler productHandlerImpl) GetList(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	handler.logger.Info("[productHandlerImpl#GetList] BEGIN BEGIN Get list product")
	defer handler.logger.Info("[productHandlerImpl#GetList] END Get list product")

	products, err := handler.productService.GetList(ctx)
	if err != nil {
		return transform.SendError(err)
	}

	return transform.SendSuccessWithData(products)
}

func (handler productHandlerImpl) GetDetail(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (handler productHandlerImpl) Create(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	handler.logger.Info("[productHandlerImpl#Create] BEGIN Create product")
	defer handler.logger.Info("[productHandlerImpl#Create] END Create product")

	productModel, err := utils.JsonDeserialize[models.CreateProductModel](req.Body)
	if err != nil {
		handler.logger.Info("[productHandlerImpl#Create] failed deserialize request body to object, err: ", err)
		return transform.SendAppError(errors.NewModelInvalidError(err))
	}

	err = handler.validate.Struct(productModel)
	if err != nil {
		handler.logger.Info("[productHandlerImpl#Create] failed validate request body schema, err: ", err)
		return transform.SendAppError(errors.NewModelInvalidError(err))
	}

	newProductInfo := handler.productConverter.ConvertCreateProductModelToProduct(*productModel)
	newProductAttributesInfo := handler.productConverter.ConvertAttributeValueModelToProductAttributeValue(*productModel)

	newProduct, err := handler.productService.Create(ctx, newProductInfo, newProductAttributesInfo)
	if err != nil {
		handler.logger.Info("[productHandlerImpl#Create] failed to create new product, err: ", err)
		return transform.SendError(err)
	}

	return transform.SendSuccessWithData(newProduct)
}

func (handler productHandlerImpl) Update(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (handler productHandlerImpl) Delete(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//TODO implement me
	panic("implement me")
}
