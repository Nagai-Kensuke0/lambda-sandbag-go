# lambda-sandbag-go

## リポジトリの説明

- AWS Lambda を中心に、色々な AWS サービス（無料または極めて安価なもの）を使ってみようという為の検証用リポジトリです。
- 基本的に、以下の２種類の Lambda を実装しています
  1. Lambda トリガーとなる AWS サービスを動かす Lambda
  2. 1 で動かした AWS サービスをトリガーとして発火する Lambda
- Lambda やその他の AWS サービスの構築は template.yaml にテンプレートを作成し、AWS SAM でデプロイしています

## issues

- template.yaml が１ファイルに全てのテンプレートがまとまっているので、整理を検討する
- 同じサービスに関連する Lambda はまとめてフォルダを切った方がわかりやすいと思うので、整理を検討する

## 使用言語

- Go 1.17

## 各フォルダ配下の Lambda 関数の説明

- lambda-invoke/main.go
- lambda-trigger/main.go
  lambda-invoke/main.go をトリガーとして lambda-invoke/main.go の Lambda を発火させる

- s3-notification/main.go
  特定の Lambdab バゲットにファイルを置くことで発火する関数

- sns-notification/main.go
- sns-send-message/main.go

  1. sns-send-message/main.go で Amazon SNS（以下 SNS） にメッセージを送る
  2. SNS は sns-notification/main.go の Lambda を発火させる

- sqs-notification/main.go
- sqs-send-message/main.go
  1. sqs-send-message/main.go で Amazon SQS（以下 SQS） にメッセージを送る
  2. SQS は sqs-notification/main.go の Lambda を発火させる
