package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
* curl -X POST -H "Content-Type: application/json" -d '{ "array":[ {"X":1, "Y":4},{"X":1, "Y":4}]}' http://127.0.0.1:8082/
 */
func main() {

	http.HandleFunc("/", Compute)
	fmt.Println("listening...")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}

}
