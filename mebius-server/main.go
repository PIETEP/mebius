package main

import (
	"github.com/PIETEP/mebius/mebius-server/db"
	"github.com/PIETEP/mebius/mebius-server/server"
)

func main() {
	db.Init()
	server.Init()
}
