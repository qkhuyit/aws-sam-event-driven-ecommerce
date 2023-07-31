package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	internalTypes "github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/utils"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/repositories"
	"github.com/sirupsen/logrus"
)

type ProductService interface {
	Create(ctx context.Context, product internalTypes.Product, attributes []internalTypes.ProductAttributeValue) (*internalTypes.Product, error)
	GetList(ctx context.Context) ([]internalTypes.Product, error)
}

func NewProductService(logger *logrus.Logger,
	productRepository repositories.ProductRepository,
	attributeRepository repositories.ProductAttributeRepository,
	attributeValueRepository repositories.ProductAttributeValueRepository,
	dynamoDbClient *dynamodb.Client,
) ProductService {
	return &productServiceImpl{
		logger:                   logger,
		productRepository:        productRepository,
		attributeRepository:      attributeRepository,
		attributeValueRepository: attributeValueRepository,
		dynamoDbClient:           dynamoDbClient,
	}
}

type productServiceImpl struct {
	logger                   *logrus.Logger
	productRepository        repositories.ProductRepository
	attributeRepository      repositories.ProductAttributeRepository
	attributeValueRepository repositories.ProductAttributeValueRepository
	dynamoDbClient           *dynamodb.Client
}

func (productService productServiceImpl) GetList(ctx context.Context) (products []internalTypes.Product, err error) {
	products = make([]internalTypes.Product, 0)
	var token map[string]types.AttributeValue
	for {
		input := &dynamodb.ScanInput{
			TableName:         aws.String("Product"),
			ExclusiveStartKey: token,
		}

		result, err := productService.dynamoDbClient.Scan(ctx, input)
		if err != nil {
			return nil, err
		}

		var fetchedProducts []internalTypes.Product
		err = attributevalue.UnmarshalListOfMaps(result.Items, &fetchedProducts)
		if err != nil {
			return nil, err
		}

		products = append(products, fetchedProducts...)
		token = result.LastEvaluatedKey
		if token == nil {
			break
		}
	}

	return products, nil
}

func (productService productServiceImpl) Create(
	ctx context.Context,
	product internalTypes.Product,
	attributes []internalTypes.ProductAttributeValue) (*internalTypes.Product, error) {

	product.Id = uuid.New().String()
	product.Alias = utils.GenerateAliasString(product.Name)
	_, err := productService.productRepository.Create(ctx, product)
	if err != nil {
		productService.logger.Info("[productServiceImpl#Create] failed create product. err: ", err)
		return nil, err
	}

	for _, at := range attributes {
		_, err = productService.attributeValueRepository.Create(ctx, internalTypes.ProductAttributeValue{
			Value:       at.Value,
			ProductId:   product.Id,
			AttributeId: at.AttributeId,
		})

		if err != nil {
			productService.logger.Info("[productServiceImpl#Create] failed create product attribute value. err: ", err)
			return nil, err
		}
	}

	return &product, nil
}
