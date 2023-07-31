package transform

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
	"net/http"
)

const (
	headerFieldContentType           = "Content-Type"
	headerContentTypeApplicationJson = "application/json"
)

func SendSuccessWithData(data interface{}) (events.APIGatewayProxyResponse, error) {
	restResult := LambdaResponseModel{
		Code:      200,
		Success:   true,
		Data:      data,
		Message:   "OK",
		MessageId: "SUCCESSES",
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			headerFieldContentType: headerContentTypeApplicationJson,
		},
		Body: restResult.ToJson(),
	}, nil
}

func SendError(err error) (events.APIGatewayProxyResponse, error) {
	restResult := LambdaResponseModel{
		Code:      http.StatusInternalServerError,
		Success:   true,
		Data:      nil,
		Message:   err.Error(),
		MessageId: "0E00500-000",
		Error:     err,
	}
	return events.APIGatewayProxyResponse{
		StatusCode: restResult.Code,
		Headers: map[string]string{
			headerFieldContentType: headerContentTypeApplicationJson,
		},
		Body: restResult.ToJson(),
	}, nil
}

func SendAppError(appError types.AppError) (events.APIGatewayProxyResponse, error) {
	restResult := LambdaResponseModel{
		Code:      appError.Status(),
		Success:   true,
		Data:      nil,
		Message:   appError.Message(),
		MessageId: appError.MessageId(),
	}
	return events.APIGatewayProxyResponse{
		StatusCode: restResult.Code,
		Headers: map[string]string{
			headerFieldContentType: headerContentTypeApplicationJson,
		},
		Body: restResult.ToJson(),
	}, nil
}
