package core

import (
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/core/cognito"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/core/config"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/core/dynamodb"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/core/logger"
	"go.uber.org/fx"
)

// Module exports dependency to container
var CoreModule = fx.Options(
	fx.Provide(
		logger.NewLogger,
		config.NewAppConfig,
		dynamodb.NewDynamoDb,
		cognito.NewCognitoIdentityProvider,
	),
)
