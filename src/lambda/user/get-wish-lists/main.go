// TODO: 認証

package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/hirotobb7/mawist/internal/db/repository/dynamo"
	"github.com/hirotobb7/mawist/internal/db/service"
	"github.com/hirotobb7/mawist/pkg/json"
	"github.com/hirotobb7/mawist/pkg/log"
	"github.com/hirotobb7/mawist/pkg/response"
	"github.com/hirotobb7/mawist/pkg/validator"
)

type RequestBody struct {
	UserId string `json:"userId" validate:"required"`
}

var logger = log.GetLogger()
var db = dynamo.GetDb()
var wishListService = service.NewWishListService(dynamo.NewWishListRepository(db))

func handleRequest(request events.APIGatewayProxyRequest) (int, interface{}, error) {
	var requestBody RequestBody
	if err := json.Parse(request.Body, &requestBody); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	validationMessages := validator.Validate(&requestBody).GetMessages()
	if len(validationMessages) != 0 {
		return http.StatusBadRequest, validationMessages, nil
	}

	wishLists, err := wishListService.FindByUserId(requestBody.UserId)
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

	logger.Info.Printf("%d", statusCode)
	logger.Info.Printf("%+v\n", result)

	apiGatewayProxyResponse, err := response.New(statusCode, result)

	if err != nil {
		logger.Error.Printf("%+v\n", err)
	}

	return apiGatewayProxyResponse, nil
}

func main() {
	lambda.Start(handler)
}
