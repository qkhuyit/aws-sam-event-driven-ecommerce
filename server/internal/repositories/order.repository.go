package repositories

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	internalTypes "github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
	"github.com/sirupsen/logrus"
)

type OrderRepository interface {
	Create(ctx context.Context, order internalTypes.Order) (newProduct *internalTypes.Order, err error)
}

func NewOrderRepository(dbClient *dynamodb.Client, logger *logrus.Logger) OrderRepository {
	return &orderRepositoryImpl{
		dbClient:       dbClient,
		logger:         logger,
		orderTableName: aws.String("Order"),
	}
}

type orderRepositoryImpl struct {
	dbClient       *dynamodb.Client
	logger         *logrus.Logger
	orderTableName *string
}

func (orderRepository orderRepositoryImpl) Create(ctx context.Context, order internalTypes.Order) (newOrder *internalTypes.Order, err error) {
	orderRepository.logger.Infoln("[orderRepository#orderRepositoryImpl] BEGIN create order")
	defer orderRepository.logger.Infoln("[orderRepository#orderRepositoryImpl] END create order")

	av, err := attributevalue.MarshalMap(order)
	if err != nil {
		orderRepository.logger.Errorln("[orderRepository#orderRepositoryImpl] failed to marshal order to attribute value map.")
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: orderRepository.orderTableName,
	}

	output, err := orderRepository.dbClient.PutItem(ctx, input)
	if err != nil {
		orderRepository.logger.Errorln("[orderRepository#orderRepositoryImpl] failed to put order to dynamodb. err: ", err)
		return nil, err
	}

	attributevalue.UnmarshalMap(output.Attributes, newOrder)

	return newOrder, nil
}
