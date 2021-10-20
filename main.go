package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	iter := 1

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
			curr := string("\033[3"+strconv.Itoa(iter)+"m") + string(element)
			if iter == 6 {
				iter = 0
			}
			iter++

			print(curr)
		}
		if iter > 1 {
			iter--
		}

	}
}
