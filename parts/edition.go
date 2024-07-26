package parts

import (
	"fmt"
	"time"
)

const (
	/*
		textBlack = iota + 30
		textRed
		textGreen
		textYellow
		textBlue
		textPurple
		textCyan
		textWhite
	*/
	colorCyan   = 36
	colorGreen  = 32
	colorYellow = 33
	colorRed    = 31
)

func Edition() {
	println()
	var Version = "  「 novice @ Siracusant | Dev.0.01 | LICENSE GPL 3.0 」  "
	var Source = " ##  Source Branch  >>  " //github.com/Siracusant/novice
	{
		currentTime := time.Now()
		formattedTime := currentTime.Format("15:04:05 01-02/2006")
		PrintInColor(Version, colorCyan)
		print("\n")
		fmt.Print(Source)
		PrintItalic("github.com/Siracusant/novice")
		print(" ##  Start up at    >>  ")
		PrintItalic(formattedTime)
	}
}

func PrintInColor(str string, colorCode int) {
	fmt.Printf("\x1b[1;%dm%s\x1b[0m\n", colorCode, str)
}

func PrintItalic(str string) {
	fmt.Printf("\x1b[3m%s\x1b[0m\n", str)
}

func Warning() {
	printInColorWithoutReturn(" [WARNING]: ", colorYellow)
}

func Done() {
	printInColorWithoutReturn(" [DONE]  :: ", colorGreen)
}

func Pass() {
	printInColorWithoutReturn(" [PASS]  :: ", colorGreen)
}

func UnPass() {
	printInColorWithoutReturn(" [UNPASS]:: ", colorRed)
}

func Failed() {
	printInColorWithoutReturn(" [FAILED]:: ", colorRed)
}

func Remind() {
	printInColorWithoutReturn(" [REMIND]:: ", colorYellow)
}

func printInColorWithoutReturn(str string, colorCode int) {
	fmt.Printf("\x1b[1;%dm%s\x1b[0m", colorCode, str)
}

func TimeHMS() {
	currentTime := time.Now()
	TimeHMS := currentTime.Format(" [15:04:05]")
	print(TimeHMS)
}
