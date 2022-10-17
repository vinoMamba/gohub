package verifycode

import (
	"time"

	"github.com/vinoMamba/gohub/pkg/redis"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) bool {
	return s.RedisClient.Set(s.KeyPrefix+key, value, time.Minute*time.Duration(15))
}

func (s *RedisStore) Get(key string) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	return val
}

func (s *RedisStore) Verify(key string, answer string) bool {
	v := s.Get(s.KeyPrefix + key)
	return v == answer
}
