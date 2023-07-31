package repositories

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	internalTypes "github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
	"github.com/sirupsen/logrus"
)

type ProductRepository interface {
	FindByAlias(ctx context.Context, alias string) (*internalTypes.Product, error)
	FindById(ctx context.Context, id string) (newProduct *internalTypes.Product, err error)
	Create(ctx context.Context, product internalTypes.Product) (newProduct *internalTypes.Product, err error)
}

func NewProductRepository(client *dynamodb.Client, logger *logrus.Logger) ProductRepository {
	return &productRepositoryImpl{
		dynamoDbClient:   client,
		logger:           logger,
		productTableName: aws.String("Product"),
	}
}

type productRepositoryImpl struct {
	dynamoDbClient   *dynamodb.Client
	productTableName *string
	logger           *logrus.Logger
}

func (productRepository productRepositoryImpl) Create(ctx context.Context, product internalTypes.Product) (newProduct *internalTypes.Product, err error) {
	av, err := attributevalue.MarshalMap(product)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: productRepository.productTableName,
	}

	productRepository.dynamoDbClient.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{
				Put: &types.Put{
					TableName: productRepository.productTableName,
					Item:      map[string]types.AttributeValue{},
				},
			},
		},
	})

	output, err := productRepository.dynamoDbClient.PutItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	attributevalue.UnmarshalMap(output.Attributes, &newProduct)

	return newProduct, nil
}

func (productRepository productRepositoryImpl) FindByAlias(ctx context.Context, alias string) (newProduct *internalTypes.Product, err error) {
	productRepository.logger.Info("[productRepositoryImpl#FindByAlias] BEGIN FindByAlias")
	defer productRepository.logger.Info("[productRepositoryImpl#FindByAlias] END FindByAlias")
	aliasAttributeValue, _ := attributevalue.Marshal(alias)
	response, err := productRepository.dynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key:       map[string]types.AttributeValue{"alias": aliasAttributeValue},
		TableName: productRepository.productTableName,
	})

	if err != nil {
		productRepository.logger.Info("[productRepositoryImpl#FindByAlias] Failed to GetItem alias", alias)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(response.Item, &newProduct)
	if err != nil {
		return nil, err
	}

	return newProduct, err
}

func (productRepository productRepositoryImpl) FindById(ctx context.Context, id string) (newProduct *internalTypes.Product, err error) {
	productRepository.logger.Info("[productRepositoryImpl#FindByAlias] BEGIN FindById")
	defer productRepository.logger.Info("[productRepositoryImpl#FindByAlias] END FindById")
	aliasAttributeValue, _ := attributevalue.Marshal(id)
	response, err := productRepository.dynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key:       map[string]types.AttributeValue{"id": aliasAttributeValue},
		TableName: productRepository.productTableName,
	})

	if err != nil {
		productRepository.logger.Info("[productRepositoryImpl#FindByAlias] Failed to GetItem id ", id)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(response.Item, &newProduct)
	if err != nil {
		return nil, err
	}

	return newProduct, err
}
