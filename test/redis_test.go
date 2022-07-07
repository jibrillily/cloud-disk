package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()

//设置redis连接
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestSetValue(t *testing.T) {
	//最后一个为过期时间
	err := rdb.Set(ctx, "key", "value", time.Second*10).Err()
	if err != nil {
		t.Error(err)
	}
}

func TestGetValue(t *testing.T) {
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
}
