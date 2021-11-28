package test

import (
	"context"
	"fmt"
	"go-app/library/boot"
	"go-app/library/log"
	_rocketmq "go-app/library/rocketmq"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func TestRocketMQTopic(t *testing.T) {
	topic := "newOne"
	// clusterName := "DefaultCluster"
	nameSrvAddr := []string{"127.0.0.1:9876"}
	brokerAddr := "127.0.0.1:30911;127.0.0.1:30921;127.0.0.1:30931"

	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	if err != nil {
		fmt.Println(err.Error())
	}

	// create topic
	err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(brokerAddr),
	)
	if err != nil {
		fmt.Println("Create topic error:", err.Error())
	}

	// delete topic
	err = testAdmin.DeleteTopic(
		context.Background(),
		admin.WithTopicDelete(topic),
		// admin.WithBrokerAddrDelete(brokerAddr),
		// admin.WithNameSrvAddr(nameSrvAddr),
	)
	if err != nil {
		fmt.Println("Delete topic error:", err.Error())
	}

	err = testAdmin.Close()
	if err != nil {
		fmt.Printf("Shutdown admin error: %s", err.Error())
	}
}

func TestRocketMQProducer(t *testing.T) {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
	)

	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	topic := "test"
	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello RocketMQ Go Client! " + strconv.Itoa(i)),
		}
		res, err := p.SendSync(context.Background(), msg)
		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}

	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}

func TestRocketMQConsumer(t *testing.T) {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)
	err := c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("subscribe callback: %v \n", msgs[i])
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}

func initRocketMQ() {
	boot.Register(&log.LogStarter{})
	boot.Register(&_rocketmq.RocketMQStarter{})
	testBootRun()
}

func TestRocketMQSendSyncMsg(t *testing.T) {
	initRocketMQ()

	err := _rocketmq.SendSyncMsg("test", []byte("ok"))
	assert.Nil(t, err)
}

func TestRocketMQSendAsyncMsg(t *testing.T) {
	initRocketMQ()

	err := _rocketmq.SendAsyncMsg("test", []byte("ok"+time.Now().String()))
	assert.Nil(t, err)
	time.Sleep(time.Second * 5)
}

func init() {
	const (
		topic   = "test"
		groupId = "test_group"
	)

	_rocketmq.RegisterHandler(groupId, topic, func(msg *primitive.MessageExt) {
		fmt.Println("test consumer rocketmq raw msg:", msg)
		fmt.Printf("test consumer rocketmq msg topic: %s, body: %s\n", msg.Topic, msg.Body)
	})
}

func TestRocketMQConsumerMsg(t *testing.T) {
	initRocketMQ()

	time.Sleep(time.Hour)
}
