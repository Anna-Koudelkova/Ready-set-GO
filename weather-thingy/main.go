package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anna-Koudelkova/Ready-set-GO/weather-thingy/handlers"
)

func main() {
	fileserver := http.FileServer(http.Dir("./pages"))
	//  here is where the different handlers go and server starts
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/weatherpage", handlers.WeatherHandler)

	fmt.Println("The most beautiful server is up and running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
