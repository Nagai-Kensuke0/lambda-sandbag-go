package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// 呼び出し元のLambdaから受け取るための構造体定義
type Event struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func handler(event Event) (Response, error) {
	fmt.Println(event)
	time.Sleep(time.Second * 2)
	return Response{
			Message: "success",
			Ok:      true,
	}, nil
}

func main() {
	lambda.Start(handler)
}