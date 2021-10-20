package main

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	/*if len(args) > 1 {
		log.Fatal("Currently only one file supported.")
	}*/

	if len(args) == 0 {
		log.Fatal("Please enter file name.")
	}

	for _, element := range args {
		filerc, err := os.Open(element)
		if err != nil {
			log.Fatal(err)
		}
		defer filerc.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(filerc)
		contents := buf.String()

		for _, element := range contents {
			curr := string("\033[3"+strconv.Itoa(rand.Intn(6-1)+1)+"m") + string(element)

			print(curr)
		}
	}
}
