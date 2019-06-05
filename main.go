package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/tomocy/server/config"
	"github.com/tomocy/server/server"
)

func main() {
	root, err := parseRoot()
	if err != nil {
		showHelp()
		return
	}
	config.Must(config.LoadConfig("./config.yml"))
	server := server.New(root)
	server.ListenAndServe(config.Current.Addr)
}

func parseRoot() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("root path is not specified")
	}

	return os.Args[1], nil
}

func showHelp() {
	fmt.Print("Usage of ./server:\n\t./server ROOT\n\n")
}
