package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		log.Fatal("Currently only one file supported.")
	}

	if len(args) == 0 {
		log.Fatal("Please enter file name.")
	}

	arg := args[0]

	filerc, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()

	newStr := ""
	iter := 1

	for _, element := range contents {
		newStr += string("\033[3"+strconv.Itoa(iter)+"m") + string(element)

		if iter == 6 {
			iter = 0
		}
		iter++
	}

	fmt.Println(newStr)
}
