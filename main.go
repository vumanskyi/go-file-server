package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
)

const MinArgs = 3

const ExplainMsg  = `usage: file server [host]:port /path/to/dir
For example:
server :9999 .`

func main() {
	if isValidInputArgs() {
		fmt.Fprint(os.Stderr, ExplainMsg)
		os.Exit(2)
	}

	fmt.Println("Server is running....")

	http.Handle("/", http.FileServer(http.Dir(os.Args[2])))
	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}

func isValidInputArgs() bool  {
	return len(os.Args) != MinArgs
}