package adapter

import (
	"context"
	"github.com/hetiansu5/msredis"
	"testing"
	"time"
)

func TestNewRedis_Set(t *testing.T) {
	redis := NewRedis(msredis.Group("cache", "cache_slave1"))
	err := redis.Set(context.TODO(), "tt01", 3, time.Second*10)
	if err != nil {
		t.Error(err)
	}
	v, err1 := redis.Get(context.TODO(), "tt01")
	if err1 != nil {
		t.Error(err1)
	}
	n, ok := v.(string)
	if !ok {
		t.Error("get not int value")
	}
	if n != "3" {
		t.Error("get wrong value")
	}
}
