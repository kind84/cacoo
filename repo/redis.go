package repo

import (
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/kind84/cacoo"
)

type redisRepo struct {
	redis *redis.Pool
}

// NewRedisRepo returns a pointer to a new repository using Redis.
func NewRedisRepo(host string) cacoo.Repo {
	redisClient := &redis.Pool{
		MaxIdle:   5,
		MaxActive: 5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host)
		},
	}
	return &redisRepo{redisClient}
}

func (r redisRepo) Save(k string, v interface{}) {
	conn := r.redis.Get()
	defer conn.Close()

	conn.Do("SET", k, v)
	log.Printf("Inserted key %s with value: %v\n", k, v)
}
