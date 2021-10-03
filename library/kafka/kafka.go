package kafka

import (
	"errors"
	"go-app/library/log"
	"time"

	"github.com/Shopify/sarama"
)

type KafkaConfig struct {
	Producer *ProducerConfig `yaml:"producer" json:"producer"`
}

type ProducerConfig struct {
	NameServers []string `yaml:"name_servers" json:"name_servers"`
}

var (
	producer sarama.AsyncProducer
)

func Init(cfg *KafkaConfig) {
	// producer
	pc := cfg.Producer
	if pc != nil {
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

		asyncProducer, err := sarama.NewAsyncProducer(pc.NameServers, config)
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
					log.Debug("kafka producer send success:", successes)
				}
			}
		}(asyncProducer)

		producer = asyncProducer
	}

	// consumer TODO
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
