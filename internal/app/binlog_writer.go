package app

import (
	"os"
	"qudecim/db/dto"
	"strings"
)

type BinlogWriter struct {
	stack chan *dto.Request

	writes    int
	maxWrites int
	directory string

	current       string
	currentSource *os.File
}

func NewBinlogWriter(directory string, maxWrites int) *BinlogWriter {
	return &BinlogWriter{
		directory: directory,
		current:   Timestamp(),
		stack:     make(chan *dto.Request),
		maxWrites: maxWrites,
	}
}

func (b *BinlogWriter) Run() {
	b.openBinlog()

	for {
		select {
		case item := <-b.stack:
			b.addToBinlog(item)

			b.writes++
			if b.writes > b.maxWrites {
				b.changeBinlog(Timestamp())
				b.writes = 0
			}

		}
	}
}

func (b *BinlogWriter) Add(request *dto.Request) {
	b.stack <- request
}

func (b *BinlogWriter) openBinlog() error {

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

func (b *BinlogWriter) addToBinlog(request *dto.Request) error {
	key := strings.ReplaceAll(request.GetKey(), "\n", "\\n")
	value := strings.ReplaceAll(request.GetValue(), "\n", "\\n")

	text := key + "\n" + value + "\n"

	if _, err := b.currentSource.WriteString(text); err != nil {
		return err
	}

	return nil
}

func (b *BinlogWriter) changeBinlog(timestampBinlog string) {
	b.current = timestampBinlog
	b.openBinlog()
}

func (b *BinlogWriter) getCurrentBinlogPath() string {
	return b.getBinlogPath(b.current)
}

func (b *BinlogWriter) getBinlogPath(binlog string) string {
	return b.directory + binlog
}
