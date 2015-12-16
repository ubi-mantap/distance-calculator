package main

type Location struct {
	X, Y float64
}

type UserRequest struct {
	Curr Location   `json:"current"`
	Prev []Location `json:"previous"`
}

type Distance struct {
	CurrDistance float64
}
