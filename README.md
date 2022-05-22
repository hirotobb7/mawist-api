# mawist-api
ほしい物リスト管理API

## 使用技術
- Docker
- AWS SAM
- AWS CloudFormation
- AWS DynamoDB
- Golang

## 開発ブランチ名
GitHubで作成したissueの番号を用いる。
`feature/${issue-number}`とする。


## ローカル開発環境

### AWSアカウントの作成
AWSを利用したサーバレスアプリのため、AWSアカウントを作成すること

### `.env`の作成
- `./.sam.env.example.json`を参考し、`./.sam.env.json`を作成する  
  sam build用の環境変数設定
- `./.docker.env.example`を参考し、`./.docker.env`を作成する  
  go-samコンテナ用環境変数の設定

### Docker containerの作成
```bash
pwd # /{your project path}/mawist/api
docker compose build
docker compose up -d
docker compose exec go-sam /bin/bash
```

### sam build
- go-samコンテナ内で実行すること
- 環境変数の設定を行なっていること

```bash
pwd # /var/sam
sam build
```

### APIエンドポイントの作成
- `sam build`実行済みであること
```bash
sam local start-api --container-host host.docker.internal --host 0.0.0.0
```

ホストマシンから、http://127.0.0.1:3000 でアクセス可能になる。  
細かいオプションについては、[API Gateway のローカルでの実行](https://docs.aws.amazon.com/ja_jp/serverless-application-model/latest/developerguide/serverless-sam-cli-using-start-api.html)を参考


## DB

### ローカルDB
テーブルの一括作成

- go-samコンテナ内で実行すること
- 環境変数の設定を行なっていること
```bash
pwd # /var/sam
bash ./scripts/create-tables.sh
```
