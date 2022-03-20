package server

import "github.com/gin-gonic/gin"

func setRoute(r *gin.Engine) {
	r.GET("/ping", ping)
}

func ping(c *gin.Context) {
	c.JSON(200, `{"message":"pong"}`)
}
