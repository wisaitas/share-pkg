package redis

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockRedis struct {
	mock.Mock
}

func NewMockRedis() *MockRedis {
	return &MockRedis{}
}

func (m *MockRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	ret := m.Called(ctx, key, value, expiration)
	return ret.Error(0)
}

func (m *MockRedis) Get(ctx context.Context, key string) (string, error) {
	ret := m.Called(ctx, key)
	return ret.String(0), ret.Error(1)
}

func (m *MockRedis) Del(ctx context.Context, keys ...string) error {
	ret := m.Called(ctx, keys)
	return ret.Error(0)
}

func (m *MockRedis) Exists(ctx context.Context, keys ...string) (bool, error) {
	ret := m.Called(ctx, keys)
	return ret.Bool(0), ret.Error(1)
}
