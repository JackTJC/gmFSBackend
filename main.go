package main

import (
	"github.com/JackTJC/gmFS_backend/dal"
	"github.com/JackTJC/gmFS_backend/http_server"
	"github.com/JackTJC/gmFS_backend/logs"
	"github.com/JackTJC/gmFS_backend/util"
)

func main() {
	logs.InitLog()
	dal.Init()
	util.Init()
	http_server.Main(true)
}
