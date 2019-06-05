package main

import (
	"github.com/tomocy/server/config"
	"github.com/tomocy/server/server"
)

func main() {
	config.Must(config.LoadConfig("./config.yml"))
	server := server.New("./public")
	server.ListenAndServe(config.Current.Addr)
}
