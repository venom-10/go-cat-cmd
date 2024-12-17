package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CatFile(fileName string) error {
	file, err := os.Open(fileName)

	if err != nil {
		return fmt.Errorf("error opening file %s: %v", fileName, err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}

		// Print the line
		fmt.Print(line)

		// Break if we've reached the end of the file
		if err == io.EOF {
			break
		}
	}

	return nil
}
