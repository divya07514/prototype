package main

import (
	"encoding/json"
	"net/http"
	"prototype/simple-chat-application/pub_sub"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var clientConns = make(map[string]*websocket.Conn)
var mutex = &sync.Mutex{}
var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}
var subscribedChannels = []string{}
var channels = map[string][]string{
	"ch-1": {"client-1", "client-2"},
}

var clientToChannels = map[string][]string{
	"client-1": {"ch-1"},
	"client-2": {"ch-1"},
}

type Message struct {
	ClientID string `json:"clientId"`
	Payload  string `json:"payload"`
}

func main() {
	r := mux.NewRouter()
	pub_sub.InitRedis()
	r.Use(mux.CORSMethodMiddleware(r))

	// Handlers for port 8080
	r8080 := r.PathPrefix("/server_one").Subrouter()
	r8080.HandleFunc("/ws/{clientId}", connectionHandler)
	r8080.HandleFunc("/message/send", messageHandler)

	// Handlers for port 8081
	r8081 := r.PathPrefix("/server_two").Subrouter()
	r8081.HandleFunc("/ws/{clientId}", connectionHandler)
	r8081.HandleFunc("/message/send", messageHandler)

	log.Info().Msg("starting web servers on ports 8080 and 8081")
	go func() {
		if err := http.ListenAndServe(":8080", r8080); err != nil {
			log.Error().Msgf("error starting web server on port 8080: %v", err)
		}
	}()
	go func() {
		if err := http.ListenAndServe(":8081", r8081); err != nil {
			log.Error().Msgf("error starting web server on port 8081: %v", err)
		}
	}()
	select {}
}

func messageHandler(writer http.ResponseWriter, request *http.Request) {
	var req Message
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		log.Error().Msgf("error decoding request: %v", err)
		http.Error(writer, "error decoding request", http.StatusBadRequest)
		return
	}
	log.Info().Msgf("received message: %s from clientId: %s", req.Payload, req.ClientID)
	marshal, err := json.Marshal(req)
	if err != nil {
		log.Error().Msgf("error encoding request: %v", err)
		http.Error(writer, "error encoding request", http.StatusInternalServerError)
		return
	}

	clientId := req.ClientID
	var channel string
	if clientId == "client-1" || clientId == "client-2" {
		channel = "ch-1"
	} else {
		channel = "ch-2"
	}

	pub_sub.Publish(channel, marshal)
}

func connectionHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	clientId := vars["clientId"]
	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Error().Msgf("error upgrading http connection to a websocket connection: %s", err.Error())
	}
	mutex.Lock()
	clientConns[clientId] = ws
	mutex.Unlock()
	log.Info().Msgf("client %s connected", clientId)
	values := clientToChannels[clientId]
	for _, value := range values {
		if !contains(subscribedChannels, value) {
			ch := pub_sub.Subscribe(value)
			log.Info().Msgf("clientId %s subscribed to channel %s", clientId, value)
			subscribedChannels = append(subscribedChannels, value)
			go handleMessages(ch)
		}
	}
}

func handleMessages(ch <-chan *redis.Message) {
	for message := range ch {
		var data Message
		err := json.Unmarshal([]byte(message.Payload), &data)
		if err != nil {
			log.Error().Msgf("error unmarshalling message: %s", err.Error())
			continue
		}
		channelMembers := channels[message.Channel]
		for _, member := range channelMembers {
			if member != data.ClientID {
				mutex.Lock()
				conn := clientConns[member]
				log.Info().Msg("writing message on ws connection")
				err = conn.WriteMessage(websocket.TextMessage, []byte(data.Payload))
				if err != nil {
					log.Error().Msgf("error writing message: %s", err.Error())
				}
				mutex.Unlock()
			}
		}
	}
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
