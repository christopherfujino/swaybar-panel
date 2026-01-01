package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("{\"version\":1}\n[")
	for {
		now := time.Now()
		fmt.Printf("[{\"full_text\": \"%s\"},{\"full_text\": \"B(%s)\"}],", now.Format("15:04"), batteryPercent())
		time.Sleep(time.Second * 5)
	}
}
