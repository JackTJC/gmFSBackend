package server

import "github.com/gin-gonic/gin"

func HTTPMain() {
	r := gin.Default()
	setRoute(r)
	r.Run(":8080")
}
func setRoute(r *gin.Engine) {
	r.GET("/ping", ping)
}

func ping(c *gin.Context) {
	c.JSON(200, `{"message":"pong"}`)
}
