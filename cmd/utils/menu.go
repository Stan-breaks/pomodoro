package utils

import (
	"os"
)

func DisplayMenu(opt int) {
	menu := []string{
		"Start a new timer",
		"Load scheduler.",
		"Edit scheduler.",
		"See your logs.",
		"Analysis",
		"Exit",
	}
	os.Stdout.Write([]byte("\033[2J\033[H\r\n"))
	os.Stdout.Write([]byte(" ===Welcome to pomo - Your pomodoro scheduler===\r\n\r\n"))
	for i, item := range menu {
		if i == opt {
			os.Stdout.Write([]byte("  >" + item + "\r\n"))
		} else {
			os.Stdout.Write([]byte("   " + item + "\r\n"))
		}
	}
}
