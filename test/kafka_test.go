package test

import (
	"fmt"
	"go-app/library/boot"
	"go-app/library/kafka"
	"go-app/library/log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func initKafka() {
	boot.Register(&log.LogStarter{})
	boot.Register(&kafka.KafkaStarter{})
	testBootRun()
}

func TestProducer(t *testing.T) {
	initKafka()

	err := kafka.SendMsg([]byte("ok"), "test")
	assert.Nil(t, err)

	err = kafka.SendKeyValMsg([]byte("key"), []byte("val"), "test")
	assert.Nil(t, err)
	time.Sleep(time.Minute)
}

func TestConsumer(t *testing.T) {
	const (
		topic   = "test"
		groupId = "test_group"
	)

	kafka.RegisterHandler(groupId, topic, func(key, val []byte) {
		fmt.Printf("kafka consumer msg key:%s, msg:%s\n", string(key), string(val))
	})
	initKafka()
	time.Sleep(time.Hour)
}
