package utils

import (
	"os"
	"strings"
)

func SeeLogs() {
	logFile := "/home/stan/projects/goProject/pomodoro"
	data, err := os.ReadFile(logFile)
	if err != nil {
		os.Stdout.Write([]byte("Error with reading data\r\n"))
		return
	}
	os.Stdout.Write([]byte("\033[H\033[2J\r\n"))
	stringData := string(data)
	sliceData := strings.Split(stringData, "\n")[:-10]
	for i, item := 10; i > 0; i-- {
		os.St
	}
}
