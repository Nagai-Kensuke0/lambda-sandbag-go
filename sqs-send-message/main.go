package main

import (
	"log"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws/session"
)

var svc *sqs.SQS
var region string

// メッセージを送信する
func SendMessage() error {
	region = os.Getenv("AWS_REGION")
	queueName := os.Getenv("AWS_SQS_QUEUENAME")
	accountNumber := os.Getenv("AWS_ACCOUNT_NUMBER")

	queueUrl := "https://sqs." + region + ".amazonaws.com/" + accountNumber + "/" + queueName
	fmt.Println("queueUrl", queueUrl)

	// 送信内容を作成
	params := &sqs.SendMessageInput{
			MessageBody:  aws.String("HogeFugaPiyo"),
			QueueUrl:     aws.String(queueUrl),
			DelaySeconds: aws.Int64(1),
	}

	sqsRes, err := svc.SendMessage(params)
	if err != nil {
			return err
	}

	fmt.Println("SQSMessageID", *sqsRes.MessageId)

	return nil
}

func handler(s3Event events.S3Event) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
	}))
	svc = sqs.New(sess)
	if err := SendMessage(); err != nil {
			log.Fatal(err)
	}
}

func main() {
	lambda.Start(handler)
}
