package main

import (
	"encoding/json"
	"fmt"
	"os"

	janus "github.com/feixiao/janus-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: janus-echotest </path/to/socket>\n")
		return
	}

	gateway, err := janus.Connect(os.Args[1])
	if err != nil {
		fmt.Printf("Connect: %s\n")
		return
	}

	info, err := gateway.Info()
	if err != nil {
		fmt.Printf("Info: %s\n", err)
		return
	}

	infoStr, _ := json.MarshalIndent(info, "", "\t")
	fmt.Printf("%s\n", string(infoStr))
	gateway.Close()
}
