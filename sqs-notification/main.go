package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(sqsEvent events.SQSEvent) {
	for _, record := range sqsEvent.Records {
			log.Printf("[%s] EventSourceArn = %s \n", record.EventSource, record.EventSourceARN)
	}
}

func main() {
	lambda.Start(handler)
}
