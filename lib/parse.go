package lib

import (
	"bufio"
	"os"
)

// Parses the input file, calling the callback function line-by-line
// Return an error in the callback function to stop the read and propagate the error
func ParseInputByLine(fileName string, cb func(line string) error) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		err := cb(scanner.Text())
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
