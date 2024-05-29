package app

import (
	"fmt"
	"os"
	"qudecim/db/appConfig"
	"qudecim/db/internal/dto"
	"sync"
)

type App struct {
	data   map[string]*Item
	binlog *BinlogWriter
	config *appConfig.Config

	rw sync.RWMutex
	Wg sync.WaitGroup
}

func NewApp(binlog *BinlogWriter, config *appConfig.Config) *App {
	return &App{
		data:   make(map[string]*Item),
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
	item, exist := a.data[request.GetKey()]
	if exist {
		item.value = request.GetValue()
	} else {
		a.data[request.GetKey()] = newItem(request.GetKey(), request.GetValue())
	}
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
	parent, exist := a.data[request.GetKey()]
	if !exist {
		parent = newItemList(request.GetKey())
		a.data[request.GetKey()] = parent
	}
	valueItem, ok := a.data[request.GetValue()]
	if ok {
		parent.items[request.GetValue()] = valueItem
	}
	fmt.Println(a.data)
	a.rw.Unlock()

	if ok {
		a.binlog.Add(request)
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

func (a *App) ForcePush(key string, value string) {
	parent, exist := a.data[key]
	if !exist {
		parent = newItemList(key)
		a.data[key] = parent
	}
	valueItem, ok := a.data[value]
	if ok {
		parent.items[value] = valueItem
	}
}
