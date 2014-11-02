package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/oschwald/geoip2-golang"
)

var db *geoip2.Reader

func main() {
	env := flag.String("c", "development", "the program config environment")
	flag.Parse()

	var c config
	err := c.read(env)
	if err != nil {
		log.Fatal("config", err)
	}

	err = initDb(c)
	if err != nil {
		log.Fatal("initDb", err)
	}

	http.HandleFunc("/lookup", handleLookup)
	http.HandleFunc("/distance", handleDistance)
	err = http.ListenAndServe(fmt.Sprintf(":%v", c.Port), nil)
	if err != nil {
		log.Fatal("unable to start server", err)
	}
}

func initDb(c config) error {
	var err error
	db, err = geoip2.Open(c.DbFile)
	return err
}

func handleLookup(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.NotFound(res, req)
		return
	}

	sIp := req.URL.Query().Get("ip")
	ip := net.ParseIP(sIp)
	if ip == nil {
		http.Error(res, fmt.Sprintf("Invalid IP address %v", sIp), 400)
		return
	}

	record, err := db.City(ip)
	if err != nil {
		http.Error(res, fmt.Sprintf("Unexpected error: %v", err), 500)
		return
	}

	location := NewLocation(record)

	data, err := json.Marshal(location)
	if err != nil {
		http.Error(res, fmt.Sprintf("Unexpected error: %v", err), 500)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(data)
}

func handleDistance(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.NotFound(res, req)
		return
	}

	latlng1 := req.URL.Query().Get("latlng1")
	latlng2 := req.URL.Query().Get("latlng2")
	if latlng1 == "" || latlng2 == "" {
		http.Error(res, "latlng1 and latlng2 are required", 400)
		return
	}

	c1, err := ParseCoordinates(latlng1)
	if err != nil {
		http.Error(res, fmt.Sprintf("invalid latlng1 format %v:%v", latlng1, err), 400)
		return
	}

	c2, err := ParseCoordinates(latlng2)
	if err != nil {
		http.Error(res, fmt.Sprintf("invalid latlng2 format %v:%v", latlng2, err), 400)
		return
	}

	d := GreatCircleDistance(c1, c2)

	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(fmt.Sprintf("{\"distance\": %.2f}", d)))
}
