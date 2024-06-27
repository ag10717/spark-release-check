package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	message := fmt.Sprintf("Hello, World! Everyone poses; this the FRESHEST! result.")
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
