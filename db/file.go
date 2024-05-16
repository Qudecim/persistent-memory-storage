package db

import (
	"bufio"
	"os"
)

func readFromFile(file *os.File) error {
	scanner := bufio.NewScanner(file)

	key := ""
	for scanner.Scan() {
		if key == "" {
			key = deescapeString(scanner.Text())
		} else {
			Data[key] = deescapeString(scanner.Text())
			key = ""
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
