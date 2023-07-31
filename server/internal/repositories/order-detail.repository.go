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

type OrderDetailRepository interface {
	Create(ctx context.Context, item internalTypes.OrderDetail) (result *internalTypes.OrderDetail, err error)
	CreateAll(ctx context.Context, items []internalTypes.OrderDetail) (err error)
}

func NewOrderDetailRepository(dbClient *dynamodb.Client, logger *logrus.Logger) OrderDetailRepository {
	return &orderDetailRepositoryImpl{
		dbClient:             dbClient,
		logger:               logger,
		orderDetailTableName: aws.String("OrderDetail"),
	}
}

type orderDetailRepositoryImpl struct {
	dbClient             *dynamodb.Client
	logger               *logrus.Logger
	orderDetailTableName *string
}

func (orderDetailRepository orderDetailRepositoryImpl) Create(ctx context.Context, item internalTypes.OrderDetail) (result *internalTypes.OrderDetail, err error) {
	orderDetailRepository.logger.Info("[orderDetailRepositoryImpl#Create] BEGIN create order detail record")
	defer orderDetailRepository.logger.Info("[orderDetailRepositoryImpl#Create] END create order detail record")

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		orderDetailRepository.logger.Info("[orderDetailRepositoryImpl#Create] failed to marshal order detail to attribute value map.")
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: orderDetailRepository.orderDetailTableName,
	}

	output, err := orderDetailRepository.dbClient.PutItem(ctx, input)
	if err != nil {
		orderDetailRepository.logger.Info("[orderDetailRepositoryImpl#Create] failed to put order detail to dynamodb.")
		return nil, err
	}

	attributevalue.UnmarshalMap(output.Attributes, result)

	return result, nil
}

func (orderDetailRepository orderDetailRepositoryImpl) CreateAll(ctx context.Context, items []internalTypes.OrderDetail) (err error) {
	orderDetailRepository.logger.Info("[orderDetailRepositoryImpl#CreateAll] BEGIN create order detail record")
	defer orderDetailRepository.logger.Info("[orderDetailRepositoryImpl#CreateAll] END create order detail record")

	transactItems := make([]types.TransactWriteItem, len(items))
	for i, item := range items {
		av, err := attributevalue.MarshalMap(item)
		if err != nil {
			orderDetailRepository.logger.Info("[orderDetailRepositoryImpl#Create] failed to marshal order detail to attribute value map.")
			return err
		}

		transactItems[i] = types.TransactWriteItem{
			Put: &types.Put{
				TableName: orderDetailRepository.orderDetailTableName,
				Item:      av,
			},
		}
	}

	_, err = orderDetailRepository.dbClient.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{TransactItems: transactItems})
	if err != nil {
		defer orderDetailRepository.logger.Info("[orderDetailRepositoryImpl#CreateAll] failed to put order detail to dynamodb.")
		return err
	}

	return nil
}
