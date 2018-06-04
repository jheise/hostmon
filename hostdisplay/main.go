package main

import (
	// standard
	"flag"
	"fmt"
	"net/http"

	// external
	"github.com/gorilla/mux"
)

var (
	address   string
	port      string
	redisaddr string
	redisport string
)

func init() {
	flag.StringVar(&address, "address", "0.0.0.0", "address to bind to")
	flag.StringVar(&port, "port", "9999", "port to bind to")
	flag.StringVar(&redisaddr, "redisaddr", "redis", "redis host to connect to")
	flag.StringVar(&redisport, "redisport", "6379", "redis port to use")
	flag.Parse()
}

func main() {
	router := mux.NewRouter()

	// handle static requests
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// handle index
	router.HandleFunc("/", IndexHandler).Methods("GET")

	fmt.Printf("Listening on %s ...\n", address+":"+port)
	http.ListenAndServe(address+":"+port, router)
}
