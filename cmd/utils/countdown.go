package utils

import (
	"fmt"
	"os"
	"pomo/cmd/models"
	"strconv"
	"strings"
	"time"
)

var number = Number()

func Count(countDown models.CountDown) {
	os.Stdout.Write([]byte("\033[H\033[2J\r\n"))
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	totalSeconds := (countDown.Minutes * 60) + countDown.Seconds
	logFile := "/home/stan/projects/goProjects/pomodoro/test.log"
	startLog := models.Log{
		Filepath: logFile,
		Message:  countDown.Message + " started",
		Date:     countDown.Date,
	}
	Log(startLog)
	for range ticker.C {
		os.Stdout.Write([]byte(countDown.Message + " started at " + countDown.Date + "\r\n"))
		if totalSeconds < 0 {
			break
		}
		minutes := totalSeconds / 60
		seconds := totalSeconds % 60
		stringMinutes := strings.Split(fmt.Sprintf("%02d", minutes), "")
		stringSeconds := strings.Split(fmt.Sprintf("%02d", seconds), "")
		intMinutes := make([]int, len(stringMinutes))
		intSeconds := make([]int, len(intMinutes))
		for i, s := range stringMinutes {
			intMinutes[i], _ = strconv.Atoi(s) // converting the strings to int
		}
		for i, s := range stringSeconds {
			intSeconds[i], _ = strconv.Atoi(s)
		}
		// Split each number into lines
		var lines [][]string
		for i, digit := range append(intMinutes, intSeconds...) {
			if i == 2 {
				lines = append(lines, strings.Split(Colon, "\n"))
			}
			lines = append(lines, strings.Split(number[digit], "\n"))
		}
		// Concatenate corresponding lines
		for j := 0; j < len(lines[0]); j++ {
			for _, line := range lines {
				os.Stdout.Write([]byte(line[j] + " ")) // adding some space btwn the no
			}
			os.Stdout.Write([]byte("\r\n"))
		}
		time.Sleep(1 * time.Second)
		os.Stdout.Write([]byte("\033[H\033[2J\r\n")) // clear the screen
		totalSeconds--
	}
	os.Stdout.Write([]byte("\r\n\r\n"))
	endLog := models.Log{
		Filepath: logFile,
		Message:  countDown.Message + " ended",
		Date:     time.Now().Format("2006-01-02 15:04:05"),
	}
	Log(endLog)
}
