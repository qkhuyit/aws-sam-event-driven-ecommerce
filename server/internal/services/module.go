package services

import "go.uber.org/fx"

// Module exports dependency to container
var ServiceModule = fx.Options(
	fx.Provide(
		NewProductService,
		NewOrderService,
	),
)
