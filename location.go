package main

import "github.com/oschwald/geoip2-golang"

type location struct {
	City        string      `json:"city"`
	Zipcode     string      `json:"zipcode"`
	Timezone    string      `json:"timezone"`
	Coordinates coordinates `json:"coordinates"`
}

func NewLocation(record *geoip2.City) location {
	return location{
		City:     record.City.Names["en"],
		Zipcode:  record.Postal.Code,
		Timezone: record.Location.TimeZone,
		Coordinates: coordinates{
			Latitude:  record.Location.Latitude,
			Longitude: record.Location.Longitude,
		},
	}
}
