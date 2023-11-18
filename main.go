package main

import (
	"flag"

	"mythos-auth/config"
	"mythos-auth/server"
)

func main() {
	env := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*env)

	if err := server.Init(); err != nil {
		panic(err)
	}
}
