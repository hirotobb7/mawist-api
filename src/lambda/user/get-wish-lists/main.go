package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/hiroto7/mawist/internal/dynamo"
	"github.com/hiroto7/mawist/pkg/json"
	"github.com/hiroto7/mawist/pkg/log"
	"github.com/hiroto7/mawist/pkg/response"
	"github.com/hiroto7/mawist/pkg/validator"
)

var logger = log.GetLogger()

type RequestBody struct {
	UserId string `json:"userId" validate:"required"`
}

func handleRequest(request events.APIGatewayProxyRequest) (int, interface{}, error) {

	var reqBody RequestBody
	if err := json.Parse(request.Body, &reqBody); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	validationMessages := validator.Validate(&reqBody).GetMessages()
	if len(validationMessages) != 0 {
		return http.StatusBadRequest, validationMessages, nil
	}

	wishLists, err := dynamo.FindWishListsByUserId(reqBody.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, wishLists, nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger.Info.Printf("request.body: %+v\n", request.Body)

	statusCode, result, err := handleRequest(request)
	if err != nil {
		logger.Error.Printf("%+v\n", err)
	}

	// rename
	res, err := response.New(statusCode, result)

	logger.Info.Printf("%+d", statusCode)
	logger.Info.Printf("%+v\n", result)

	if err != nil {
		logger.Error.Printf("%+v\n", err)
	}

	return res, nil
}

func main() {
	lambda.Start(handler)
}
