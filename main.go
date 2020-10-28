package main

import (
	"fmt"
	"github.com/suotas/bitbank-client/client"
)

func main() {
	c, _ := client.NewClient("apiToken", "apiSecret")
	fmt.Println(c.SendRequest("http://google.com"))
}