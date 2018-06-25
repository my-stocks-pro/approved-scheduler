package main

import (
	"github.com/my-stocks-pro/approved-scheduler/scheduler"
)

func main() {

	Scheduler := scheduler.New()

	Scheduler.Run()
}
