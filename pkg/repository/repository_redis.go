package repository

import (
	"fmt"
	"time"

	"github.com/willyrgf/rerest/config"
	"github.com/garyburd/redigo/redis"
)

var (
	repo *redis.Pool
)

func initRepo() {
	c := *config.RerestConf
	server := fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort)

	pool := &redis.Pool{

		MaxIdle:     10,
		MaxActive:   10000,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	repo = pool
}

// Get repo
func Get() *redis.Pool {
	if repo == nil {
		initRepo()
	}
	return repo
}

// GetConn get a conn from pool select data database
func GetConn() (conn redis.Conn) {
	repo = Get()
	conn = repo.Get()
	conn.Do("SELECT", config.RerestConf.RedisDatabaseData)
	return
}

// GetConnToControlDB get a conn from pool selecting control database
func GetConnToControlDB() (conn redis.Conn) {
	repo = Get()
	conn = repo.Get()
	conn.Do("SELECT", config.RerestConf.RedisDatabaseControl)
	return
}
