package slackbot

// LambdaWorkerEvent is for Invoking another lambda directly from a lambda
type LambdaWorkerEvent struct {
	Message     string `json:"message,omitempty"`
	ResponseURL string `json:"responseURL,omitempty"`
	ChannelID   string `json:"channelID,omitempty"`
	UserID      string `json:"userID,omitempty"`
}
