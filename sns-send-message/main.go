package main

import (
	"log"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var svc *sns.SNS
var region string

// メッセージを送信する
func SendMessage() error {
	region = os.Getenv("AWS_REGION")
	accountNumber := os.Getenv("AWS_ACCOUNT_NUMBER")
	snsTopicName := os.Getenv("AWS_SNS_TOPICNAME")

	topicArn  := "arn:aws:sns:" + region + ":" + accountNumber + ":" + snsTopicName

	fmt.Println("%s", topicArn)

	message := "This is sample message."

	// メッセージを送信するための構造体を作成
	inputPublish := &sns.PublishInput{
			Message:  aws.String(message),
			TopicArn: aws.String(topicArn),
	}

	// メッセージの送信(Publish)
	MessageId, err := svc.Publish(inputPublish)
	if err != nil {
			fmt.Println("Publish Error: ", err)
	}

	fmt.Println(MessageId)

	return nil
}

func handler() {
	sess := session.Must(session.NewSession())
	svc = sns.New(sess)
	if err := SendMessage(); err != nil {
			log.Fatal(err)
	}
}

func main() {
	lambda.Start(handler)
}
