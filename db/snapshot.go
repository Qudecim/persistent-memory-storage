package db

import (
	"fmt"
	"os"
	"strconv"
)

func getSnapshotPath(snapshot string) string {
	return Config.Snapshot.Directory + snapshot
}

func saveSnapshot() error {
	snapshot := timestamp()
	f, err := os.OpenFile(getSnapshotPath(snapshot), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer f.Close()

	for key, value := range Data {
		key = escapeString(key)
		value = escapeString(value)

		text := key + "\n" + value + "\n"

		if _, err := f.WriteString(text); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func readSnapshot(snapshot string) error {
	f, err := os.Open(getSnapshotPath(snapshot))
	if err != nil {
		return err
	}

	defer f.Close()

	return readFromFile(f)
}

func getLastSnapshot() int {
	files, err := os.ReadDir(Config.Snapshot.Directory)
	if err != nil {
		panic(err)
	}

	lastBinlog := 0
	for _, file := range files {
		if !file.IsDir() {
			if number, err := strconv.Atoi(file.Name()); err == nil {
				lastBinlog = max(lastBinlog, number)
			}
		}
	}

	return lastBinlog
}
