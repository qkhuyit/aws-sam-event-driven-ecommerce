package repositories

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	internalTypes "github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
	"github.com/sirupsen/logrus"
)

type ProductAttributeValueRepository interface {
	Create(ctx context.Context, attribute internalTypes.ProductAttributeValue) (*internalTypes.ProductAttributeValue, error)
}

func NewProductAttributeValueRepository() ProductAttributeValueRepository {
	return &productAttributeValueRepositoryImpl{}
}

type productAttributeValueRepositoryImpl struct {
	dynamoDbClient                 *dynamodb.Client
	productAttributeValueTableName *string
	logger                         *logrus.Logger
}

func (productAttributeValueRepository productAttributeValueRepositoryImpl) Create(ctx context.Context, attribute internalTypes.ProductAttributeValue) (newAttributeVal *internalTypes.ProductAttributeValue, err error) {
	av, err := attributevalue.MarshalMap(attribute)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: productAttributeValueRepository.productAttributeValueTableName,
	}

	output, err := productAttributeValueRepository.dynamoDbClient.PutItem(ctx, input)
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(output.Attributes, &newAttributeVal)
	if err != nil {
		return nil, err
	}

	return newAttributeVal, nil
}
