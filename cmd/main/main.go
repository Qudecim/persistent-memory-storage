package main

import (
	"flag"
	"fmt"
	"os"
	"qudecim/db/appConfig"
	"qudecim/db/internal/app"
	"qudecim/db/internal/transport"
	"runtime/pprof"
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

	binlogWriter := app.NewBinlogWriter(config.Binlog.Directory, 100)
	application := app.NewApp(binlogWriter, config)
	application.Init()

	binlogReader := app.NewBinlogReader(application, config.Binlog.Directory)
	binlogReader.Read()

	go binlogWriter.Run()
	transport.Run(application)

	if *flagProfile {
		profile()
	}
}

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
