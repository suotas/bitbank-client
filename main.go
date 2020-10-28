package main

import (
	"github.com/suotas/bitbank-client/client"
)

func main() {
	client.NewClient("apiToken", "apiSecret")
}