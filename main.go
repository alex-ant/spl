package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	delimiter       = flag.String("d", "", "delimiter substring")
	part            = flag.String("p", "", "part to output (starting from 1)")
	outputDelimiter = flag.String("od", ",", "output delimiter")
	sequence        = flag.Bool("s", false, "splits string by as long sequence of provided delimiter as found")
)

var parts []int

func process(data []byte) {
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if strings.Contains(line, *delimiter) {
			result := strings.Split(line, *delimiter)
			if *sequence {
				var resultTmp []string
				for _, v := range result {
					if v != "" {
						resultTmp = append(resultTmp, v)
					}
				}
				result = resultTmp
			}

			var out string
			for k, p := range parts {
				if k > 0 {
					out += *outputDelimiter
				}
				if p <= len(result) {
					out += result[p-1]
				}
			}
			fmt.Println(out)
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

	if *delimiter != "" && *part != "" {
		partsStr := strings.Split(*part, ",")
		for _, p := range partsStr {
			partInt, partErr := strconv.Atoi(p)
			if partErr != nil {
				log.Fatal(partErr)
			}
			parts = append(parts, partInt)
		}

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
