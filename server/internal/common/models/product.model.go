package models

type CreateProductModel struct {
	Name       string                       `json:"name"`
	Quantity   int                          `json:"quantity"`
	Attributes []ProductAttributeValueModel `json:"attributes"`
	BrandId    string                       `json:"brand_id"`
	CategoryId string                       `json:"category_id"`
}

type ProductAttributeValueModel struct {
	AttributeId string `json:"attribute_id"`
	Value       string `json:"value"`
}
