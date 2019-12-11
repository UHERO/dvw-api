package data

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"log"
)

type Cache struct {
	Prefix string
	Pool *redis.Pool
	TTL  int
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func CreateCache(prefix string, pool *redis.Pool, ttlMin int) (r *Cache) {
	r = &Cache{Prefix: prefix, Pool: pool, TTL: 60 * ttlMin} // actual TTL is in seconds
	c := r.Pool.Get()
	defer c.Close()
	_, err := c.Do("PING")
	if err != nil {
		log.Printf("**** Cannot contact Redis server (%v). No caching!", err.Error())
		return nil
	}
	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func (r *Cache) GetCache(key string) ([]byte, error) {
	c := r.Pool.Get()
	defer c.Close()
	if r.Prefix != "" {
		key = r.Prefix + "_" + key
	}
	value, err := c.Do("GET", key)
	if err != nil {
		log.Printf("Redis error on GET: %v", err)
		return nil, err
	}
	if value == nil {
		//log.Printf("Redis cached val nil on GET: %v", err)
		return nil, err
	}
	log.Printf("Redis GET: %s", key)
	return value.([]byte), err
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func (r *Cache) SetCache(key string, value []byte) (err error) {
	c := r.Pool.Get()
	defer c.Close()
	if r.Prefix != "" {
		key = r.Prefix + "_" + key
	}
	c.Send("MULTI")
	c.Send("SET", key, value)
	c.Send("EXPIRE", key, r.TTL)
	response, err := redis.Values(c.Do("EXEC"))
	if err != nil {
		log.Printf("Redis error on SET or EXPIRE: %v", err)
		return
	}
	var setResponse string
	var expireResponse int
	if _, err := redis.Scan(response, &setResponse, &expireResponse); err != nil {
		log.Print("Error on scan of redis response")
	}
	if setResponse != "OK" {
		err = errors.New("did not get OK from Redis SET")
		log.Print(err)
		return
	}
	if expireResponse != 1 {
		log.Printf("Did not set expiration to %v", r.TTL)
	}
	log.Printf("Redis SET: %s", key)
	return
}
