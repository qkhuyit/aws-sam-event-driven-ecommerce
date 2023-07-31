package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/handlers"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		internal.RootModule,
		fx.Invoke(func(h handlers.ProductHandler) {
			lambda.Start(h.Delete)
		}),
	).Run()
}
