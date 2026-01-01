package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PREFIX = "/sys/class/power_supply/qcom-battmgr-bat/energy_"

func batteryPart(part string) float64 {
	diskPath := fmt.Sprintf("%s%s", PREFIX, part)
	bytes, err := os.ReadFile(diskPath)
	if err != nil {
		return -1; // TODO
	}
	floatValue, err := strconv.ParseFloat(strings.TrimSpace(string(bytes)), 64)
	if err != nil {
		return -1; // TODO
	}

	return floatValue
}

func batteryPercent() string {
	full := batteryPart("full")
	current := batteryPart("now")
	empty := batteryPart("empty")
	return fmt.Sprintf("%.1f%%", 100*(current-empty)/(full-empty))
}
