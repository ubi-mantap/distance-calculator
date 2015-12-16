package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
