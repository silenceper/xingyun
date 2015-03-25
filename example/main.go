package main

import (
	"net/http"

	"github.com/xiaoenai/xingyun"
)

func main() {
	cfg := &xingyun.Config{}
	server := xingyun.NewServer(cfg)
	logger := server.Logger()

	pipe := server.NewPipe("test", xingyun.PipeHandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		logger.Tracef("before")
		next(w, r)
		logger.Tracef("after")
	}))

	server.Get("/hello", pipe.Wrap(func(ctx *xingyun.Context) {
		logger.Tracef("hello world")
		ctx.WriteString("hello world")
	}))

	err := server.ListenAndServe("127.0.0.1:9000")
	logger.Errorf("%s", err)
}
