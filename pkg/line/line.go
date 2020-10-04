package line

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Client struct {
	*linebot.Client
}

func NewLine() (*Client, error) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		return nil, err
	}
	return &Client{bot}, nil
}

func (l *Client) ParseRequest(req *http.Request) ([]*Event, error) {
	events, err := l.Client.ParseRequest(req)
	if err != nil {
		return nil, err
	}
	return toEvent(events), nil
}

func (l *Client) ReplyMessage(ctx context.Context, token, message string) error {
	res, err := l.Client.ReplyMessage(token, linebot.NewTextMessage(message)).WithContext(ctx).Do()
	log.Println(res)
	return err
}

func (l *Client) ReplyImage(ctx context.Context, token, image string) error {
	res, err := l.Client.ReplyMessage(token, linebot.NewImageMessage(image, image)).WithContext(ctx).Do()
	log.Println(res)
	return err
}
