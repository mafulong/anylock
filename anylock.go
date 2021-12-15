package anylockgo

import (
	"context"
	"time"
)

type AnyLock interface {
	Lock(ctx context.Context) error
	Unlock(ctx context.Context) error
	LockWithTTL(ctx context.Context, ttl time.Duration) error
	LockWithTTLAndTimeout(ctx context.Context, ttl time.Duration, timeout time.Duration) error
}

// todo
type RedisAnyLock struct {
}

// todo
//func NewRedisAnyLock(ctx context.Context, lockKey string, redisClient redis.Client) *RedisAnyLock {
//	return &RedisAnyLock{}
//}

func (receiver *RedisAnyLock) Lock(ctx context.Context) error {
	return nil
}

func (receiver *RedisAnyLock) Unlock(ctx context.Context) error {
	return nil
}
func (receiver *RedisAnyLock) LockWithTTL(ctx context.Context, ttl time.Duration) error {
	return nil
}
func (receiver *RedisAnyLock) LockWithTTLAndTimeout(ctx context.Context, ttl time.Duration, timeout time.Duration) error {
	return nil
}
