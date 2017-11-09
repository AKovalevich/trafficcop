package main

import (
	"os"

	"github.com/AKovalevich/trafficcop/cmd/trafficcop"
)

// Run application
func main() {
	os.Exit(trafficcop.Run(os.Args[1:]))
}
