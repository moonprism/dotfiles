package main

import (
	"os/exec"
	"time"
	"bytes"
	"strings"
	"io/ioutil"
	"math"
	"strconv"
	"fmt"
	"runtime"
	"os"
	"bufio"
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
	return powerIcon// + " " + capacityStr + "%"
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

var cores = runtime.NumCPU()

func getLoadAVG() string {
	var load1,load5,load15 float32
	var loadavg, err = ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		panic(err)
	}
	_, err = fmt.Sscanf(string(loadavg), "%f %f %f", &load1, &load5, &load15)
	if err != nil {
		panic(err)
	}
	return " " + fmt.Sprintf("%.2f,%.1f,%.1f/%d", load1, load5, load15, cores)
}

func getMemInfo() string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var total, available = -1.0, -1.0
	for info := bufio.NewScanner(file); info.Scan(); {
		var prop, val = "", 0.0
		if _, err = fmt.Sscanf(info.Text(), "%s %f", &prop, &val); err != nil {
			panic(err)
		}
		if prop == "MemTotal:" {
			total = val
		}
		if prop == "MemAvailable:" {
			available = val
		}
		if total != -1.0 && available != -1.0 {
			break
		}
	}
	p := (total - available) * 100.0 / total
	return " " + fmt.Sprintf("%.1f%%", p)
}

var timeIconArr = [12]string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}

func getTime() string {
	var now = time.Now()
	var hour = now.Hour()
	return timeIconArr[hour%12-1] + fmt.Sprintf("%d:%02d", hour, now.Minute())
}

func main() {
	for {
		status := []string{
			getLoadAVG(),
			getMemInfo(),
			getTime(),
			getPower(),
			getDwmVersion(),
		}

		s := " " + strings.Join(status, " ")

		exec.Command("xsetroot", "-name", s).Run()
		var now = time.Now()
		time.Sleep(now.Truncate(time.Second).Add(time.Second).Sub(now))
	}
}
