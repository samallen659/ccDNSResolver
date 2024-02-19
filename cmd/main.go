package main

import (
	"log"

	"github.com/samallen659/ccDNSResolver/internal/connection"
)

func main() {
	c, err := connection.NewClient("8.8.8.8")
	if err != nil {
		log.Fatal(err)
	}

	err = c.Resolve("www.northlincs.gov.uk")
	if err != nil {
		log.Fatal(err)
	}
}
