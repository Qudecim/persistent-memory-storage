package main

import (
	"fmt"
	"os"
	"qudecim/db/db"
	"runtime/pprof"
	"strconv"
	"time"
)

func profile() {
	// Создаем файл для записи профиля CPU
	f, err := os.Create("cpuprofile.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Запускаем профилирование CPU
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println(err)
		return
	}
	defer pprof.StopCPUProfile()

	// Ваш код, который вы хотите профилировать
	for i := 0; i < 1000000; i++ {
		db.Set("test_key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
	}

	// Для примера добавим небольшую паузу
	time.Sleep(2 * time.Second)
}
