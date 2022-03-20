package main

import (
	"github.com/JackTJC/gmFS_backend/dal"
	"github.com/JackTJC/gmFS_backend/server"
)

func main() {
	dal.InitDB()
	server.Main()
}
