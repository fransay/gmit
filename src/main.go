package main

import "flag"

func main() {
	add := flag.String("add", "", "stage files, eg. add main.go")
	commit := flag.String("commit", "", "commit files, eg. commit main.go, defaults to llm generated msg")

	flag.Parse()
}
