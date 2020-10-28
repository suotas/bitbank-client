package api

import (
	"fmt"
)

func (t *Ticker) GetTicker(pair string) {
	path := fmt.Sprintf("/%s/ticker", pair)
	res, _ := t.Client.SendRequest(path)
	fmt.Println(res)
}