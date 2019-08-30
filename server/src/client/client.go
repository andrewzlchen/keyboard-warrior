package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Client is a structure that can perform operations
// that is related to getting words for the typing test
type Client struct {
	HTTPClient *http.Client
	Router     *gin.Engine
}

// New makes a new Client that can obtain text for the players to type
func New(confPath string) (*Client, error) {
	httpClient := &(http.Client{})

	client := &(Client{
		HTTPClient: httpClient,
		Router:     gin.Default(),
	})

	return client, nil
}

// Start client web server
func (c *Client) Start() {
	c.Router.Run()
}
