package main

import (
	"fmt"
	"net/url"
)

func get(s string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover with err", err)
		}
	}()
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("url is", u)
}
func main() {
	get("https://grab.com")
	get("https://abc.com 123")
}
