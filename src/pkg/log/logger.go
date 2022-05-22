package log

import (
	"log"
	"os"
)

type logger struct {
	UserId string
	Info   *log.Logger
	Error  *log.Logger
}

// NOTICE: シングルスレッドのみ
var singletonLogger *logger

func GetLogger() *logger {
	if singletonLogger == nil {
		log.Print("create")
		singletonLogger = &logger{
			Info: log.New(
				os.Stdout,
				"[INFO]",
				log.Llongfile, //　時刻, リクエストとの関連はLambdaに委譲
			),
			Error: log.New(
				os.Stderr,
				"[ERROR]",
				log.Llongfile, //　時刻, リクエストとの関連はLambdaに委譲
			),
		}
	}

	return singletonLogger
}
