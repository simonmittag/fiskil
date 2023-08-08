package main

import (
	"fmt"
	"github.com/simonmittag/fiskil"
)

func main() {
	fmt.Print("\nloading..")
	s := fiskil.Server{}
	s.Bootstrap()
}
