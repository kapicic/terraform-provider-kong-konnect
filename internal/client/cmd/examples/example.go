package main

import (
	"fmt"

	"github.com/kong-sdk/pkg/client"
)

func main() {
	client := client.NewClient("http://example.com", "TOKEN")

		fmt.Printf("%#v", client)
}
