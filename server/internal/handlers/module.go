package handlers

import "go.uber.org/fx"

// Module exports dependency to container
var HandlerModule = fx.Options(
	fx.Provide(
		NewProductHandler,
		NewOrderHandler,
	),
)
