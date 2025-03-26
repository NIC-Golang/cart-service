package tests

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

type MockRedis struct {
	Client *redis.Client
	Mini   *miniredis.Miniredis
}

func NewMockRedis() (*MockRedis, error) {
	mini, err := miniredis.Run()
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(&redis.Options{
		Addr: mini.Addr(),
	})
	return &MockRedis{Client: client, Mini: mini}, nil
}

func (m *MockRedis) Set(key string, value string) error {
	return m.Client.Set(key, value, 0).Err()
}

func (m *MockRedis) Get(key string) (string, error) {
	return m.Client.Get(key).Result()
}

func TestMockRedis(t *testing.T) {
	mockRedis, err := NewMockRedis()
	assert.NoError(t, err)

	key, value := "foo", "bar"

	err = mockRedis.Set(key, value)
	assert.NoError(t, err)

	result, err := mockRedis.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, value, result)
}
