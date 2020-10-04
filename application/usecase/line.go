package usecase

import (
	"context"
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/ryomak/line-login-example/pkg/line"
)

func NewLineUsecase(l *line.Client) *LineUsecase {
	return &LineUsecase{lc: l}

}

type LineUsecase struct {
	lc *line.Client
}

func (l *LineUsecase) Join(ctx context.Context, event *line.Event) error {
	profile, err := l.lc.Client.GetProfile(event.Source.UserID).WithContext(ctx).Do()
	if err != nil {
		return err
	}

	return l.lc.ReplyMessage(ctx, event.ReplyToken, fmt.Sprintf("%sこんにちは!", profile.DisplayName))
}

func (l *LineUsecase) Reply(ctx context.Context, event *line.Event) error {
	/*
		profile, err := l.lc.Client.GetProfile(event.Source.UserID).WithContext(ctx).Do()
		if err != nil {
			return err
		}
	*/
	fmt.Println(event.Message)
	fmt.Print("返信:")
	s := "aaaa:"
	fmt.Scanln(&s)
	_, err := l.lc.Client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("%s", s))).Do()
	return err
}
