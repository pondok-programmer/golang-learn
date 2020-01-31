package main

import (
	"fmt"

	"github.com/pondok-programmer/golang-learn/sidik"
	"github.com/pondok-programmer/golang-learn/syarif"
	"github.com/pondok-programmer/golang-learn/roihan"
)

func main() {
	fmt.Println("Start golang-learn")
  fmt.Println("Run Syarif")
	syarif.Main()
  fmt.Println()
  fmt.Println("Run Sidik")
	sidik.Main()
  fmt.Println()
  fmt.Println("Run Roihan")
	roihan.Main()
  fmt.Println()
	fmt.Println("End golang-learn")
}
