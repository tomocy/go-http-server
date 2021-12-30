package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tomocy/server/server"
)

func main() {
	conf := parseConfig()
	server := server.New(conf.root)
	if err := server.ListenAndServe(conf.addr); err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen and serve: %v", err)
		os.Exit(1)
	}
}

func parseConfig() *config {
	conf := new(config)
	flag.StringVar(&conf.root, "root", "./", "root of static files to be served")
	flag.StringVar(&conf.addr, "addr", ":80", "address of server")
	flag.Parse()

	return conf
}

type config struct {
	root string
	addr string
}
