package redis

import "github.com/go-redis/redis"

type RedisConfig struct {
	Single  *SingleItemConfig  `yaml:"single" json:"single"`
	Cluster *ClusterItemConfig `yaml:"cluster" json:"cluster"`
}

type SingleItemConfig struct {
	Primary string                         `yaml:"primary" json:"primary"`
	Source  map[string]*SingleClientConfig `yaml:"source" json:"source"`
}

type SingleClientConfig struct {
	Addr     string `yaml:"addr" json:"addr"`
	Password string `yaml:"password" json:"password"`
	PoolSize int    `yaml:"pool_size" json:"pool_size"`
}

type ClusterItemConfig struct {
	Primary string                          `yaml:"primary" json:"primary"`
	Source  map[string]*ClusterClientConfig `yaml:"source" json:"source"`
}

type ClusterClientConfig struct {
	Addrs    []string `yaml:"addrs" json:"addrs"`
	Password string   `yaml:"password" json:"password"`
	PoolSize int      `yaml:"pool_size" json:"pool_size"`
}

var (
	singleClientMap = make(map[string]*redis.Client)
	singlePrimary   string

	clusterClientMap = make(map[string]*redis.ClusterClient)
	clusterPrimary   string
)

func Init(cfg *RedisConfig) {
	// single
	sc := cfg.Single
	if sc != nil {
		singlePrimary = sc.Primary
		if singlePrimary == "" {
			panic("redis single config primary empty")
		}
		if _, ok := sc.Source[singlePrimary]; !ok {
			panic("redis single config source primary empty")
		}

		for k, v := range sc.Source {
			opt := &redis.Options{
				Addr:     v.Addr,
				Password: v.Password,
				PoolSize: v.PoolSize,
			}
			client := redis.NewClient(opt)
			singleClientMap[k] = client
		}

		for _, v := range singleClientMap {
			if err := v.Ping().Err(); err != nil {
				panic(err)
			}
		}
	}

	// cluster
	cc := cfg.Cluster
	if cc != nil {
		clusterPrimary = cc.Primary
		if clusterPrimary == "" {
			panic("redis cluster config primary empty")
		}
		if _, ok := cc.Source[clusterPrimary]; !ok {
			panic("redis cluster config source primary empty")
		}

		for k, v := range cc.Source {
			opt := &redis.ClusterOptions{
				Addrs:    v.Addrs,
				Password: v.Password,
				PoolSize: v.PoolSize,
			}
			client := redis.NewClusterClient(opt)
			clusterClientMap[k] = client
		}

		for _, v := range clusterClientMap {
			if err := v.Ping().Err(); err != nil {
				panic(err)
			}
		}
	}
}

func GetClient(key ...string) *redis.Client {
	if len(key) == 0 {
		return singleClientMap[singlePrimary]
	}
	return singleClientMap[key[0]]
}

func GetClusterClient(key ...string) *redis.ClusterClient {
	if len(key) == 0 {
		return clusterClientMap[singlePrimary]
	}
	return clusterClientMap[key[0]]
}
