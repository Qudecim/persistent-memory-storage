package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func profile() {
	f, err := os.Create("cpuprofile.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println(err)
		return
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 100; i++ {
		// db.Set("test_key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
	}
}
