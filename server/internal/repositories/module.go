package repositories

import "go.uber.org/fx"

// Module exports dependency to container
var RepositoryModule = fx.Options(
	fx.Provide(
		NewProductRepository,
		NewProductAttributeRepository,
		NewProductAttributeValueRepository,
		NewOrderRepository,
		NewOrderDetailRepository,
	),
)
