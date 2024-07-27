package utils

import (
	"encoding/json"
	"github.com/Stan-breaks/pomo/cmd/models"
	"os"
	"time"
)

func LoadSchedule() {
	customizeFile := "/home/stan/projects/goProjects/pomodoro/customise.json"
	data, err := os.ReadFile(customizeFile)
	if err != nil {
		os.Stdout.Write([]byte("\033[H\033[2J\r\n"))
		os.Stdout.Write([]byte("Error occurred with custom schedule \r\n"))
		return
	}
	var items []models.ScheduleItem
	err = json.Unmarshal(data, &items)
	if err != nil {
		os.Stdout.Write([]byte("\033[H\033[2J\r\n"))
		os.Stdout.Write([]byte("Error occurred with custom schedule \r\n"))
		return
	}
	for _, item := range items {
		currentTime := time.Now()
		countDown := models.CountDown{
			Message: item.Activity,
			Minutes: item.Minutes,
			Seconds: item.Seconds,
			Date:    currentTime.Format("2006-01-02 15:04:05"),
		}
		Count(countDown)
	}
}
