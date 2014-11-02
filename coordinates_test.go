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
