package main

type Location struct {
	X, Y float64
}

type PreviousLocations struct {
	Previous []Location `json:"array"`
}
