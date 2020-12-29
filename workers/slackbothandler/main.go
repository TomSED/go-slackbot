package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/TomSED/go-slackbot"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	sdklambda "github.com/aws/aws-sdk-go/service/lambda"
)

func main() {
	lambda.Start(slackBotHandler)
}

func slackBotHandler(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// For debugging
	requestBytes, _ := json.MarshalIndent(e, "", "    ")
	fmt.Println(string(requestBytes))

	formValues, err := url.ParseQuery(e.Body)
	if err != nil {
		fmt.Printf("url.ParseQuery error: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusForbidden}, nil
	}

	// Check request token
	requestToken := formValues.Get("token")
	if !authenticate(requestToken) {
		fmt.Printf("invalid request token: %v\n", requestToken)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusForbidden}, nil
	}

	// Transform slack command to payload for worker function
	event := slackbot.LambdaWorkerEvent{
		Message:     formValues.Get("text"),
		ResponseURL: formValues.Get("response_url"),
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("json.Marshal error: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, fmt.Errorf("json.Marshal error: %v", err)
	}

	// Execute worker function
	sess := session.Must(session.NewSession())
	lambdaSdk := sdklambda.New(sess)
	_, err = lambdaSdk.InvokeWithContext(context.Background(), &sdklambda.InvokeInput{
		FunctionName:   aws.String(os.Getenv("SLACKBOT_WORKER_FUNCTION_NAME")),
		InvocationType: aws.String("Event"),
		Payload:        []byte(eventBytes),
	})
	if err != nil {
		fmt.Printf("lambdaSdk.InvokeWithContext error: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, fmt.Errorf("lambdaSdk.InvokeWithContext error: %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Success",
	}, nil
}

func authenticate(requestToken string) bool {

	token := os.Getenv("SLACKBOT_AUTH_TOKEN")

	return requestToken == token
}
