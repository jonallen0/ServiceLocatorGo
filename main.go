package main

import (
	"log"
	"net/http"

	"github.com/jcpa/ServiceLocatorGo/ServiceLocatorGo"
)

func main() {
	ServiceLocatorGo.AddServices()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
