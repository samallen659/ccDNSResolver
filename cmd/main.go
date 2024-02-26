package main

import (
	"flag"
	"github.com/samallen659/ccDNSResolver/internal/connection"
	"log"
)

func main() {
	var server = flag.String("server", "8.8.8.8", "DNS server to send queries too")
	var hostname = flag.String("hostname", "", "Hostname to query")
	flag.Parse()

	if *hostname == "" {
		log.Fatal("Hostname cannot be blank")
	}
	c, err := connection.NewClient(*server)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Resolve(*hostname)
	if err != nil {
		log.Fatal(err)
	}
	resp.Print()
}
