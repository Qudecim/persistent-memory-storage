package main

import (
	"fmt"
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

	k, ok := db.Get("test_key99")
	fmt.Println(k, ok)

	k, ok = db.Get("test_key199")
	fmt.Println(k, ok)

	k, ok = db.Get("test_key299")
	fmt.Println(k, ok)

	db.Snapshot()

}
