package main

import (
	client "github.com/elleven11/pantegana/client"
)

func main() {
	config := client.LoadClientConfig()

	client.RunClient(config.Host, config.Port)
}
