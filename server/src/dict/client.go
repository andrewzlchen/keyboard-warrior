package dict

import "net/http"

// Client is a structure that can perform operations
// that is related to getting words for the typing test
type Client struct {
	config     *Config
	httpClient *http.Client
}

// NewClient makes a new Client that can obtain text for the players to type
func NewClient(confPath string) (*Client, error) {
	conf, err := ParseConfig(confPath)
	if err != nil {
		return nil, err
	}

	httpClient := &(http.Client{})

	client := &(Client{
		config:     conf,
		httpClient: httpClient,
	})

	return client, nil
}
