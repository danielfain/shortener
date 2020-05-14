package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// RequestBody represents the url to be shortened from the request
type RequestBody struct {
	URL string `json:"url"`
}

// byte represents ASCII characters; use rune for unicode support
var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numLetters = 6

// Handler for the lambda function
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := RequestBody{}

	err := json.Unmarshal([]byte(request.Body), &body)

	if err != nil {
		println(err)
	}

	shortURL := os.Getenv("host") + randomLetters(numLetters)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		Body:       shortURL,
	}, nil

}

func randomLetters(n int) string {
	seq := make([]byte, n)
	for i := range seq {
		seq[i] = letters[rand.Intn(len(letters))]
	}
	return string(seq)
}

func main() {
	lambda.Start(Handler)
}