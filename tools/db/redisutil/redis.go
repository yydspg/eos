package redisutil

import (
	"context"
	"github.com/eos/tools/errs"
	"github.com/go-redis/redis/v9"
)

type Config struct {
	ClusterMode bool
	Address     []string
	Username    string
	Password    string
	MaxRetry    int
	DB          int
	PoolSize    int
}

func init() {
}

func NewRedisClient(ctx context.Context, config *Config) (redis.UniversalClient, error) {
	if len(config.Address) == 0 {
		return nil,errs.New("redis address is empty").Wrap()
	}	
	var cli redis.UniversalClient
	if config.ClusterMode || len(config.Address) > 1 {
		opt := &redis.ClusterOptions{
			Addrs:      config.Address,
			Username:   config.Username,
			Password:   config.Password,
			PoolSize:   config.PoolSize,
			MaxRetries: config.MaxRetry,
		}
		cli = redis.NewClusterClient(opt)
	} else {
		opt := &redis.Options{
			Addr:       config.Address[0],
			Username:   config.Username,
			Password:   config.Password,
			DB:         config.DB,
			PoolSize:   config.PoolSize,
			MaxRetries: config.MaxRetry,
		}
		cli = redis.NewClient(opt)
	}
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, errs.WrapMsg(err, "Redis Ping failed", "Address", config.Address, "Username", config.Username, "ClusterMode", config.ClusterMode)
	}
	return cli, nil
}
func Check(ctx context.Context, config *Config) error {
	client, err := NewRedisClient(ctx, config)
	if err != nil {
		return err
	}
	defer client.Close()

	// Ping the Redis server to check connectivity.
	if err := client.Ping(ctx).Err(); err != nil {
		return errs.WrapMsg(err, "Redis ping failed", "config", config)
	}

	return nil
}
