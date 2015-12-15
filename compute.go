package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Location struct {
	X, Y float64
}

type PreviousLocations struct {
	Previous []Location `json:"array"`
}

func main() {
	http.HandleFunc("/", compute)
	fmt.Println("listening...")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func midllePoint(x float64, y float64, tot float64) Location {
	currX := x / tot
	currY := y / tot

	curr := Location{
		X: currX,
		Y: currY,
	}
	return curr
}

func compute(res http.ResponseWriter, req *http.Request) {
	var msg PreviousLocations
	var totX float64
	var totY float64
	var tot float64

	decodeJson := json.NewDecoder(req.Body)
	err := decodeJson.Decode(&msg)

	if err != nil {
		panic(err)
	}
	prevs := msg.Previous
	for i := range prevs {
		totX += prevs[i].X
		totY += prevs[i].Y
		tot += 1.0
	}

	curr := midllePoint(totX, totY, tot)

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(res).Encode(curr); err != nil {
		log.Fatal(err)
	}
}
