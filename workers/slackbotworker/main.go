package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/TomSED/go-slackbot"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(slackBotWorker)
}

func slackBotWorker(ctx context.Context, e slackbot.LambdaWorkerEvent) error {
	// For debugging
	requestBytes, _ := json.MarshalIndent(e, "", "    ")
	fmt.Println(string(requestBytes))

	// Send response to slack
	delayedResponse := fmt.Sprintf("Finish Long running task")
	err := slackbot.SendSlackMessage(delayedResponse, e.ResponseURL)
	if err != nil {
		fmt.Printf("slackbot.SendSlackMessage error: %v", err)
		return nil
	}

	return nil
}
