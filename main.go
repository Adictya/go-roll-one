package main

import (
	"flag"
	"fmt"
	"os"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) < 2 {
		fmt.Println(rand.Intn(6)+1)
		os.Exit(0)
	} else {
		n := flag.Int("n",6,"Number of sides in a die")
		fmt.Println(rand.Intn(*n)+1)
		os.Exit(0)
	}
}
