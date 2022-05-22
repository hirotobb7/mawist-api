package main

import (
	"log"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-cmp/cmp"

	"github.com/hiroto7/mawist/db/seeders"
	"github.com/hiroto7/mawist/pkg/json"
)

func TestMain(m *testing.M) {
	if err := seeders.CreateWishLists(); err != nil {
		if err := seeders.DeleteWishLists(); err != nil {
			log.Printf("seed rollback error %+v\n", err)
		}
		log.Fatalf("seed create error %+v\n", err)
	}

	m.Run()
}

func TestHandler(t *testing.T) {
	t.Cleanup(func() {
		if err := seeders.DeleteWishLists(); err != nil {
			t.Errorf("seed clean up error %+v\n", err)

		}
	})

	request := events.APIGatewayProxyRequest{
		Path: "/wish-lists",
	}

	t.Run("400 Response", func(t *testing.T) {
		type ExpectBody []string

		expectBody := ExpectBody{
			"UserIdは必須です。",
		}

		request.Body, _ = json.Stringify(map[string]string{})

		result, _ := handler(request)

		resultBody := ExpectBody{}
		if err := json.Parse(result.Body, &resultBody); err != nil {
			t.Fatalf("unexpected error: %+v\n", err)
		}

		if diff := cmp.Diff(400, result.StatusCode); diff != "" {
			t.Errorf("(-expect +result):\n%s", diff)
		}

		if diff := cmp.Diff(expectBody, resultBody); diff != "" {
			t.Errorf("(-expect +result):\n%s", diff)
		}
	})

	t.Run("200 Fill Response", func(t *testing.T) {
		type ExpectBody [2]map[string]interface{}

		expectBody := ExpectBody{
			{
				"userId":     "test-user-id-1",
				"id":         "test-id-1",
				"name":       "マイリスト",
				"createdAt":  "2022-05-08T17:00:00Z",
				"updatedAt":  "2022-05-08T17:00:00Z",
				"isDisabled": false,
			},
			{
				"userId":     "test-user-id-1",
				"id":         "test-id-2",
				"name":       "プレゼントリスト",
				"createdAt":  "2022-05-10T09:00:00Z",
				"updatedAt":  "2022-05-10T09:00:00Z",
				"isDisabled": false,
			},
		}

		request.Body, _ = json.Stringify(map[string]string{
			"userId": "test-user-id-1",
		})

		result, _ := handler(request)

		resultBody := ExpectBody{}
		if err := json.Parse(result.Body, &resultBody); err != nil {
			t.Fatalf("unexpected error: %+v\n", err)
		}

		if diff := cmp.Diff(200, result.StatusCode); diff != "" {
			t.Errorf("(-expect +result):\n%s", diff)
		}

		if diff := cmp.Diff(expectBody, resultBody); diff != "" {
			t.Errorf("(-expect +result):\n%s", diff)
		}
	})

	t.Run("200 Empty Response", func(t *testing.T) {
		request.Body, _ = json.Stringify(map[string]string{
			"userId": "not-existed-id",
		})

		result, _ := handler(request)

		var resultBody interface{}
		if err := json.Parse(result.Body, &resultBody); err != nil {
			t.Fatalf("unexpected error: %+v\n", err)
		}

		if diff := cmp.Diff(200, result.StatusCode); diff != "" {
			t.Errorf("(-expect +result):\n%s", diff)
		}

		if diff := cmp.Diff(resultBody, nil); diff != "" {
			t.Errorf("(-expect +result):\n%s", diff)
		}
	})
}
