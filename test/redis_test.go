package test

import (
	"fmt"
	"go-app/library/boot"
	"go-app/library/redis"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initRedis() {
	boot.Register(&redis.RedisStarter{})
	testBootRun()
}

func TestRedis(t *testing.T) {
	initRedis()

	rc := redis.GetClient()
	err := rc.Ping().Err()
	assert.Nil(t, err)
}

func TestSetGetDel(t *testing.T) {
	initRedis()

	const (
		key = "hello"
		val = "world"
	)

	rc := redis.GetClient()
	setResult, err := rc.Set(key, val, 0).Result()
	assert.Nil(t, err)
	fmt.Println("set result:", setResult)

	getResult, err := rc.Get(key).Result()
	assert.Nil(t, err)
	assert.Equal(t, val, getResult)

	delResult, err := rc.Del(key).Result()
	assert.Nil(t, err)
	fmt.Println("del result:", delResult)
}
