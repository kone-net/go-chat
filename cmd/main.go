package main

import (
	"chat-room/global/log"
	"chat-room/config"
	"chat-room/router"
	"chat-room/server"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	log.InitLogger(config.GetConfig().Log.Path, config.GetConfig().Log.Level)
	log.Info("config", zap.Any("config", config.GetConfig()))

	log.Info("start server", zap.String("start", "start web sever..."))

	newRouter := router.NewRouter()

	go server.MyServer.Start()

	s := &http.Server{
		Addr:           "127.0.0.1:8888",
		Handler:        newRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if nil != err {
		log.Error("server error", zap.Any("serverError", err))
	}
}
