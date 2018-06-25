package scheduler

type TypeReadisData struct {
	Date string
	IDlist []string
}

type TypeConfig struct {
	Tick     uint64
	URLGET    string
	URLPOST   string
}

type TypeSchaduler struct {
	Config    *TypeConfig
	RedisData *TypeReadisData
}


func ReadisDataNew () *TypeReadisData{
	return &TypeReadisData{}
}


func New() *TypeSchaduler {
	return &TypeSchaduler{
		GetConfig(),
		ReadisDataNew(),
	}
}
