package main

import (
	"fmt"

	"github.com/xchenny/keyboard-warrior/src/dict"
)

func main() {
	client, err := dict.NewClient("./src/dict/conf.yaml")
	if err != nil {
		fmt.Printf(fmt.Sprintf("Error occurred while creating client. Err: %v\n", err))
	}

	lowercase := false
	toks, err := client.GetRandomText(lowercase)
	if err != nil {
		fmt.Printf(fmt.Sprintf("Error occurred while getting random text. Err: %v\n", err))
	}

	for i := 0; i < len(toks); i++ {
		fmt.Println(toks[i])
	}
}
