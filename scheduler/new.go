package scheduler

import "github.com/my-stocks-pro/approved-scheduler/config"

type TypeReadisApproved struct {
	Date   string
	IDlist []string
}


type TypeSchaduler struct {
	Config    *config.TypeConfig
	RedisData *TypeReadisApproved
}

func ReadisApprovedNew() *TypeReadisApproved {
	return &TypeReadisApproved{}
}

func New() *TypeSchaduler {
	return &TypeSchaduler{
		config.GetConfig(),
		ReadisApprovedNew(),
	}
}
