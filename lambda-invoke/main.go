package main

import (
	"log"
	"fmt"
	"os"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	l "github.com/aws/aws-lambda-go/lambda"	// importでlambdaが２つある為、'l'で呼べるよう指定
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-lambda-go/events"
)

var svc *lambda.Lambda
var lambdaArn string

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// 別のLambdaに渡すための構造体定義
type Event struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// メッセージを送信する
func LambdaInvoke() ([]byte, error) {
	// 呼び出したLambdaに渡すデータ
	payload := Event{
		Id:   1,
		Name: "yukpiz",
	}
	jsonBytes, _ := json.Marshal(payload)

	input := &lambda.InvokeInput{
			FunctionName:   aws.String(lambdaArn),
			Payload:        jsonBytes,
			// 非同期で呼ぶ場合はInvocationTypeを設定する
			// InvocationType: aws.String("Event"),
	}

	resp, _ := svc.Invoke(input)
	fmt.Println(resp)

	return resp.Payload, nil
}

func handler() (events.APIGatewayProxyResponse, error) {
	// 環境変数の取得
	lambdaArn = os.Getenv("AWS_LAMBDA_TRIGGER_FUNCTION")
	fmt.Println(lambdaArn)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
	}))
	svc = lambda.New(sess)
	resp, err := LambdaInvoke()
	if err != nil {
			log.Fatal(err)
	}
	fmt.Println(resp)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", resp),
		StatusCode: 200,
	}, nil
}

func main() {
	l.Start(handler)
}
