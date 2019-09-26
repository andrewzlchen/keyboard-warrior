package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xchenny/keyboard-warrior/src/client"
)

// GetWordsHandler returns a Handler function that uses a client to get words for the typing test
func GetWordsHandler(c *client.Client, lowercase bool) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// generate new words
		// TODO: Figure out how to store a central word list amongst
		// competitors so they all have the same words
		toks, err := client.GetRandomText(c, true)
		if err != nil {
			errMsg := fmt.Sprintf("Error occurred while getting random text. Err: %v\n", err)
			ctx.JSON(500, gin.H{
				"error": errMsg,
			})
			return
		}

		if toks == nil {
			ctx.JSON(500, gin.H{
				"error": "Server error. No tokens generated.",
			})
			return
		}

		// return the words
		ctx.JSON(200, gin.H{
			"words": toks,
		})
	}

}
