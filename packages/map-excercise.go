package main

import (
 	"golang.org/x/tour/wc"
 	"strings"
)

func WordCount(s string) map[string]int {
	var maps = make(map[string]int)
	for _, value := range strings.Fields(s) {
		count := maps[value]
		maps[value] = count + 1
	}
	return maps
}

func main() {
	wc.Test(WordCount)
}