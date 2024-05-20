package main

import (
	"qudecim/db/appConfig"
	"qudecim/db/db"
)

// TODO:
// socket
// expired_time

func main() {
	config, err := appConfig.LoadConfig("config.yaml")
	if err != nil {
		return
	}

	db.Init(config)

	// for i := 200; i < 300; i++ {
	// 	db.Set("test_key"+strconv.Itoa(i), "test_value"+strconv.Itoa(i))
	// }

	//profile()

	db.Snapshot()

}
