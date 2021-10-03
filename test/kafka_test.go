package test

import (
	"go-app/library/boot"
	"go-app/library/kafka"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initKafka() {
	boot.Register(&kafka.KafkaStarter{})
	testBootRun()
}

func TestProducer(t *testing.T) {
	initKafka()

	err := kafka.SendMsg([]byte("ok"), "test")
	assert.Nil(t, err)
}
