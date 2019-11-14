package repository

import (
	"github.com/garyburd/redigo/redis"
)

// GetFromSet is a layer to get members from a set
func GetFromSet(key string) (members []string, err error) {
	conn := GetConn()
	defer conn.Close()
	members, err = redis.Strings(conn.Do("SMEMBERS", key))
	return
}
