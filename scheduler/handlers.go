package scheduler

import (
	"net/http"
	"io/ioutil"
	"bytes"
)

func (s *TypeScheduler) HandleGET() ([]byte, error) {

	req, errReq := http.NewRequest("GET", s.Config.MethodGET, nil)
	if errReq != nil {
		return nil, errReq
	}

	q := req.URL.Query()
	q.Add("namedb", s.Config.RedisDB)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	b, errReadAll := ioutil.ReadAll(resp.Body)
	if errReadAll != nil {
		return nil, errReadAll
	}

	return b, nil
}

func (s *TypeScheduler) HandlePOST (data []byte) (*http.Response, error) {

	req, errReq := http.NewRequest("POST", s.Config.MethodPOST, bytes.NewReader(data))
	if errReq != nil {
		return nil, errReq
	}

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	return resp, nil
}
