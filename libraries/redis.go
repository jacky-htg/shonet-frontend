package libraries

import (
	"fmt"
	"github.com/jacky-htg/shonet-frontend/config"
	r "gopkg.in/redis.v5"
	"time"
)

var redisName = config.GetString("laravel.redisName")
var SHONET_REDIS = fmt.Sprintf("%s:", redisName)

var client = r.NewClient(&r.Options{
	Addr:		config.GetString("database.redis.address"),
	Password:	config.GetString("database.redis.password"),
})

func RedisSet(key string, val interface{}, duration time.Duration) error {
	key = SHONET_REDIS + key
	if err := client.SetNX(key, val, duration).Err(); err != nil {
		return err
	}

	return nil
}

func RedisGet(key string) ([]byte, error) {
	key = SHONET_REDIS + key
	getting, err := client.Get(key).Bytes()
	if err != nil {
		return []byte{}, err
	}

	return getting, nil
}

func RedisExists(key string) bool {
	key = SHONET_REDIS + key
	if exixts := client.Exists(key).Val(); exixts {
		return true
	}

	return false
}


func RedisDelete(key string) (bool, error) {
	key = SHONET_REDIS + key
	if err := client.Del(key).Err(); err != nil {
		return false, err
	}

	return true, nil
}

func RedisHashSet(key string, field string, val interface{}) error {
	key = SHONET_REDIS + key
	if err := client.HSet(key, field, val).Err(); err != nil {
		return err
	}

	return nil
}

func RedisHashExists(key string, field string) bool {
	key = SHONET_REDIS + key
	if exists := client.HExists(key, field).Val(); exists {
		return true
	}

	return false
}
