package redis

import (
	"context"
	"sync"

	"bit-project/gateway/config"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var (
	once   sync.Once
	client *redis.Client
)

// Open create a new *redis client
func Open(conf *config.Config) *redis.Client {
	once.Do(func() { // atomic, does not allow repeating
		client = connect(conf)
	})
	return client
}

func connect(conf *config.Config) *redis.Client {
	var (
		err     error
		rClient *redis.Client
	)
	rClient = redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost + ":" + conf.RedisPort,
		Password: conf.RedisPassword,
		DB:       0,
	})
	pong, err := rClient.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("Error connecting to Redis: %s", err)
	}
	log.Infof("Connected to Redis: %s", pong)
	return rClient
}
