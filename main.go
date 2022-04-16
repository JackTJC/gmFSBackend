package main

import (
	"github.com/JackTJC/gmFS_backend/dal"
	"github.com/JackTJC/gmFS_backend/http_server"
	"github.com/JackTJC/gmFS_backend/logs"
)

func main() {
	logs.InitLog()
	dal.Init()
	// grpc_server.GRPCMain()
	http_server.Main()
}
