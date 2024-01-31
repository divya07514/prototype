package main

import (
	"github.com/rs/zerolog/log"
	"sync"
)

func main() {
	foodId := 1
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			order, err := PlaceOrder(foodId)
			wg.Done()
			if err != nil {
				log.Error().Msg("order not placed")
			} else {
				log.Error().Msgf("order placed: ", order.OrderId)
			}
		}()
	}
	wg.Wait()
}
