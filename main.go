package main

import (
	"flag"
	"qudecim/db/appConfig"
	"qudecim/db/db"
	socket "qudecim/db/transport"
)

// TODO:
// expired_time
// logs and errors
// beauty

func main() {

	flagProfile := flag.Bool("profile", false, "Is profile action")

	config, err := appConfig.LoadConfig("config.yaml")
	if err != nil {
		return
	}

	db.Init(config)

	binlog := db.NewBinlog(config.Binlog.Directory, config.Binlog.EveryCheckOversize, config.Binlog.ChanceCheckOversize)
	go binlog.Run()
	db.GlobalBinlog = binlog

	socket.Run()

	if *flagProfile {
		profile()
	}

	//db.Snapshot()
}
