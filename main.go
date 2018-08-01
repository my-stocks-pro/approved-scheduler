package main

import (
	"github.com/my-stocks-pro/approved-scheduler/scheduler"
	"net/http"
	"log"
	"os/signal"
	"os"
	"context"
	"time"
)

func main() {

	//TODO -> This is test
	//rds.Test()

	//router := gin.Default()
	//
	//Scheduler := scheduler.New()
	//
	//router.GET(Scheduler.Config.Service, func(c *gin.Context) {
	//	go Scheduler.Run()
	//})
	//
	//router.GET(fmt.Sprintf("%s/stop", Scheduler.Config.Service), func(c *gin.Context) {
	//	if err := router.Shutdown(nil); err != nil {
	//		panic(err) // failure/timeout shutting down the server gracefully
	//	}
	//})
	//
	//router.Run(":8002")

	Scheduler := scheduler.New()

	go Scheduler.Routing()

	go func() {
		if err := Scheduler.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	signal.Notify(Scheduler.QuitOS, os.Interrupt)
	select {
	case <-Scheduler.QuitOS:
		log.Println("Shutdown Server by OS signal...")
	case <-Scheduler.QuitRPC:
		log.Println("Shutdown Server by RPC signal...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := Scheduler.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")


}
