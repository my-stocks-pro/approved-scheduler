package main

import (
		"github.com/my-stocks-pro/approved-scheduler/scheduler"
	"github.com/gin-gonic/gin"
)

func main() {

	//TODO -> This is test
	//rds.Test()

	router := gin.Default()

	Scheduler := scheduler.New()

	go Scheduler.Run()

	router.Run(":8002")
}
