package response

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/hiroto7/mawist/pkg/json"
)

func InternalServerError() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       "エラーが発生しました。",
	}
}

func New(statusCode int, result interface{}) (events.APIGatewayProxyResponse, error) {
	if statusCode == http.StatusInternalServerError {
		return InternalServerError(), nil
	}

	responseBody, err := json.Stringify(result)
	if err != nil {
		return InternalServerError(), err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       responseBody,
	}, nil
}
