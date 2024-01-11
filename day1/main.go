package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var path = "input.txt"

func main() {
	start := time.Now()
	readFile()
	fmt.Println(time.Since(start))
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func readFile() {
	var file, err = os.Open(path)
	if isError(err) {
		return
	}

	textByte := make([]byte, 22468)
	text, err := file.Read(textByte)
	if isError(err) {
		return
	}
	word := string(textByte[:text])
	var result = 0

	re := regexp.MustCompile(`[^0-9]+(\d+)[^0-9]+`)
	matches := re.FindAllString(word, -1)
	fmt.Println(matches)

	for _, match := range matches {
		numbs, _ := strconv.Atoi(match)
		result = result + numbs
	}

	fmt.Println(result)
	// fmt.Printf("%d bytes: %s\n", text, string(textByte[:text]))
}
