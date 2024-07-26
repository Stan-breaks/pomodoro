package utils

import (
	"os"
	"strings"
)

func SeeLogs() {
	logFile := "/home/stan/projects/goProjects/pomodoro/test.log"
	data, err := os.ReadFile(logFile)
	if err != nil {
		os.Stdout.Write([]byte("\033[H\033[2J\r\n"))
		os.Stdout.Write([]byte("Error with reading data\r\n"))
		return
	}
	stringData := string(data)
	sliceData := strings.Split(stringData, "\n")
	if len(sliceData) > 10 {
		sliceData = sliceData[len(sliceData)-10:]
	}
	for _, item := range sliceData {
		os.Stdout.Write([]byte(item + "\r\n"))
	}
}
