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
	if event == nil {
		return nil, fmt.Errorf("recieved nil event")
	}

	message := fmt.Sprintf("Hello, %s; We are delighted to have you in Version 1.1.2", event.Name)
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
