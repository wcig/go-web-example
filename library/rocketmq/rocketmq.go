package rocketmq

import (
	"context"
	"fmt"
	"go-app/library/log"
	"strings"

	"github.com/apache/rocketmq-client-go/v2/consumer"

	"github.com/apache/rocketmq-client-go/v2/producer"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type RocketMQConfig struct {
	Brokers  string            `yaml:"brokers" json:"brokers"`
	Consumer []*ConsumerConfig `yaml:"consumer" json:"consumer"`
}

type ConsumerConfig struct {
	Topic   string `yaml:"topic" json:"topic"`
	GroupId string `yaml:"group_id" json:"group_id"`
}

type HandleFunc = func(msg *primitive.MessageExt)

var (
	rp            rocketmq.Producer
	handleFuncMap = make(map[string]HandleFunc)
)

func Init(cfg *RocketMQConfig) {
	if cfg == nil {
		return
	}

	brokers := strings.Split(cfg.Brokers, ",")
	if len(brokers) == 0 {
		return
	}

	initProducer(brokers)
	initConsumer(brokers, cfg.Consumer)
}

func initProducer(brokers []string) {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(brokers)),
		producer.WithRetry(2),
	)
	if err != nil {
		fmt.Printf("start rocketmq producer error: %s", err.Error())
		panic(err)
	}

	rp = p
	if err = p.Start(); err != nil {
		fmt.Printf("start rocketmq producer error: %s", err.Error())
		panic(err)
	}
}

func initConsumer(brokers []string, configs []*ConsumerConfig) {
	for _, config := range configs {
		c, err := rocketmq.NewPushConsumer(
			consumer.WithGroupName(config.GroupId),
			consumer.WithNsResolver(primitive.NewPassthroughResolver(brokers)),
		)
		if err != nil {
			log.Errorf("start rocketmq consumer error: %s", err.Error())
		}

		key := getConsumerKey(config.GroupId, config.Topic)
		handle := handleFuncMap[key]
		if handle == nil {
			err = fmt.Errorf("rocketmq consumer empty, group: %s, topic: %s", config.GroupId, config.Topic)
			panic(err)
		}

		f := func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgs {
				handle(msgs[i])
			}
			return consumer.ConsumeSuccess, nil
		}
		err = c.Subscribe(config.Topic, consumer.MessageSelector{}, f)
		if err != nil {
			log.Errorf("start rocketmq consumer error: %s", err.Error())
		}

		if err = c.Start(); err != nil {
			log.Errorf("start rocketmq consumer error: %s", err.Error())
		}
	}
}

func SendSyncMsg(topic string, body []byte) error {
	msg := &primitive.Message{
		Topic: topic,
		Body:  body,
	}
	res, err := rp.SendSync(context.Background(), msg)
	if err != nil {
		return err
	}
	log.Debugf("send rocketmq message success: result=%s\n", res.String())
	return nil
}

func SendAsyncMsg(topic string, body []byte) error {
	msg := &primitive.Message{
		Topic: topic,
		Body:  body,
	}
	mq := func(ctx context.Context, res *primitive.SendResult, err error) {
		if err != nil {
			log.Error("send rocketmq async msg error:", err)
		} else {
			log.Debugf("send rocketmq message success: result=%s\n", res.String())
		}
	}
	err := rp.SendAsync(context.Background(), mq, msg)
	return err
}

func RegisterHandler(groupId string, topic string, handle HandleFunc) {
	key := getConsumerKey(groupId, topic)
	handleFuncMap[key] = handle
}

func getConsumerKey(groupId, topic string) string {
	return fmt.Sprintf("%s-%s", topic, groupId)
}
