package web

import (
	"github.com/gin-gonic/gin"
	"github.com/xchenny/keyboard-warrior/src/client"
)

// CORSMiddleware returns a HandlerFunc that handles CORS errors
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// InitRouter connects HandlerFuncs to Router
func InitRouter(c *client.Client) {
	c.Router.Use(CORSMiddleware())
	c.Router.GET("/words", GetWordsHandler(c, true))
}
