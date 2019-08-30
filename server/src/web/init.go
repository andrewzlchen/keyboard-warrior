package web

import (
	"github.com/xchenny/keyboard-warrior/src/client"
)

// InitRouter connects HandlerFuncs to Router
func InitRouter(c *client.Client) {
	c.Router.GET("/words", GetWordsHandler(c, true))
}
