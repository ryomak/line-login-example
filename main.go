package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/ryomak/line-login-example/application/usecase"
	"github.com/ryomak/line-login-example/handler"
	"github.com/ryomak/line-login-example/pkg/line"
)

const (
	redirectURL = "http://localhost:8080/redirect"
)

func main() {

	client, err := line.NewLine()
	if err != nil {
		log.Fatalln(err)
	}
	lineHandler := &handler.LineHandler{
		Usecase: usecase.NewLineUsecase(client),
		Client:  client,
	}

	r := chi.NewRouter()
	r.Route("/hook", func(r chi.Router) {
		r.Post("/", lineHandler.WebHook)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))

}
