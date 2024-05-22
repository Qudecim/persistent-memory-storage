package main

import (
	"flag"
	"qudecim/db/appConfig"
	"qudecim/db/db"
	"qudecim/db/internal/binlog"
	socket "qudecim/db/transport"
)

// TODO:
// expired_time
// logs and errors
// beauty
// snapshot to internal

func main() {

	flagProfile := flag.Bool("profile", false, "Is profile action")

	config, err := appConfig.LoadConfig("config.yaml")
	if err != nil {
		return
	}

	binlogReader := binlog.NewBinlogReader(config.Binlog.Directory, 100)
	binlogReader.Init()
	for binlogReader.Next() {
		// set to db
	}

	if err := binlogReader.Err(); err != nil {
		return
	}

	db.Init(config)

	binlog := binlog.NewBinlogWriter(config.Binlog.Directory, 100)
	go binlog.Run()
	db.GlobalBinlog = binlog

	socket.Run()

	if *flagProfile {
		profile()
	}

	//db.Snapshot()
}
