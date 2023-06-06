package main

import (
	"os/exec"
	"time"
	"bytes"
	"strings"
	"io/ioutil"
	"math"
	"strconv"
)

var (
	powerIconArr = [5]string{" ", " ", " ", " ", " "}
	powerIconCount = len(powerIconArr)
	powerIconFi = 0
)

func getPower() string {
	var powerIcon string
	path := "/sys/class/power_supply/BAT0/"
	status, err := ioutil.ReadFile(path + "status")
	if err != nil {
		panic(err)
	}

	capacity, err := ioutil.ReadFile(path + "capacity")
	if err != nil {
		panic(err)
	}
	capacityStr := strings.TrimSpace(string(capacity))

	if strings.TrimSpace(string(status)) == "Charging" {
		if powerIconFi == powerIconCount-1 {
			powerIconFi = 0
		} else {
			powerIconFi++
		}
		powerIcon = powerIconArr[powerIconFi]
	} else {
		capacityInt, err := strconv.Atoi(capacityStr)
		if err != nil {
			panic(err)
		}
		goConversion_v_ := math.Ceil(float64(capacityInt) * (float64(powerIconCount)/100))
		powerIcon = powerIconArr[int(goConversion_v_)-1]
	}
	return powerIcon + " " + capacityStr + "%"
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
			" " + time.Now().Local().Format("1-2 15:04"),
			getPower(),
			getDwmVersion(),
		}

		s := " " + strings.Join(status, " ")

		exec.Command("xsetroot", "-name", s).Run()
		var now = time.Now()
		time.Sleep(now.Truncate(time.Second).Add(time.Second).Sub(now))
	}
}
