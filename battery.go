package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PREFIX = "/sys/class/power_supply/qcom-battmgr-bat/energy_"

func batteryPart(part string) (float64, error) {
	diskPath := fmt.Sprintf("%s%s", PREFIX, part)
	bytes, err := os.ReadFile(diskPath)
	if err != nil {
		return 0, err;
	}
	floatValue, err := strconv.ParseFloat(strings.TrimSpace(string(bytes)), 64)
	if err != nil {
		return 0, err;
	}

	return floatValue, nil
}

func batteryPercent() string {
	full, err := batteryPart("full")
	if err != nil {
		panic(err)
	}
	current, err := batteryPart("now")
	if err != nil {
		panic(err)
	}
	empty, err := batteryPart("empty")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%.1f%%", 100*(current-empty)/(full-empty))
}
