package server

import "github.com/gin-gonic/gin"

func Main() {
	r := gin.Default()
	setRoute(r)
	r.Run(":8080")
}
