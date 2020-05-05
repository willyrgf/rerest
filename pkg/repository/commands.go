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

// GetFromList is a layer to get members from a list
func GetFromList(key string) (members []string, err error) {
	conn := GetConn()
	defer conn.Close()
	members, err = redis.Strings(conn.Do("LRANGE", key, 0, -1))
	return
}
