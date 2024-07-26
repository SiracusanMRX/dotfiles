package main

import (
	"Novice/parts"
	"Novice/parts/pkg"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	parts.Edition()
	parts.CheckerMain()
	{
		parts.Done()
		println("CHECKER_EXITED")
	}
	{
		currentUser, err := user.Current()
		if err != nil {
			parts.Warning()
			fmt.Println("ERRORED when getting USR_NAME:", err)
		}
		{
			parts.Done()
			fmt.Print("USR_NAME | ")
			parts.PrintItalic(currentUser.Username)
			whetherOrNot()
			pkg.Comparison()
		}
	}
}

func whetherOrNot() {
	parts.Warning()
	fmt.Println("If you think the current device can proceed with the NOVICE installation, enter \"y\", or if you think there is an actual anomaly, enter \"n\".")
	for {
		parts.Warning()
		fmt.Print("ARE U SURE? [y/n] = ")
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			parts.Warning()
			fmt.Println("INPUT WRONG", err)
			continue
		}
		input = strings.ToLower(input)
		if input == "y" {
			parts.TimeHMS()
			fmt.Println(" Continuing...")
			break
		} else if input == "n" {
			parts.TimeHMS()
			fmt.Println(" Process Ending...")
			os.Exit(0)
		} else {
			parts.Failed()
			fmt.Println("INVALID INPUT, re-enter.")
		}
	}
}
