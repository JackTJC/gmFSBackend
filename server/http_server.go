package server

import "github.com/gin-gonic/gin"

func HTTPMain() {
	r := gin.Default()
	setRoute(r)
	r.Run(":8080")
}
