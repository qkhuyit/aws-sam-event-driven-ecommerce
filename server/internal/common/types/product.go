package types

type Product struct {
	Id    string `json:"id" dynamodbav:"id"`
	Name  string `json:"name" dynamodbav:"name"`
	Alias string `json:"alias" dynamodbav:"alias"`
	Price int    `json:"price"`
}

type ProductAttribute struct {
	Id     string `json:"id" dynamodbav:"id"`
	Name   string `json:"name" dynamodbav:"name"`
	Type   string `json:"type" dynamodbav:"type"`
	Remark string `json:"remark" dynamodbav:"remark"`
}

type ProductAttributeValue struct {
	ProductId   string `json:"product_id" dynamodbav:"product_id"`
	AttributeId string `json:"attribute_id" dynamodbav:"attribute_id"`
	Value       string `json:"value" dynamodbav:"value"`
}

type Brand struct {
	Id          string `json:"id" dynamodbav:"id"`
	Name        string `json:"name" dynamodbav:"name"`
	Alias       string `json:"alias" dynamodbav:"alias"`
	Description string `json:"description" dynamodbav:"description"`
}

type Category struct {
	Id          string `json:"id" dynamodbav:"id"`
	Name        string `json:"name" dynamodbav:"name"`
	Alias       string `json:"alias" dynamodbav:"alias"`
	Description string `json:"description" dynamodbav:"description"`
}

type Supplier struct {
	Id          string `json:"id" dynamodbav:"id"`
	Name        string `json:"name" dynamodbav:"name"`
	Alias       string `json:"alias" dynamodbav:"alias"`
	Description string `json:"description" dynamodbav:"description"`
	OwnerId     string `json:"owner_id" dynamodbav:"owner_id"`
}
