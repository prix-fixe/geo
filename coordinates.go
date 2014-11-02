package main

import "math"

const (
	earthRadius      = 3959 // in miles
	degreesPerRadian = math.Pi / 180.0
)

type coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Calculates the Haversine distance between two coordinates in miles.
// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
func GreatCircleDistance(c1 *coordinates, c2 *coordinates) float64 {
	lat1 := c1.Latitude * degreesPerRadian
	lat2 := c2.Latitude * degreesPerRadian

	dLat := (c2.Latitude - c1.Latitude) * degreesPerRadian
	dLon := (c2.Longitude - c1.Longitude) * degreesPerRadian

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}
