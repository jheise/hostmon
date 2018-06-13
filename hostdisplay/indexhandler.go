package main

import (
	// standard
	"encoding/json"
	// "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sort"

	// local
	"github.com/jheise/hostmon/hostutil"

	// external
	"github.com/garyburd/redigo/redis"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := redis.Dial("tcp", redisaddr+":"+redisport)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		panic(err)
	}
	// sort keys
	sort.Strings(keys)

	// hosts := []hostutil.DockerList{}
	hosts := []hostutil.Host{}

	for _, hostname := range keys {
		data, err := redis.String(conn.Do("GET", hostname))
		if err != nil {
			panic(err)
		}

		var host hostutil.Host
		json.Unmarshal([]byte(data), &host)

		hosts = append(hosts, host)
	}

	hostdata := HostMonData{hosts}

	// load template and return html
	template_bytes, err := ioutil.ReadFile(templates + "/hostmon.template")
	if err != nil {
		panic(err)
	}
	template_data := string(template_bytes)

	tmpl := template.New("hostmon-template")
	tmpl, err = tmpl.Parse(template_data)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, hostdata)

	// fmt.Fprintf(w, "keys are %s", hosts)
}

// {"Hostname":"alita","Containers":[{"Name":"/zen_cori","Image":"ubuntu","Status":"exited","Ports":null}]}
