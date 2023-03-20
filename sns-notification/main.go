package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
			log.Printf("[%s] EventSubscriptionArn = %s \n", record.EventSource, record.EventSubscriptionArn)
	}
}

func main() {
	lambda.Start(handler)
}
