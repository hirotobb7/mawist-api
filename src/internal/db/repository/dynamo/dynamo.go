package dynamo

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// NOTICE: シングルスレッドのみ
var singletonDb *dynamo.DB

func GetDb() *dynamo.DB {
	if singletonDb == nil {
		singletonDb = dynamo.New(
			session.Must(session.NewSession()),
			&aws.Config{
				Region:   aws.String(os.Getenv("AWS_REGION")),
				Endpoint: aws.String(os.Getenv("DB_ENDPOINT")),
			})
	}

	return singletonDb
}
