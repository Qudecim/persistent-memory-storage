package db

import (
	"fmt"
	"os"
	"qudecim/db/appConfig"
	"qudecim/db/dto"
	"qudecim/db/internal/binlog"
	"strconv"
	"sync"
)

var Data map[string]string

var GlobalBinlog *binlog.BinlogWriter

var Config *appConfig.Config

var rw sync.RWMutex
var Wg sync.WaitGroup

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

}

func Set(request *dto.Request) {
	rw.Lock()
	Data[request.GetKey()] = request.GetValue()
	rw.Unlock()

	GlobalBinlog.Add(request)
}

func Get(request *dto.Request) (string, bool) {
	rw.RLock()
	value, ok := Data[request.GetKey()]
	rw.RUnlock()
	return value, ok
}

func Snapshot() {
	saveSnapshot()
}
