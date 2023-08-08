package main

import (
	"fmt"
	"github.com/simonmittag/fiskil"
)

// cli program 'server'
func main() {
	fmt.Print("\nloading..")
	s := fiskil.Server{}
	s.Bootstrap()
}
