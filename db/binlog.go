package db

import (
	"log"
	"os"
	"qudecim/db/dto"
	"strconv"
)

type Binlog struct {
	stack chan *dto.Request

	directory           string
	current             string
	everyCheckOversize  bool
	chanceCheckOversize int
	currentSource       *os.File
}

func NewBinlog(directory string, everyCheckOversize bool, chanceCheckOversize int) *Binlog {
	return &Binlog{directory: directory, current: timestamp(), everyCheckOversize: everyCheckOversize, chanceCheckOversize: chanceCheckOversize, stack: make(chan *dto.Request)}
}

func (b *Binlog) Run() {
	b.openBinlog()

	for {
		select {
		case item := <-b.stack:
			b.addToBinlog(item)
		}
	}
}

func (b *Binlog) add(request *dto.Request) {
	b.stack <- request
}

func (b *Binlog) openBinlog() error {

	if b.currentSource != nil {
		b.currentSource.Close()
	}

	f, err := os.OpenFile(b.getCurrentBinlogPath(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	b.currentSource = f

	return nil
}

func (b *Binlog) addToBinlog(request *dto.Request) error {

	if b.everyCheckOversize || randomInThouthand(b.chanceCheckOversize) {
		if b.isOverSizeBinlog() {
			b.changeBinlog(timestamp())
		}
	}

	key := escapeString(request.GetKey())
	value := escapeString(request.GetValue())

	text := key + "\n" + value + "\n"

	if _, err := b.currentSource.WriteString(text); err != nil {
		return err
	}

	return nil
}

func (b *Binlog) changeBinlog(timestampBinlog string) {
	b.current = timestampBinlog
	b.openBinlog()
}

func (b *Binlog) getCurrentBinlogPath() string {
	return b.getBinlogPath(b.current)
}

func (b *Binlog) getBinlogPath(binlog string) string {
	return b.directory + binlog
}

func (b *Binlog) isOverSizeBinlog() bool {
	fileInfo, err := os.Stat(b.getCurrentBinlogPath())
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()
	return fileSize > int64(Config.Binlog.Oversize)
}

func readBinlog(binlog string) error {
	f, err := os.Open(Config.Binlog.Directory + binlog)
	if err != nil {
		return err
	}

	return readFromFile(f)
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
