package main

import "github.com/gomodule/redigo/redis"

var host string = "127.0.0.1"
var port string = "6379"
var password string = "azerty123"
var redisServer string = host + ":" + port
var connectionError error
var connection redis.Conn
var valeur string
var cle string

// RedisGetKey function
func RedisGetKey(cle string) string {
	connection, connectionError = redis.Dial("tcp", redisServer)
	if connectionError != nil {
		panic(connectionError)
	}
	_, connectionError = connection.Do("AUTH", password)
	if connectionError != nil {
		panic(connectionError)
	}
	valeur, connectionError = redis.String(connection.Do("GET", cle))
	defer connection.Close()
	return valeur
}

// RedisSetKey function
func RedisSetKey(cle string, valeur string) error {
	connection, connectionError = redis.Dial("tcp", redisServer)
	if connectionError != nil {
		panic(connectionError)
	}
	_, connectionError = connection.Do("AUTH", password)
	if connectionError != nil {
		panic(connectionError)
	}
	valeur, connectionError = redis.String(connection.Do("SET", cle, valeur))
	if connectionError != nil {
		panic(connectionError)
	}
	defer connection.Close()
	return nil
}
