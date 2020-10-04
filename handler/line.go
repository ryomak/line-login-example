package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ryomak/line-login-example/application/usecase"
	"github.com/ryomak/line-login-example/pkg/line"
)

type LineHandler struct {
	Usecase *usecase.LineUsecase
	Client  *line.Client
}

func (l *LineHandler) WebHook(w http.ResponseWriter, r *http.Request) {
	events, _ := l.Client.ParseRequest(r)

	ctx := r.Context()
	es := []error{}
	for i, v := range events {
		fmt.Println(i, v)
		switch v.Type {
		case "join":
			if err := l.Usecase.Join(ctx, v); err != nil {
				es = append(es, err)
			}
		default:
			if err := l.Usecase.Reply(ctx, v); err != nil {
				es = append(es, err)
			}
		}
	}
	if len(es) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(es)
		return
	}
	w.WriteHeader(http.StatusOK)

}
