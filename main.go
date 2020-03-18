package main

import (
	"context"
	"fmt"
	"github.com/suisai337/goblog/models"
	"github.com/suisai337/goblog/pkg/gredis"
	"github.com/suisai337/goblog/pkg/logging"
	"github.com/suisai337/goblog/pkg/setting"
	"github.com/suisai337/goblog/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	err := gredis.Setup()
	if err != nil {
		logging.Error(err)
	}

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = s.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
