package fileutil

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err1 := f.Close(); err != nil {
		err = err1
	}

	if err != nil {
		return nil, err
	}

	return lines, nil
}

func WriteLines(file string, lines []string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	_, err = f.WriteString(strings.Join(lines, "\n"))
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}

func AppendFile(file string, data []byte) error {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}
