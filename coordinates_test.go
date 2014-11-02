package main

import "testing"

func TestGreatCircleDistance(t *testing.T) {
	c1 := &coordinates{Latitude: 37.7797, Longitude: -122.417}
	c2 := &coordinates{Latitude: 37.8116, Longitude: -122.242}
	d := GreatCircleDistance(c1, c2)
	if d != 9.806118100713892 {
		t.Error("expected distance to be 9.806118100713892", d)
	}
}

func TestParseCoordinates(t *testing.T) {
	c, err := ParseCoordinates("37.7797,-122.417")
	if err != nil {
		t.Error("unexpected error!")
	}
	if c.Latitude != 37.7797 {
		t.Error("expected latitude to be 37.7797", c)
	}
	if c.Longitude != -122.417 {
		t.Error("expected latitude to be -122.417", c)
	}

	c, err = ParseCoordinates("-122.417")
	if err == nil {
		t.Error("expected invalid format error")
	}
	c, err = ParseCoordinates("invalid,-122.417")
	if err == nil {
		t.Error("expected not a number error")
	}
}
