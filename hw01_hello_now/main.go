package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	logger := log.New(os.Stderr, "", 0)

	curTime := time.Now()
	exTime, err := ntp.Time("0.ru.pool.ntp.org")
	if err != nil {
		logger.Fatal("Error: can't get network time: ", err)
	}

	fmt.Println("current time:", curTime.Round(time.Second))
	fmt.Println("exact time:", exTime.Round(time.Second))
}
