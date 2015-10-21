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

func main() {
	flag.Parse()

	var result []string

	if *delimiter != "" && *part > 0 {
		stdInBytes, stdInErr := ioutil.ReadAll(os.Stdin)
		if stdInErr != nil {
			log.Fatal(errors.New("unable to process STDIN, " + stdInErr.Error()))
		}
		bytesWithoutNewLine := strings.Split(string(stdInBytes), "\n")[0]
		result = strings.Split(bytesWithoutNewLine, *delimiter)
		if *part <= len(result) {
			fmt.Println(result[*part-1])
		} else {
			log.Fatal(errors.New("index outside result's range"))
		}
	} else {
		log.Fatal(errors.New("both delimiter and part parameters shouldn't be empty"))
	}
}
