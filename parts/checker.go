package parts

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func CheckerMain() {
	currentTime := time.Now()
	formattedTime := currentTime.Format(" [15:04:05]")
	{
		fmt.Println()
		TimeHMS()
	}
	fmt.Println(" Checker is running...")
	var linuxText = formattedTime + "       Linux : "
	var hyprText = formattedTime + "    Hyprland : "
	{
		if checkOS() {
			print(linuxText)
			PrintInColor("true", colorGreen)
		} else {
			print(linuxText)
			PrintInColor("false", colorRed)
		}
		if checkDesktopEnv() {
			print(hyprText)
			PrintInColor("true", colorGreen)
		} else {
			print(hyprText)
			PrintInColor("false", colorRed)
		}
		pingGitHub(formattedTime)
		pingArchWiki(formattedTime)
	}
	checkProxy()
	checkBasicRefreshRate()
	{
		var stat syscall.Statfs_t
		err := syscall.Statfs(".", &stat)
		if err != nil {
			fmt.Println("ERRORED when scan file-system:", err)
			os.Exit(1)
		}
		free := stat.Bfree * uint64(stat.Bsize)
		if free >= 100*1024*1024*1024 {
			Pass()
			println("DISK_SPACE_PASS")
		} else {

			PrintInColor("DISK_SPACE MAY BE INSUFFICIENT", colorRed)
		}
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		totalMemory := m.Sys / (1024 * 1024 * 1024)
		if totalMemory >= 0 {
			Pass()
			println("RAM_SPACE_PASS")
		} else {
			UnPass()
			PrintInColor("RAM_SIZE MAY BE INSUFFICIENT", colorRed)
		}
	}
}

func pingGitHub(formattedTime string) {
	var NetWorkText = formattedTime + "  Network_GH : "
	print(NetWorkText)
	out, err := exec.Command("ping", "-c 1", "github.com").Output()
	if err != nil {
		PrintInColor("false", colorRed)
		return
	}

	if strings.Contains(string(out), "1 received") {
		PrintInColor("true", colorGreen)
	} else {
		PrintInColor("false", colorRed)
	}
}

func pingArchWiki(formattedTime string) {
	var NetWorkText = formattedTime + "  Network_AW : "
	print(NetWorkText)
	out, err := exec.Command("ping", "-c 1", "archlinux.org").Output()
	if err != nil {
		PrintInColor("false", colorRed)
		return
	}

	if strings.Contains(string(out), "1 received") {
		PrintInColor("true", colorGreen)
	} else {
		PrintInColor("false", colorRed)
	}
}

func checkOS() bool {
	out, err := exec.Command("uname", "-a").Output()
	if err != nil {
		Warning()
		fmt.Println("Error getting OS information:", err)
		return false
	}

	return strings.Contains(strings.ToLower(string(out)), "linux")
}

func checkDesktopEnv() bool {
	out, err := exec.Command("pacman", "-Q").Output()
	if err != nil {
		Warning()
		fmt.Println("Error getting pacman information:", err)
		return false
	}
	return strings.Contains(string(out), "hyprland")
}

func checkProxy() {
	cmd := exec.Command("env")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("ERROR GETTING ENV:", err)
		return
	}

	output := string(out)
	TimeHMS()
	print("  http_proxy : ")
	if strings.Contains(output, "HTTP_PROXY") || strings.Contains(output, "http_proxy") {
		PrintInColor("true", colorCyan)
	} else {
		PrintInColor("false", colorGreen)
	}
	TimeHMS()
	print(" https_proxy : ")
	if strings.Contains(output, "HTTPS_PROXY") || strings.Contains(output, "https_proxy") {
		PrintInColor("true", colorCyan)
	} else {
		PrintInColor("false", colorGreen)
	}

}

func checkBasicRefreshRate() {
	cmd := exec.Command("xrandr")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("执行 xrandr 命令时出错:", err)
		return
	}
	output := strings.Split(string(out), "\n")
	if len(output) >= 3 {
		thirdLine := output[2]
		rateStart := strings.LastIndex(thirdLine, " ") + 1
		rateStr := thirdLine[rateStart : len(thirdLine)-1]
		rateStr = strings.ReplaceAll(rateStr, "*", "")
		rateStr = strings.ReplaceAll(rateStr, "+", "")
		rateStr = strings.ReplaceAll(rateStr, "-", "")
		rateStr = strings.ReplaceAll(rateStr, "/", "")
		rate, err := strconv.ParseFloat(rateStr, 64)
		if err != nil {
			Failed()
			fmt.Println("CANNOT TURN TO FLAOT")
			rate, err = strconv.ParseFloat(rateStr, 64)
			if err != nil {
				Failed()
				fmt.Println("STILL CANNOT TURN TO FLAOT", err)
				return
			}
		}
		if rate >= 59.94 {
			Pass()
		} else {
			UnPass()
		}
		println("REFRESH_RATE")
	} else {
		fmt.Println("INFO MISSING")
	}
}
