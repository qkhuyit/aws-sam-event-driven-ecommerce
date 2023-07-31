package repositories

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	internalTypes "github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
	"github.com/sirupsen/logrus"
)

type ProductAttributeRepository interface {
	FindById(ctx context.Context, id string) (*internalTypes.ProductAttribute, error)
}

func NewProductAttributeRepository(dynamodbClient *dynamodb.Client) ProductAttributeRepository {
	return &productAttributeRepositoryImpl{}
}

type productAttributeRepositoryImpl struct {
	dynamoDbClient            *dynamodb.Client
	logger                    *logrus.Logger
	productAttributeTableName *string
}

func (productAttributeRepository productAttributeRepositoryImpl) FindById(ctx context.Context, id string) (attribute *internalTypes.ProductAttribute, err error) {
	productAttributeRepository.logger.Info("[productAttributeRepositoryImpl#FindById] BEGIN FindById")
	defer productAttributeRepository.logger.Info("[productAttributeRepositoryImpl#FindById] END FindById")

	aliasAttributeValue, _ := attributevalue.Marshal(id)
	response, err := productAttributeRepository.dynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key:       map[string]types.AttributeValue{"id": aliasAttributeValue},
		TableName: productAttributeRepository.productAttributeTableName,
	})

	if err != nil {
		productAttributeRepository.logger.Info("[productAttributeRepositoryImpl#FindById] Failed to GetItem id ", id)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(response.Item, &attribute)
	if err != nil {
		productAttributeRepository.logger.Info("[productAttributeRepositoryImpl#FindById] Failed to UnmarshalMap. err: ", err)
		return nil, err
	}

	return attribute, err
}
