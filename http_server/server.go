package http_server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func Main(isHTTPS bool) {
	r := gin.Default()
	setRoute(r)
	if isHTTPS {
		r.Use(tlsHandler(9000))
		r.RunTLS(":9000", "./config/cert/domain.pem", "./config/cert/domain.key")
	} else {
		r.Run(":9000")
	}
}

func tlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			c.Abort()
			return
		}

		c.Next()
	}
}
