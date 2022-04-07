package main

import (
	"github.com/JackTJC/gmFS_backend/dal"
	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/server"
)

func main() {
	logs.InitLog()
	dal.Init()
	server.GRPCMain()
	//server.Main()
}
