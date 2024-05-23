package app

import (
	"fmt"
	"os"
	"qudecim/db/appConfig"
	"qudecim/db/dto"
	"sync"
)

type App struct {
	data   map[string]string
	binlog *BinlogWriter
	config *appConfig.Config

	rw sync.RWMutex
	Wg sync.WaitGroup
}

func NewApp(binlog *BinlogWriter, config *appConfig.Config) *App {
	return &App{
		data:   make(map[string]string),
		binlog: binlog,
		config: config,
	}
}

func (a *App) Init() {
	err := os.MkdirAll(a.config.Binlog.Directory, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	err = os.MkdirAll(a.config.Snapshot.Directory, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
}

func (a *App) Set(request *dto.Request) {
	a.rw.Lock()
	a.data[request.GetKey()] = request.GetValue()
	a.rw.Unlock()

	a.binlog.Add(request)
}

func (a *App) Get(request *dto.Request) (string, bool) {
	a.rw.RLock()
	value, ok := a.data[request.GetKey()]
	a.rw.RUnlock()
	return value, ok
}

func (a *App) ForceSet(key string, value string) {
	a.data[key] = value
}
