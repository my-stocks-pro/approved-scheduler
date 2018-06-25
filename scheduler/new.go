package scheduler

type TypeReadisApproved struct {
	Date   string
	IDlist []string
}

type TypeConfig struct {
	Tick    uint64
	URLGET  string
	URLPOST string
	RadisDB string
}

type TypeSchaduler struct {
	Config    *TypeConfig
	RedisData *TypeReadisApproved
}

func ReadisApprovedNew() *TypeReadisApproved {
	return &TypeReadisApproved{}
}

func New() *TypeSchaduler {
	return &TypeSchaduler{
		GetConfig(),
		ReadisApprovedNew(),
	}
}
