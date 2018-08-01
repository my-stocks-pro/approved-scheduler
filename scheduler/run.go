package scheduler

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)


func (s *TypeScheduler) ApprovedTask() {

	//TODO pass Redis DB name
	data, errGET := s.HandleGET()
	if errGET != nil {
		fmt.Println(errGET)
	}
	fmt.Println("HandleGET", string(data))

	//data := Test()
	//
	//resp, errPOST := s.HandlePOST(data)
	//if errPOST != nil {
	//	fmt.Println(errPOST)
	//}
	//
	//fmt.Println(resp)

	//os.Exit(1)
}


func (s *TypeScheduler) Run() {

	//TODO NEED use right time of schedule -> must change in conf

	fmt.Println("NewScheduler ApprovedTask Create")

	approvedScheduler := gocron.NewScheduler()
	approvedScheduler.Every(s.Config.Tick).Seconds().Do(s.ApprovedTask)
	<- approvedScheduler.Start()
}
