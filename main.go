package main

import (
	"fmt"

	"github.com/JackTJC/gmFS_backend/config"
	"github.com/JackTJC/gmFS_backend/dal"
)

func main() {
	dal.InitDB()
	fmt.Println("Hello World")
	config.GetInstance()
}
