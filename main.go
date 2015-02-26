package main // import "github.com/CenturyLinkLabs/sample-go-adapter"

import (
	"log"
	"github.com/CenturyLinkLabs/pmxadapter"
)

func main() {
	adapter := new(SampleAdapter)
	server := pmxadapter.NewServer(adapter)

	log.Printf("Starting Sample Adapter")
	server.Start()
}
