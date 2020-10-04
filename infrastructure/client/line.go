package client

import (
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Line struct{ *linebot.Client }

func NewLineClient() (*Line, error) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		return nil, err
	}
	return &Line{bot}, nil
}
