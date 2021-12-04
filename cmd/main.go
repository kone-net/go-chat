package main

import (
	"chat-room/config"
	"chat-room/internal/kafka"
	"chat-room/internal/router"
	"chat-room/internal/server"
	"chat-room/pkg/common/constant"
	"chat-room/pkg/global/log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func main() {
	log.InitLogger(config.GetConfig().Log.Path, config.GetConfig().Log.Level)
	log.Info("config", zap.Any("config", config.GetConfig()))

	if config.GetConfig().MsgChannelType.ChannelType == constant.KAFKA {
		kafka.InitProducer(config.GetConfig().MsgChannelType.KafkaTopic, config.GetConfig().MsgChannelType.KafkaHosts)
		kafka.InitConsumer(config.GetConfig().MsgChannelType.KafkaHosts)
		go kafka.ConsumerMsg(server.ConsumerKafkaMsg)
	}

	log.Info("start server", zap.String("start", "start web sever..."))

	newRouter := router.NewRouter()

	go server.MyServer.Start()

	s := &http.Server{
		Addr:           ":8888",
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
