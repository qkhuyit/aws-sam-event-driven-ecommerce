package utils

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func ToAttributeValueMap[T any](inp T) map[string]types.AttributeValue {
	av, err := attributevalue.MarshalMap(inp)
	if err != nil {
		return nil
	}

	return av
}
