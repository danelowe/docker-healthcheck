package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	healthCheckURL := "http://localhost"
	if len(os.Args) > 2 {
		fmt.Println("Must have a single argument")
	} else if len(os.Args) == 2 {
		healthCheckURL = os.Args[1]
	}
	_, err := http.Get(healthCheckURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
