package binlog

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type BinlogReader struct {
	directory string

	fromDate           int
	binlogs            []int
	currentBinlogIndex int

	scanner *bufio.Scanner

	key   string
	value string
}

func NewBinlogReader(directory string, fromDate int) *BinlogReader {
	return &BinlogReader{
		directory: directory,
		fromDate:  fromDate,
	}
}

func (b *BinlogReader) openNext() error {

	binlog := b.binlogs[b.currentBinlogIndex]

	file, err := os.Open(b.directory + strconv.Itoa(binlog))
	if err != nil {
		return err
	}

	b.currentBinlogIndex++
	b.scanner = bufio.NewScanner(file)

	return nil
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

func (b *BinlogReader) Init() {
	b.binlogs = b.getBinlogs(b.fromDate)
	b.openNext()
}

func (b *BinlogReader) Next() bool {

	if b.read() {
		return true
	} else {
		if b.currentBinlogIndex >= len(b.binlogs)-1 {
			b.openNext()
			return b.read()
		}
	}

	return false
}

func (b *BinlogReader) read() bool {

	if b.scanner.Scan() {
		b.key = strings.ReplaceAll(b.scanner.Text(), "\\n", "\n")

		if b.scanner.Scan() {
			b.value = strings.ReplaceAll(b.scanner.Text(), "\\n", "\n")
		}

		return true
	}

	return false

}

func (b *BinlogReader) GetKey() string {
	return b.key
}

func (b *BinlogReader) GetValue() string {
	return b.value
}

func (b *BinlogReader) Err() error {
	return b.scanner.Err()
}
