package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/**
* curl -X POST -H "Content-Type: application/json" -d '{"current":{"X":0,"Y":0},"previous":[ {"X":1, "Y":4},{"X":1, "Y":4}]}' http://127.0.0.1:8082/
 */
func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8082"
	}

	http.HandleFunc("/", Compute)
	fmt.Println("listening at " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
