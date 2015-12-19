package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Compute(res http.ResponseWriter, req *http.Request) {
	var msg UserRequest

	decodeJson := json.NewDecoder(req.Body)
	err := decodeJson.Decode(&msg)

	if err != nil {
		panic(err)
	}

	prevs := msg.Prev
	currLoc := msg.Curr

	curr := MidPoint(prevs)
	latMid := curr.Lat
	lonMid := curr.Lon

	currDist := Hervesine(latMid, lonMid, currLoc.Lat, currLoc.Lon)

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(res).Encode(currDist); err != nil {
		log.Fatal(err)
	}
}
