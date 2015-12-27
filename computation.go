package main

import (
	"math"
)

const rEarth = 6372.8 // km

func Haversine(θ float64) float64 {
	return .5 * (1 - math.Cos(θ))
}

//Convert deg to rad
func DegToRad(lat, lon float64) Location {
	return Location{lat * math.Pi / 180, lon * math.Pi / 180}
}

// Convert Rad to deg
func RadToDeg(lat, lon float64) Location {
	return Location{lat * 180 / math.Pi, lon * 180 / math.Pi}
}

// return Hervesine distance in km
func HaversineDist(p1, p2 Location) Distance {
	return Distance{2 * rEarth * math.Asin(math.Sqrt(Haversine(p2.Lat-p1.Lat)+
		math.Cos(p1.Lat)*math.Cos(p2.Lat)*Haversine(p2.Lon-p1.Lon)))}
}

func MidPoint(prevs []Location) Location {
	var x, y, z, tot, midLon, midSquareRoot, midLat float64

	for i := range prevs {
		curr := DegToRad(prevs[i].Lat, prevs[i].Lon)
		x += math.Cos(curr.Lat) * math.Cos(curr.Lon)
		y += math.Cos(curr.Lat) * math.Sin(curr.Lon)
		z += math.Sin(curr.Lat)
		tot += 1.0
	}

	x = x / tot
	y = y / tot
	z = z / tot
	midLon = math.Atan2(y, x)
	midSquareRoot = math.Sqrt(x*x + y*y)
	midLat = math.Atan2(z, midSquareRoot)

	return RadToDeg(midLat, midLon)
}
