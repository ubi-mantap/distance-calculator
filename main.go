package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", Compute)
	fmt.Println("listening...")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}

}
