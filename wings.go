package main

import (
	"github.com/naysaku/flame-wings/cmd"
	"math/rand"
	"time"
)

func main() {
	// Since we make use of the math/rand package in the code, especially for generating
	// non-cryptographically secure random strings we need to seed the RNG. Just make use
	// of the current time for this.
	rand.Seed(time.Now().UnixNano())

	// Execute the main binary code.
	cmd.Execute()
}
