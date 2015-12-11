package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	delimiter = flag.String("d", "", "delimiter substring")
	part      = flag.Int("p", 0, "part to output (starting from 1)")
)

func process(data []byte) {
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if strings.Contains(line, *delimiter) {
			result := strings.Split(line, *delimiter)
			if *part <= len(result) {
				fmt.Println(result[*part-1])
			} else {
				// Printing an empty line if requested part is beyond result's slice.
				fmt.Println()
			}
		} else {
			// Printing the entire line if it doesn't contain the provided delimiter.
			fmt.Println(line)
		}
		if i == len(lines)-2 {
			break
		}
	}
}

func main() {
	flag.Parse()

	if *delimiter != "" && *part > 0 {
		filePath := flag.Arg(0)

		if filePath != "" {
			// Processing data from file.
			stat, fileErr := os.Stat(filePath)
			switch {
			case os.IsNotExist(fileErr):
				log.Fatal(fileErr)
			case stat.IsDir():
				log.Fatal(filePath + " is a directory")
			default:
				data, dataErr := ioutil.ReadFile(filePath)
				if dataErr != nil {
					log.Fatal(dataErr)
				}
				process(data)
			}
		} else {
			// Processing data from STDIN.
			data, stdInErr := ioutil.ReadAll(os.Stdin)
			if stdInErr != nil {
				log.Fatal(errors.New("unable to process STDIN, " + stdInErr.Error()))
			}
			process(data)
		}
	} else {
		log.Fatal(errors.New("both delimiter and part parameters shouldn't be empty"))
	}
}
