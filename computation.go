package main

import (
	"math"
)

func MidPoint(x float64, y float64, tot float64) Location {
	currX := x / tot
	currY := y / tot

	curr := Location{
		X: currX,
		Y: currY,
	}
	return curr
}

func CurrentDistance(currX float64, currY float64, midX float64, midY float64) Distance {
	x := math.Abs(currX - midX)
	y := math.Abs(currY - midY)
	xPow := math.Pow(x, 2)
	yPow := math.Pow(y, 2)
	dist := xPow + yPow
	sqrtDist := math.Sqrt(dist)
	currDist := Distance{
		CurrDistance: sqrtDist,
	}
	return currDist
}
