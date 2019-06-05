package main

import (
	"flag"

	"github.com/tomocy/server/server"
)

func main() {
	conf := parseConfig()
	server := server.New(conf.root)
	server.ListenAndServe(conf.addr)
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
