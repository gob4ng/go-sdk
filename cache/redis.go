package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redis"
	"github.com/gob4ng/go-sdk/utils"
	"time"
)

type RedisSetup struct {
	PrefixKey string
	Host      string
	Port      string
	Username  string
	Password  string
	DB        int
}

type redisConfig struct {
	redisSetup RedisSetup
	client     *redis.Client
	context    context.Context
}

func NewRedisConfig(setup RedisSetup) (*redisConfig, *error) {

	client := redis.NewClient(&redis.Options{
		Addr:     setup.Host + ":" + setup.Port,
		Password: setup.Password,
		DB:       setup.DB,
		Username: setup.Username,
	})

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, &err
	}

	redisConfig := redisConfig{
		redisSetup: setup,
		client:     client,
		context:    ctx,
	}

	return &redisConfig, nil

}

func (r redisConfig) SetString(key string, value string, expiration int) *error {

	redisKey := r.redisSetup.PrefixKey + "_" + key
	if statusCmd := r.client.Set(r.context, redisKey, value,
		time.Millisecond*time.Duration(expiration)); statusCmd != nil {

		if statusCmd.Err() != nil {
			newError := errors.New(statusCmd.Err().Error())
			return &newError

		}
	}

	return nil

}

func (r redisConfig) GetString(key string) (*string, *error) {

	redisKey := r.redisSetup.PrefixKey + "_" + key
	value, err := r.client.Get(r.context, redisKey).Result()
	if err != nil {
		return nil, &err
	}

	return &value, nil
}

func (r redisConfig) SetJson(key string, json interface{}, expiration int) *error {

	redisKey := r.redisSetup.PrefixKey + "_" + key
	if statusCmd := r.client.Set(r.context, redisKey, utils.JsonToString(json),
		time.Millisecond*time.Duration(expiration)); statusCmd != nil {

		if statusCmd.Err() != nil {
			newError := errors.New(statusCmd.Err().Error())
			return &newError

		}
	}

	return nil
}

func (r redisConfig) GetJson(key string) (*interface{}, *error) {

	redisKey := r.redisSetup.PrefixKey + "_" + key
	value, err := r.client.Get(r.context, redisKey).Result()
	if err != nil {
		return nil, &err
	}

	structJson, errJson := utils.JsonStringToStruct(value)
	if errJson != nil {
		return nil, errJson
	}

	return structJson, nil
}

func (r redisConfig) SetXml(key string, xml interface{}, expiration int) *error {

	redisKey := r.redisSetup.PrefixKey + "_" + key
	if statusCmd := r.client.Set(r.context, redisKey, utils.XmlToString(xml),
		time.Millisecond*time.Duration(expiration)); statusCmd != nil {

		if statusCmd.Err() != nil {
			newError := errors.New(statusCmd.Err().Error())
			return &newError

		}
	}

	return nil
}

func (r redisConfig) GetXml(key string) (*interface{}, *error) {

	redisKey := r.redisSetup.PrefixKey + "_" + key
	value, err := r.client.Get(r.context, redisKey).Result()
	if err != nil {
		return nil, &err
	}

	structJson, errJson := utils.XmlStringToStruct(value)
	if errJson != nil {
		return nil, errJson
	}

	return structJson, nil
}
