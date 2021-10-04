package kafka

import (
	"errors"
	"fmt"
	"go-app/library/log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

type KafkaConfig struct {
	Brokers  string            `yaml:"brokers" json:"brokers"`
	Consumer []*ConsumerConfig `yaml:"consumer" json:"consumer"`
}

type ConsumerConfig struct {
	Topic   string `yaml:"topic" json:"topic"`
	GroupId string `yaml:"group_id" json:"group_id"`
}

type HandleFunc = func(msgKey, msgVal []byte)

var (
	producer      sarama.AsyncProducer
	consumerMap   = make(map[string]*cluster.Consumer)
	handleFuncMap = make(map[string]HandleFunc)
)

func Init(cfg *KafkaConfig) {
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
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	config.Net.DialTimeout = 30 * time.Second
	config.Net.ReadTimeout = 60 * time.Second
	config.Net.WriteTimeout = 60 * time.Second
	config.Net.MaxOpenRequests = 8
	config.Net.KeepAlive = 6 * time.Minute
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 2
	config.ChannelBufferSize = 1024 * 1024

	asyncProducer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	go func(p sarama.AsyncProducer) {
		for {
			select {
			case errs := <-p.Errors():
				if errs != nil {
					log.Error("kafka producer err:", errs)
				}
			case successes := <-p.Successes():
				log.Debugf("kafka producer send success:%+v", successes)
			}
		}
	}(asyncProducer)

	producer = asyncProducer
}

func initConsumer(brokers []string, configs []*ConsumerConfig) {
	if len(configs) == 0 {
		return
	}

	initClient := func(brokers []string, topics []string, groupId string) {
		config := cluster.NewConfig()
		config.Version = sarama.V2_1_0_0
		config.Net.DialTimeout = 30 * time.Second
		config.Net.ReadTimeout = 60 * time.Second
		config.Net.WriteTimeout = 60 * time.Second
		config.Net.MaxOpenRequests = 32
		config.Net.KeepAlive = 6 * time.Minute
		config.ClientID = groupId
		config.Consumer.Return.Errors = true
		config.Group.Return.Notifications = true
		config.Group.Mode = cluster.ConsumerModeMultiplex
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
		config.Consumer.Offsets.CommitInterval = time.Second

		consumer, err := cluster.NewConsumer(brokers, groupId, topics, config)
		if err != nil {
			panic(err)
		}

		key := getConsumerKey(groupId, strings.Join(topics, ","))
		consumerMap[key] = consumer

		go func() {
			for err := range consumer.Errors() {
				log.Error("kafka consumer err:", err)
			}
		}()

		go func() {
			for ntf := range consumer.Notifications() {
				log.Debugf("kafka consumer notification:%+v", ntf)
			}
		}()

		go func() {
			for {
				select {
				case msg, ok := <-consumer.Messages():
					if ok && msg != nil {
						log.Debugf("kafka consumer msg: topic:%s, partition:%d, offset:%d, key:%s, val:%s",
							msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
						kafkaMsgHandle(groupId, msg.Topic, msg.Key, msg.Value)
						consumer.MarkOffset(msg, "")
					}
				}
			}
		}()
	}

	for _, v := range configs {
		initClient(brokers, []string{v.Topic}, v.GroupId)
	}
}

func SendMsg(msg []byte, topic string) error {
	if len(msg) == 0 {
		return errors.New("send msg empty")
	}
	if topic == "" {
		return errors.New("topic empty")
	}

	producerMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	}

	select {
	case producer.Input() <- producerMsg:
	case <-time.After(10 * time.Second):
		return errors.New("kafka send msg timeout")
	}
	return nil
}

func SendKeyValMsg(key, msg []byte, topic string) error {
	if len(msg) == 0 {
		return errors.New("send msg empty")
	}
	if topic == "" {
		return errors.New("topic empty")
	}

	producerMsg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(msg),
	}

	select {
	case producer.Input() <- producerMsg:
	case <-time.After(10 * time.Second):
		return errors.New("kafka send msg timeout")
	}
	return nil
}

func kafkaMsgHandle(groupId string, topic string, msgKey []byte, msgVal []byte) {
	key := getConsumerKey(groupId, topic)
	handle := handleFuncMap[key]
	if handle != nil {
		handle(msgKey, msgVal)
	}
}

func RegisterHandler(groupId string, topic string, handle HandleFunc) {
	key := getConsumerKey(groupId, topic)
	handleFuncMap[key] = handle
}

func getConsumerKey(groupId, topic string) string {
	return fmt.Sprintf("%s-%s", topic, groupId)
}
