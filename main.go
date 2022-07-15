package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	for _, item := range sd {
		if string(item) != ":" {
			whatIsIt = string(item) + whatIsIt
		} else {
			whatIsIt = " " + whatIsIt
		}
	}

	fmt.Println(whatIsIt)
}
