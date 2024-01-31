package main

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net/http"
	"time"
)

func PlaceOrder(foodId int) (*OrderDetails, error) {
	//-------------------------------------------------
	body, err := json.Marshal(map[string]interface{}{
		"FoodId": foodId,
	})
	requestBody := bytes.NewBuffer(body)
	resp, err := http.NewRequest(http.MethodPost, "http://localhost:8081/store/food/reserve", requestBody)
	resp.Header.Add("Accept", `application/json`)
	resp.Header.Add("Content-Type", `application/json`)
	client := createHttpClient()
	_, err = client.Do(resp)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, nil
	}

	//-------------------------------------------------
	resp, err = http.NewRequest(http.MethodPost, "http://localhost:8082/delivery/agent/reserve", requestBody)
	resp.Header.Add("Accept", `application/json`)
	resp.Header.Add("Content-Type", `application/json`)
	_, err = client.Do(resp)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, nil
	}

	//-------------------------------------------------
	rand.Seed(time.Now().UnixNano())
	orderId := rand.Intn(1000000) + 1
	body, _ = json.Marshal(map[string]interface{}{
		"OrderId": orderId,
		"FoodId":  foodId,
	})
	requestBody = bytes.NewBuffer(body)
	resp, err = http.NewRequest(http.MethodPost, "http://localhost:8081/store/food/book", requestBody)
	resp.Header.Add("Accept", `application/json`)
	resp.Header.Add("Content-Type", `application/json`)
	_, err = client.Do(resp)
	if err != nil {
		log.Error().Msg("could not book food packet")
		return nil, nil
	}

	//-------------------------------------------------
	body, _ = json.Marshal(map[string]interface{}{
		"OrderId": orderId,
	})
	requestBody = bytes.NewBuffer(body)
	resp, err = http.NewRequest(http.MethodPost, "http://localhost:8082/delivery/agent/book", requestBody)
	resp.Header.Add("Accept", `application/json`)
	resp.Header.Add("Content-Type", `application/json`)
	_, err = client.Do(resp)
	if err != nil {
		log.Error().Msg("could not assign delivery agent")
		return nil, nil
	}

	return &OrderDetails{OrderId: orderId}, nil
}

func createHttpClient() *http.Client {
	tr := &http.Transport{
		MaxConnsPerHost:     100,
		MaxIdleConnsPerHost: 100,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
	return client
}
