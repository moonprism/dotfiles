package main

import (
	"os/exec"
	"time"
	"bytes"
	"strings"
)

func getPower() string {
	path = "/sys/class/power_supply/BAT0/"
	status = "status"
	capacity = "capacity"
	return ""
}

var dwmVersion string

func getDwmVersion() string {
	if dwmVersion != "" {
		return dwmVersion
	}
	// 为啥是stderr呢，想不明白
	cmd := exec.Command("dwm", "-v")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Run()
	dwmVersion = stderr.String()
	return dwmVersion
}

func main() {
	for {
		status := []string{
			time.Now().Local().Format("1-2 15:04"),
			getDwmVersion(),
		}

		s := " " + strings.Join(status, " ")

		exec.Command("xsetroot", "-name", s).Run()
		var now = time.Now()
		time.Sleep(now.Truncate(time.Second).Add(time.Second).Sub(now))
	}
}
