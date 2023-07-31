package internal

import (
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/core"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/handlers"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/repositories"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/services"
	"go.uber.org/fx"
)

// Module exports dependency to container
var RootModule = fx.Options(
	core.CoreModule,
	repositories.RepositoryModule,
	services.ServiceModule,
	handlers.HandlerModule,
)
