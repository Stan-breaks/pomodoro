package main

import (
	"fmt"
	"pomo/models"
	"pomo/utils"
	"time"
)

func main() {
	now := time.Now()
	year, month, day := now.Date()
	hour, minn, sec := now.Clock()
	date := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minn, sec)
	message := "Dev work started"
	notification := models.Notification{
		Message: message,
	}
	countDown := models.CountDown{
		Date:    date,
		Minutes: 1,
		Seconds: 30,
		Message: message,
	}
	log := models.Log{
		Message:  message,
		Filepath: "/home/stanley/ogProjects/pomodoro/test.log",
		Date:     date,
	}
	utils.Notify(notification)
	utils.Log(log)
	utils.Count(countDown)
}
