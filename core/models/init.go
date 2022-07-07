package models

import (
	"cloud_disk/core/internal/config"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Init(c config.Config) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", c.Mysql.DataSource)
	if err != nil {
		log.Printf("Xorm New Engine Error: %v", err)
	}
	return engine
}
