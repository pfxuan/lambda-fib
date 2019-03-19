package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"time"
)

type MyEvent struct {
	Number uint64 `json:"number"`
}

type MyResponse struct {
	ID          string `json:"id:"`
	ComputeTime int64  `json:"computeTime:"`
}

func fib(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func HandleRequest(ctx context.Context, event MyEvent) (MyResponse, error) {
	startTime := time.Now()
	uuid := uuid.New().String()
	fib(event.Number)
	duration := time.Since(startTime)
	return MyResponse{ID: uuid, ComputeTime: duration.Nanoseconds()}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
