package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var c config
var initialized bool = false

func setupTest(t *testing.T) {
	if initialized {
		return
	}

	initialized = true
	env := "test"
	err := c.read(&env)
	if err != nil {
		t.Error("unable to read config file", err)
	}

	err = initDb(c)
	if err != nil {
		t.Error("error initializing the db", err)
	}
}

// lookup test
func TestValidLookup(t *testing.T) {
	setupTest(t)

	url := fmt.Sprintf("http://localhost:%v/lookup?ip=67.188.210.242", c.Port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("error creating request", err)
	}

	res := httptest.NewRecorder()
	handleLookup(res, req)

	var l location
	err = json.NewDecoder(res.Body).Decode(&l)
	if err != nil {
		t.Error("json parse error", err)
	}

	if l.City != "San Francisco" {
		t.Error("expected city to be 'San Francisco'", l.City)
	}
	if l.Zipcode != "94102" {
		t.Error("expected zipcode to be '94102'", l.Zipcode)
	}
	if l.Timezone != "America/Los_Angeles" {
		t.Error("expected time zone to be 'America/Los_Angeles'", l.Timezone)
	}
	if l.Coordinates.Latitude != 37.7794 {
		t.Error("expected latitude to be 37.7794", l.Coordinates.Latitude)
	}
	if l.Coordinates.Longitude != -122.417 {
		t.Error("expected longitude to be -122.417", l.Coordinates.Longitude)
	}
}

func TestLookup404(t *testing.T) {
	setupTest(t)

	url := fmt.Sprintf("http://localhost:%v/lookup?ip=67.188.210.242", c.Port)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		t.Error("error creating request", err)
	}

	res := httptest.NewRecorder()
	handleLookup(res, req)
	if res.Code != 404 {
		t.Error("StatusCode should be 404", res.Code)
	}
}

func TestInvalidIp(t *testing.T) {
	setupTest(t)

	url := fmt.Sprintf("http://localhost:%v/lookup?ip=invalid", c.Port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("error creating request", err)
	}

	res := httptest.NewRecorder()
	handleLookup(res, req)
	if res.Code != 400 {
		t.Error("StatusCode should be 400", res.Code)
	}
}

// distance test
func TestValidDistance(t *testing.T) {
	setupTest(t)

	url := fmt.Sprintf("http://localhost:%v/distance?latlng1=37.7797,-122.417&latlng2=37.8116,-122.242", c.Port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("error creating request", err)
	}

	res := httptest.NewRecorder()
	handleDistance(res, req)

	if res.Body.String() != "{\"distance\": 9.81}" {
		t.Error("expected distance to be 9.81", res.Body.String())
	}
}

func TestDistance404(t *testing.T) {
	setupTest(t)

	url := fmt.Sprintf("http://localhost:%v/distance?latlng1=37.7797,-122.417&latlng2=37.8116,-122.242", c.Port)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		t.Error("error creating request", err)
	}

	res := httptest.NewRecorder()
	handleDistance(res, req)
	if res.Code != 404 {
		t.Error("StatusCode should be 404", res.Code)
	}
}

func TestInvalidLatLng(t *testing.T) {
	setupTest(t)

	url := fmt.Sprintf("http://localhost:%v/distance?latlng1=invalid,-122.417&latlng2=37.8116,-122.242", c.Port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("error creating request", err)
	}

	res := httptest.NewRecorder()
	handleDistance(res, req)
	if res.Code != 400 {
		t.Error("StatusCode should be 400", res.Code)
	}
}
