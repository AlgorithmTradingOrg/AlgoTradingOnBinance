package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

func main() {
	binance.UseTestnet = false
	wsDepthHandler := func(event *futures.WsDepthEvent) {

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

		fmt.Printf("price: %+v, bidsSum:%+v,asksSum:%+v\n", event.Bids[0].Price, bidsRatio, asksRatio)

		// fmt.Printf("%+v", event.Bids[0])
		//计算N个买盘与卖盘的差价
		// panic("implement me")
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, stopC, err := futures.WsDiffDepthServe("BTCUSDT", wsDepthHandler, errHandler)
	// doneC, stopC, err := futures.WsDepthServe("BTCUSDT", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(50 * time.Second)
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneC
}
