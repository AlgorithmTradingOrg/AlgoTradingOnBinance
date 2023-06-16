package main

import (
	"fmt"
	"sync"

	"github.com/adshao/go-binance/v2/futures"
)

type DataHandler struct {
}

func NewDataHandler() *DataHandler {
	return &DataHandler{}
}

func (DH *DataHandler) HandlerDepthEvent(event *futures.WsDepthEvent) {

	bidsSum := 0.0
	asksSum := 0.0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		count := 0
		for _, v := range event.Bids {
			if count >= 10 {
				break
			}
			price, amount, err := v.Parse()
			if err != nil {
				fmt.Println(err)
				continue
			}
			bidsSum = bidsSum + price*amount

			count++
		}
		wg.Done()
	}(wg)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		count := 0
		for _, v := range event.Asks {
			if count >= 30 {
				break
			}
			price, amount, err := v.Parse()
			if err != nil {
				fmt.Println(err)
				continue
			}
			asksSum = asksSum + price*amount

			count++
		}
		wg.Done()
	}(wg)
	wg.Wait()
	bidsRatio := bidsSum / (bidsSum + asksSum)
	asksRatio := asksSum / (bidsSum + asksSum)

	fmt.Printf("price0: %+v,priceN: %+v, bidsSum:%+v,asksSum:%+v\n", event.Bids[0].Price, event.Bids[len(event.Bids)-1].Price, bidsRatio, asksRatio)
}

func (DH *DataHandler) HandlerAggTradeEvent(event *futures.WsAggTradeEvent) {
	fmt.Printf("time:%+v, tradetime:%+v,diff:%+v, \tcurrent price:%+v, \tamount:%+v, \tisMaker:%+v\n", event.Time, event.TradeTime, event.Time-event.TradeTime, event.Price, event.Quantity, event.Maker)
}

//至此证明该项目比较好的封装了binance的接口，稳定性需要进一步验证
//下一步，需要完善数据处理的逻辑，以及数据的存储
//再下一步，需要完善交易逻辑，以及交易的存储
//add codes for statistic data
// func (DH *DataHandler)
