package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	
	if err != nil {
		l := log.New(os.Stderr, "", 0)
		l.Println("log ntp")
	}

	fmt.Println(time)
}
