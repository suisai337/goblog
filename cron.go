package main

//import (
//	"github.com/EDDYCJY/go-gin-example/models"
//	"github.com/robfig/cron"
//	"log"
//	"time"
//)
//
//func main() {
//	log.Println("starting....")
//
//	c := cron.New()
//
//	c.AddFunc("* * * * * *", func() {
//		log.Println("Run models.CleanAllTag...")
//		models.CleanAllTag()
//	})
//	c.AddFunc("* * * * * *", func() {
//		log.Println("Run models.CleanAllArticle...")
//		models.CleanAllArticle()
//	})
//
//	c.Start()
//
//	t := time.NewTimer(time.Second * 10)
//	for {
//		select {
//			case <-t.C:
//				t.Reset(time.Second * 10)
//		}
//	}

//}
