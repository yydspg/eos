package redisutil

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type Config struct {
	ClusterMode bool
	Address     []string
	Username    string
	Password    string
	MaxRetry    int
	DB          int
	PoolSize    uint32
}

func init() {
}

func NewRedisClient(ctx context.Context, config *Config) (redis.UniversalClient, error) {
	return nil,nil
}
