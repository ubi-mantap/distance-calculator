package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
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

func Compute(res http.ResponseWriter, req *http.Request) {
	var msg UserRequest
	var totX float64
	var totY float64
	var tot float64

	decodeJson := json.NewDecoder(req.Body)
	err := decodeJson.Decode(&msg)

	if err != nil {
		panic(err)
	}

	prevs := msg.Prev
	currLoc := msg.Curr

	for i := range prevs {
		totX += prevs[i].X
		totY += prevs[i].Y
		tot += 1.0
	}

	curr := MidPoint(totX, totY, tot)
	xMid := curr.X
	yMid := curr.Y

	currDist := CurrentDistance(xMid, yMid, currLoc.X, currLoc.Y)

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(res).Encode(currDist); err != nil {
		log.Fatal(err)
	}
}
