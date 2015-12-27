package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// curl -X POST -H "Content-Type: application/json" -d '{"current":{"Lat":36.12,"Lon":86.67},"previous":[{"Lat":36.12, "Lon":86.67},{"Lat":36.12, "Lon":86.67}]}' http://127.0.0.1:8082/

func Compute(res http.ResponseWriter, req *http.Request) {
	var msg UserRequest

	decodeJson := json.NewDecoder(req.Body)
	err := decodeJson.Decode(&msg)

	if err != nil {
		panic(err)
	}

	currLoc := msg.Curr
	prevLoc := msg.Prev

	mid := MidPoint(prevLoc)
	currDist := HaversineDist(DegToRad(mid.Lat, mid.Lon), DegToRad(currLoc.Lat, currLoc.Lon))
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(res).Encode(currDist); err != nil {
		log.Fatal(err)
	}
}
