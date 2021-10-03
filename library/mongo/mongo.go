package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConfig struct {
	Primary string                      `yaml:"primary" json:"primary"`
	Source  map[string]*MongoItemConfig `yaml:"source" json:"source"`
}

type MongoItemConfig struct {
	Addr        string `yaml:"addr" json:"addr"`
	MaxPoolSize uint64 `yaml:"max_pool_size" json:"max_pool_size"`
	MinPoolSize uint64 `yaml:"min_pool_size" json:"min_pool_size"`
	ShowCmd     bool   `yaml:"show_cmd" json:"show_cmd"`
}

var (
	clientMap = make(map[string]*mongo.Client)
	primary   string
)

func Init(cfg *MongoConfig) {
	// check primary
	primary = cfg.Primary
	if primary == "" {
		panic("mongo config primary empty")
	}
	if _, ok := cfg.Source[primary]; !ok {
		panic("mongo config source primary empty")
	}

	// init client
	for k, v := range cfg.Source {
		opt := options.Client().ApplyURI(v.Addr)
		opt.SetMaxPoolSize(v.MaxPoolSize)
		opt.SetMinPoolSize(v.MinPoolSize)

		// print cmd
		if v.ShowCmd {
			monitor := &event.CommandMonitor{
				Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
					fmt.Printf("[mongo] cmd: %s\n", evt.Command)
				},
			}
			opt.SetMonitor(monitor)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, opt)
		if err != nil {
			panic(err)
		}
		clientMap[k] = client
	}

	// check connection
	for _, client := range clientMap {
		if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
			panic(err)
		}
	}
}

func GetClient(key ...string) *mongo.Client {
	if len(key) == 0 {
		return clientMap[primary]
	}
	return clientMap[key[0]]
}
