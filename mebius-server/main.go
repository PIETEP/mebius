package main

import (
    "github.com/PIETEP/mebius-server/server"
    "github.com/PIETEP/mebius-server/db"
)

func main() {
    db.Init()
    server.Init()
}
