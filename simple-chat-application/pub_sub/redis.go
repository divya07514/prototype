package pub_sub

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

var rdb *redis.Client
var pubsub *redis.PubSub

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	log.Error().Msg("initializing redis client")
}

func Subscribe(channel string) <-chan *redis.Message {
	pubsub = rdb.Subscribe(context.Background(), channel)
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
	return pubsub.Channel()
}

func Publish(channel string, message []byte) {
	_, err := rdb.Publish(context.Background(), channel, message).Result()
	if err != nil {
		log.Error().Msgf("error {%s} publishing message on channel %s", err.Error(), channel)
	} else {
		log.Info().Msgf("published message on channel %s", channel)
	}

}
