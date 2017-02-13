package main

import (
	"bufio"
	"os"
	"fmt"
	"golang.org/x/text/width"
)

func main() {
	var bio = bufio.NewReader(os.Stdin)

	for {
		var line, _, err= bio.ReadLine()
		if err != nil {
			os.Exit(0)
		}
		var wideline = width.Widen.String(string(line))
		fmt.Println(wideline)
	}
}
