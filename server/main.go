package main

import (
	"fmt"

	"github.com/xchenny/keyboard-warrior/src/client"
	"github.com/xchenny/keyboard-warrior/src/web"
)

func main() {
	client, err := client.New("./src/dict/conf.yaml")
	if err != nil {
		fmt.Printf(fmt.Sprintf("Error occurred while creating client. Err: %v\n", err))
	}
	web.InitRouter(client)

	client.Start()
}
