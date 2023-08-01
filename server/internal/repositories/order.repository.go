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

type OrderRepository interface {
	Create(ctx context.Context, order internalTypes.Order) (newProduct *internalTypes.Order, err error)
	Update(ctx context.Context, order internalTypes.Order) (err error)
	Patch(ctx context.Context, id string, fields map[string]interface{}) (err error)
	FindById(ctx context.Context, id string) (order *internalTypes.Order, err error)
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

func (orderRepository orderRepositoryImpl) Update(ctx context.Context, order internalTypes.Order) (err error) {
	orderRepository.logger.Infoln("[orderRepositoryImpl#Update] BEGIN update order")
	defer orderRepository.logger.Infoln("[orderRepositoryImpl#Update] END update order")

	av, err := attributevalue.MarshalMap(order)
	if err != nil {
		orderRepository.logger.Errorln("[orderRepositoryImpl#Update] failed to marshal order to attribute value map.")
		return err
	}

	attributeUpdates := make(map[string]types.AttributeValueUpdate)
	var key types.AttributeValue
	for k, value := range av {
		if k == "id" {
			key = value
		} else {
			attributeUpdates[k] = types.AttributeValueUpdate{
				Action: types.AttributeActionPut,
				Value:  value,
			}
		}
	}

	input := &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"id": key,
		},
		AttributeUpdates: attributeUpdates,
		TableName:        orderRepository.orderTableName,
	}

	_, err = orderRepository.dbClient.UpdateItem(ctx, input)

	if err != nil {
		orderRepository.logger.Errorln("[orderRepositoryImpl#Update] failed to put order to dynamodb. err: ", err)
		return err
	}

	return nil
}

func (orderRepository orderRepositoryImpl) Patch(ctx context.Context, id string, fields map[string]interface{}) (err error) {
	orderRepository.logger.Infoln("[orderRepositoryImpl#Patch] BEGIN patch update order")
	defer orderRepository.logger.Infoln("[orderRepositoryImpl#Patch] END path update order")

	attributeUpdates := make(map[string]types.AttributeValueUpdate)
	key, err := attributevalue.Marshal(id)
	if err != nil {
		orderRepository.logger.Errorln("[orderRepositoryImpl#Patch] failed to marshal order id to attribute value.")
		return err
	}

	for k, value := range fields {
		attributeValue, err := attributevalue.Marshal(value)
		if err != nil {
			orderRepository.logger.Errorln("[orderRepositoryImpl#Patch] failed to marshal order id to attribute value.")
			return err
		}

		attributeUpdates[k] = types.AttributeValueUpdate{
			Action: types.AttributeActionPut,
			Value:  attributeValue,
		}
	}

	input := &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"id": key,
		},
		AttributeUpdates: attributeUpdates,
		TableName:        orderRepository.orderTableName,
	}

	_, err = orderRepository.dbClient.UpdateItem(ctx, input)

	if err != nil {
		orderRepository.logger.Errorln("[orderRepositoryImpl#Patch] failed to put order to dynamodb. err: ", err)
		return err
	}

	return nil
}

func (orderRepository orderRepositoryImpl) FindById(ctx context.Context, id string) (order *internalTypes.Order, err error) {
	//TODO implement me
	panic("implement me")
}

func (orderRepository orderRepositoryImpl) Create(ctx context.Context, order internalTypes.Order) (newOrder *internalTypes.Order, err error) {
	orderRepository.logger.Infoln("[orderRepositoryImpl#Create] BEGIN create order")
	defer orderRepository.logger.Infoln("[orderRepositoryImpl#Create] END create order")

	av, err := attributevalue.MarshalMap(order)
	if err != nil {
		orderRepository.logger.Errorln("[orderRepositoryImpl#Create] failed to marshal order to attribute value map.")
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: orderRepository.orderTableName,
	}

	output, err := orderRepository.dbClient.PutItem(ctx, input)
	if err != nil {
		orderRepository.logger.Errorln("[orderRepositoryImpl#Create] failed to put order to dynamodb. err: ", err)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(output.Attributes, newOrder)
	if err != nil {
		orderRepository.logger.Errorln("[orderRepositoryImpl#Create] failed to unmarshal put result to order. err: ", err)
		return nil, err
	}

	return newOrder, nil
}
