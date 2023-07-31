package converters

import (
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/models"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
)

type ProductConverter struct {
}

func (c ProductConverter) ConvertCreateProductModelToProduct(m models.CreateProductModel) types.Product {
	return types.Product{}
}

func (c ProductConverter) ConvertAttributeValueModelToProductAttributeValue(m models.CreateProductModel) []types.ProductAttributeValue {
	result := make([]types.ProductAttributeValue, len(m.Attributes))
	return result
}
