package main

import (
	"bufio"
	"os"
	"fmt"
	"golang.org/x/text/width"
)

func main() {
	var bio = bufio.NewReader(os.Stdin)

	var line, _, err = bio.ReadLine()
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}

	var widener = width.Widen
	var wideline = widener.String(string(line))
	fmt.Println(wideline)
}
