package scheduler

import (
	"github.com/my-stocks-pro/approved-scheduler/config"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
)

type TypeReadisApproved struct {
	Date   string
	IDlist []string
}


type TypeScheduler struct {
	QuitOS    chan os.Signal
	QuitRPC   chan bool
	Router    *gin.Engine
	Server    *http.Server
	Config    *config.TypeConfig
	RedisData *TypeReadisApproved
}

func ReadisApprovedNew() *TypeReadisApproved {
	return &TypeReadisApproved{}
}

func New() *TypeScheduler {
	router := gin.Default()
	return &TypeScheduler{
		QuitOS:  make(chan os.Signal),
		QuitRPC: make(chan bool),
		Router:  router,
		Server: &http.Server{
			Addr:    ":8002",
			Handler: router,
		},
		Config:    config.GetConfig(),
		RedisData: ReadisApprovedNew(),
	}
}
