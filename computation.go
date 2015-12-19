package main

import (
	"math"
)

func MidPoint(prevs []Location) Location {
	var x, y, z, tot, midLon, midSquareRoot, midLat float64

	for i := range prevs {
		lat := prevs[i].Lat * math.Pi / 180
		lon := prevs[i].Lon * math.Pi / 180
		x += math.Cos(lat) * math.Cos(lon)
		y += math.Cos(lat) * math.Sin(lon)
		z += math.Sin(lat)
		tot += 1.0
	}

	x = x / tot
	y = y / tot
	y = z / tot

	midLon = math.Atan2(y, x)
	midSquareRoot = math.Sqrt(x*x + y*y)
	midLat = math.Atan2(z, midSquareRoot)

	curr := Location{
		Lat: DegToRad(midLat),
		Lon: DegToRad(midLon),
	}

	return curr
}

func DegToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func Hervesine(lat1 float64, lat2 float64, lon1 float64, lon2 float64) Distance {
	var dLat, dLon, a, c, d float64

	earthRad := 6371.0

	dLat = DegToRad(math.Abs(lat1 - lat2))
	dLon = DegToRad(math.Abs(lon1 - lon2))
	a = math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(DegToRad(lat1))*math.Cos(DegToRad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d = earthRad * c

	distance := Distance{
		CurrDistance: d,
	}

	return distance
}
