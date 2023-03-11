package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(s3Event events.S3Event) {
	for _, record := range s3Event.Records {
			s3rec := record.S3
			log.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3rec.Bucket.Name, s3rec.Object.Key)
	}
}

func main() {
	lambda.Start(handler)
}
