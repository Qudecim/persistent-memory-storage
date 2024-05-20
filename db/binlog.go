package db

import (
	"log"
	"os"
	"strconv"
)

func openBinlog() error {

	if CurrentBinlogSource != nil {
		CurrentBinlogSource.Close()
	}

	f, err := os.OpenFile(getCurrentBinlogPath(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	CurrentBinlogSource = f
	return nil
}

func addToBinlog(key string, value string) error {

	if Config.Binlog.EveryCheckOversize || randomInThouthand(Config.Binlog.ChanceCheckOversize) {
		if isOverSizeBinlog() {
			changeBinlog(timestamp())
		}
	}

	key = escapeString(key)
	value = escapeString(value)

	text := key + "\n" + value + "\n"

	if _, err := CurrentBinlogSource.WriteString(text); err != nil {
		return err
	}

	return nil
}

func readBinlog(binlog string) error {
	f, err := os.Open(getBinlogPath(binlog))
	if err != nil {
		return err
	}
	CurrentBinlogSource = f

	return readFromFile(CurrentBinlogSource)
}

func getBinlogs(fromDate int) []int {
	files, err := os.ReadDir(Config.Binlog.Directory)
	if err != nil {
		panic(err)
	}

	var binlogs []int
	for _, file := range files {
		if !file.IsDir() {
			if number, err := strconv.Atoi(file.Name()); err == nil {
				if number > fromDate {
					binlogs = append(binlogs, number)
				}
			}
		}
	}

	return binlogs
}

func changeBinlog(timestampBinlog string) {
	CurrentBinlog = timestampBinlog
	openBinlog()
}

func getCurrentBinlogPath() string {
	return getBinlogPath(CurrentBinlog)
}

func getBinlogPath(binlog string) string {
	return Config.Binlog.Directory + binlog
}

func isOverSizeBinlog() bool {
	fileInfo, err := os.Stat(getCurrentBinlogPath())
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()
	return fileSize > int64(Config.Binlog.Oversize)
}
