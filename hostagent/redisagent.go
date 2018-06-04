package main

import (
	// external
	"github.com/garyburd/redigo/redis"
)

type RedisStorageAgent struct {
	host  string
	port  string
	ident string
	ttl   int

	rconn redis.Conn
}

func newRedisStorageAgent(host string, port string, ident string, ttl int) (*RedisStorageAgent, error) {
	newAgent := new(RedisStorageAgent)
	newAgent.host = host
	newAgent.port = port
	newAgent.ident = ident
	newAgent.ttl = ttl

	// try to connect to redis server
	rconn, err := redis.Dial("tcp", host+":"+port)
	if err != nil {
		return newAgent, err
	}
	newAgent.rconn = rconn

	return newAgent, nil
}

func (self *RedisStorageAgent) Process(data string) error {
	// store data
	_, err := self.rconn.Do("SET", self.ident, data)
	if err != nil {
		return err
	}

	// set ttl
	_, err = self.rconn.Do("EXPIRE", self.ident, self.ttl)

	return err
}
