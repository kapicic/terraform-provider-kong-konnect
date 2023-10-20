package main

import (
	"fmt"

	"github.com/kong-sdk/pkg/client"
)

func main() {
	client := client.NewClient("https://api-dev.kong.com")

	fmt.Printf("%#v", client)
}
