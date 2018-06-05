package main

import (
	// std
	"flag"
	// "fmt"
	"encoding/json"
	// "os"
	"time"

	// local
	"github.com/jheise/hostmon/hostutil"

	// external
	"github.com/fsouza/go-dockerclient"
)

var (
	client   *docker.Client
	port     string
	address  string
	interval int
	identity string

	storage StorageAgent
)

func init() {
	flag.StringVar(&address, "address", "redis", "address to redis server")
	flag.StringVar(&port, "port", "6379", "redis port")
	flag.IntVar(&interval, "interval", 3, "update interval")
	flag.StringVar(&identity, "identity", "unnamedAgent", "agent identity")
	flag.Parse()
}

func main() {
	endpoint := "unix:///var/run/docker.sock"
	localclient, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	client = localclient

	// build storage agent
	storage, err = newRedisStorageAgent(address, port, identity, interval+5)
	if err != nil {
		panic(err)
	}

	for {
		hostdata := hostutil.Host{}
		hostdata.Hostname = identity

		docker, err := check_docker()
		if err != nil {
			panic(err)
		}

		hostdata.Containers = docker.Containers
		jdata, err := json.Marshal(hostdata)
		if err != nil {
			panic(err)
		}

		err = storage.Process(string(jdata))
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
