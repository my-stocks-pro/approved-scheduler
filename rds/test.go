package rds

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"fmt"
)

func Set(conn *redis.Client, key string, val interface{}) {
	b, e := json.Marshal(val)
	if e != nil {
		panic(e)
	}
	err := conn.Set(key, b, 0).Err()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Test() {
	rds := NewRDS()
	conn := rds.NewConn("1")

	date := "2018-06-25"
	listids := []string{"999999", "777777", "888888"}

	Set(conn, "date", date)

	Set(conn, "listids", listids)

}
