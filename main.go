package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const PREFIX = "/sys/class/power_supply/qcom-battmgr-bat/energy_"

func main() {
	fmt.Printf("{\"version\":1}\n[")
	for {
		now := time.Now()
		fmt.Printf("[{\"full_text\": \"%s\"},{\"full_text\": \"B(%s)\"}],", now.Format("15:04"), batteryPercent())
		time.Sleep(time.Second * 5)
	}
}

func batteryPart(part string) float64 {
	diskPath := fmt.Sprintf("%s%s", PREFIX, part)
	bytes, err := os.ReadFile(diskPath)
	if err != nil {
		panic(err)
	}
	floatValue, err := strconv.ParseFloat(strings.TrimSpace(string(bytes)), 64)
	if err != nil {
		panic(err)
	}

	return floatValue
}

func batteryPercent() string {
	full := batteryPart("full")
	current := batteryPart("now")
	empty := batteryPart("empty")
	return fmt.Sprintf("%.1f%%", 100*(current-empty)/(full-empty))
}
