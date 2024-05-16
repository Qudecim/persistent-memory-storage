package db

import (
	"fmt"
	"os"
	"qudecim/db/appConfig"
	"strconv"
)

var Data map[string]string
var CurrentBinlog string
var CurrentBinlogSource *os.File

var Config *appConfig.Config

func Init(appConfig *appConfig.Config) {
	Data = make(map[string]string)

	Config = appConfig

	err := os.MkdirAll(Config.Binlog.Directory, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	err = os.MkdirAll(Config.Snapshot.Directory, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	lastSnapshot := getLastSnapshot()
	if lastSnapshot > 0 {
		readSnapshot(strconv.Itoa(lastSnapshot))
	}

	binlogs := getBinlogs(lastSnapshot)
	lastbinlog := 0
	for _, binlog := range binlogs {
		readBinlog(strconv.Itoa(binlog))
		lastbinlog = max(lastbinlog, binlog)
	}

	if lastbinlog > 0 {
		changeBinlog(strconv.Itoa(lastbinlog))
	} else {
		changeBinlog(timestamp())
	}
}

func Set(key string, value string) {
	Data[key] = value
	addToBinlog(key, value)
}

func Get(key string) (string, bool) {
	value, ok := Data[key]
	return value, ok
}

func Snapshot() {
	saveSnapshot()
}
