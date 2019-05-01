package repo

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// RedisRepo implements Repo using Redis.
type RedisRepo struct {
	redis *redis.Pool
}

// NewRedisRepo returns a pointer to a new repository using Redis.
func NewRedisRepo(host string) *RedisRepo {
	redisClient := &redis.Pool{
		MaxIdle:   5,
		MaxActive: 5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host)
		},
	}
	return &RedisRepo{redisClient}
}

// Save a key/value pair into Redis.
func (r RedisRepo) Save(k interface{}, v interface{}) {
	conn := r.redis.Get()
	defer conn.Close()

	conn.Do("SET", k, v)
	log.Printf("Inserted key %s\n", k)
}

// SaveSet creates/adds values to a set in Redis.
func (r RedisRepo) SaveSet(k interface{}, v ...interface{}) {
	conn := r.redis.Get()
	defer conn.Close()

	args := []interface{}{k}
	args = append(args, v...)
	res, _ := conn.Do("SADD", args...)
	if res != 0 {
		log.Printf("Values added to set [%s]\n", k)
	}
}

// Get an item from repository at the given key.
func (r RedisRepo) Get(k interface{}) string {
	conn := r.redis.Get()
	defer conn.Close()

	res, _ := redis.String(conn.Do("GET", k))
	return res
}

// GetASet returns the set stored at the given key.
func (r RedisRepo) GetASet(k interface{}) []string {
	conn := r.redis.Get()
	defer conn.Close()

	res, _ := redis.Strings(conn.Do("SMEMBERS", k))
	return res
}
