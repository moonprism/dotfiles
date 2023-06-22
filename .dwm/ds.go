// dwm status
// usage: sudo go build -o /usr/local/bin/ds ./ds.go
// ♥ from moonprism

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	powerIconArr = [5]string{" ", " ", " ", " ", " "}
	powerIconFi  = 0 // frame index
)

func getPower() string {
	icon := ""
	path := "/sys/class/power_supply/BAT0/"

	statusTex, err := os.ReadFile(path + "status")
	if err != nil {
		panic(err)
	}

	capacityTex, err := os.ReadFile(path + "capacity")
	if err != nil {
		panic(err)
	}
	capacity := strings.TrimSpace(string(capacityTex))

	if strings.TrimSpace(string(statusTex)) == "Charging" {
		// 充电中...动画
		if powerIconFi == len(powerIconArr)-1 {
			powerIconFi = 0
		} else {
			powerIconFi++
		}
		icon = powerIconArr[powerIconFi]
	} else {
		capacityF64, err := strconv.ParseFloat(capacity, 64)
		if err != nil {
			panic(err)
		}
		goConversion_v_ := math.Ceil(capacityF64 * (float64(len(powerIconArr)) / 100))
		icon = powerIconArr[int(goConversion_v_)-1]
	}
	return icon // + " " + capacity + "%"
}

var dwmVersion string

func getDwmVersion() string {
	if dwmVersion != "" {
		return dwmVersion
	}
	cmd := exec.Command("dwm", "-v")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Run()
	dwmVersion = stderr.String()
	return dwmVersion
}

var cores = runtime.NumCPU()

func getLoadAVG() string {
	var load1, load5, load15 float32
	loadavg, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		panic(err)
	}
	if _, err = fmt.Sscanf(string(loadavg), "%f %f %f", &load1, &load5, &load15); err != nil {
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

	total, available := 0.0, 0.0
	for info := bufio.NewScanner(file); info.Scan(); {
		prop, val := "", 0.0
		if _, err = fmt.Sscanf(info.Text(), "%s %f", &prop, &val); err != nil {
			panic(err)
		}
		if prop == "MemTotal:" {
			total = val
		}
		if prop == "MemAvailable:" {
			available = val
		}
		if total != 0.0 && available != 0.0 {
			break
		}
	}
	p := (total - available) * 100.0 / total
	return " " + fmt.Sprintf("%.f%%", p)
}

var timeIconArr = [12]string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}

func getTime() string {
	now := time.Now()
	hour := now.Hour()
	return timeIconArr[hour%12] + fmt.Sprintf("%d:%02d", hour, now.Minute())
}

func main() {
	// 如果想显示秒数，请参考: https://github.com/schachmat/gods/blob/75cdf14a346f5ce5d8ca34953456f7f0f98067a8/gods.go#L299
	t := time.Tick(time.Second)
	for range t {
		status := []string{
			"",
			getLoadAVG(),
			getMemInfo(),
			getTime(),
			getPower(),
			getDwmVersion(),
		}
		s := strings.Join(status, " ")
		exec.Command("xsetroot", "-name", s).Run()
	}
}
