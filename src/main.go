package main

import (
	"flag"
	"fmt"
)

func main() {
	var arg Args
	flag.StringVar(&arg.add, "add", "", "stage files, eg. add main.go")
	flag.StringVar(&arg.commit, "commit", "", "commit files, eg. commit main.go, defaults to llm generated msg")
	flag.Parse()

	fmt.Println(arg.add)
	fmt.Println(arg.commit)
}
