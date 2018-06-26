package main

import (
	"github.com/my-stocks-pro/approved-scheduler/scheduler"
)

func main() {

	//TODO -> This is test
	//rds.Test()

	Scheduler := scheduler.New()

	Scheduler.Run()
}
