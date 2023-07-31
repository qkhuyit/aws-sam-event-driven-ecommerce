package cognito

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func NewCognitoIdentityProvider() *cognitoidentityprovider.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	cip := cognitoidentityprovider.NewFromConfig(cfg)
	return cip
}
