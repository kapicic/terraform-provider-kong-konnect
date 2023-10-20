package main

import (
	"fmt"

	"github.com/liblab-sdk/pkg/client"
)

func main() {
	client := client.NewClient("https://api-dev.liblab.com")

	fmt.Printf("%#v", client)
}
