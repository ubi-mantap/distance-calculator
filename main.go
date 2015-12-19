package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/**
* curl -X POST -H "Content-Type: application/json" -d '{"current":{"Lat":-0.626466,"Lon":100.1158333},"previous":[ {"Lat":-7.8981666, "Lon":110.2226681},{"Lat":-6.9034443, "Lon":107.5731165},{"Lat":-0.3027606, "Lon":100.3544704}]}' http://127.0.0.1:8082/
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
