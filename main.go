package main

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

func main() {
	binance.UseTestnet = false

	dataHandler := NewDataHandler()
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsDiffDepthServe("IDUSDT", dataHandler.HandlerDepthEvent, errHandler) //wsDepthHandler, errHandler)
	// doneC, stopC, err := futures.WsDepthServe("BTCUSDT", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	futures.WsAggTradeServe("IDUSDT", dataHandler.HandlerAggTradeEvent, errHandler)
	// use stopC to exit
	// go func() {
	// 	time.Sleep(50 * time.Second)
	// 	stopC <- struct{}{}
	// }()
	// remove this if you do not want to be blocked here
	<-doneC
}
