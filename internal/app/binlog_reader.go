package app

import (
	"bufio"
	"os"
	"strconv"
)

type BinlogReader struct {
	directory string
	app       *App
}

func NewBinlogReader(app *App, directory string) *BinlogReader {
	return &BinlogReader{
		directory: directory,
		app:       app,
	}
}

func (b *BinlogReader) Read() {

	binlogs := b.getBinlogs(0)
	for _, binlog := range binlogs {
		b.readBinlog(strconv.Itoa(binlog))
	}

}

func (b *BinlogReader) readBinlog(binlog string) error {
	f, err := os.Open(b.directory + binlog)
	if err != nil {
		return err
	}

	return b.readFromFile(f)
}

func (b *BinlogReader) getBinlogs(fromDate int) []int {
	files, err := os.ReadDir(b.directory)
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

func (b *BinlogReader) readFromFile(file *os.File) error {
	scanner := bufio.NewScanner(file)

	key := ""
	for scanner.Scan() {
		if key == "" {
			key = deescapeString(scanner.Text())
		} else {
			b.app.ForceSet(key, deescapeString(scanner.Text()))
			key = ""
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
