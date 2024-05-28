package app

import (
	"fmt"
	"os"
	"qudecim/db/appConfig"
	"qudecim/db/internal/dto"
	"sync"
)

type App struct {
	data   map[string]Item
	binlog *BinlogWriter
	config *appConfig.Config

	rw sync.RWMutex
	Wg sync.WaitGroup
}

func NewApp(binlog *BinlogWriter, config *appConfig.Config) *App {
	return &App{
		data:   make(map[string]Item),
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
	a.data[request.GetKey()] = newItem(request.GetKey(), request.GetValue())
	a.rw.Unlock()

	a.binlog.Add(request)
}

func (a *App) Get(request *dto.Request) (string, bool) {
	a.rw.RLock()
	value, ok := a.data[request.GetKey()]
	a.rw.RUnlock()
	return value.getValue(), ok
}

func (a *App) Push(request *dto.Request) bool {
	a.rw.Lock()
	value, ok := a.data[request.GetKey()]
	if ok {
		valueItem, ok := a.data[request.GetValue()]
		if ok {
			value.items[request.GetValue()] = &valueItem
		}
	}
	a.rw.Unlock()

	if ok {
		a.binlog.Add(request) // TODO IT
	}
	return ok
}

func (a *App) Pull(request *dto.Request) ([]string, bool) {
	var items []string

	a.rw.Lock()
	value, ok := a.data[request.GetKey()]
	if ok {
		for _, item := range value.items {
			items = append(items, item.getValue())
		}
	}
	a.rw.Unlock()

	return items, ok
}

func (a *App) ForceSet(key string, value string) {
	a.data[key] = newItem(key, value)
}
