package main

import (
	"fmt"
	"os"
	"pomo/models"
	"pomo/utils"
	"time"
)

func main() {
	icon := "/home/stanley/Downloads/wallpapers/pomo/pomo.png"
	logFile := "/home/stanley/ogProjects/pomodoro/test.log"
	message := os.Args[1]

	for {
		now := time.Now()
		year, month, day := now.Date()
		hour, minn, sec := now.Clock()
		date := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minn, sec)
		startNotification := models.Notification{
			Message:  message + " time started",
			IconPath: icon,
		}
		workCountDown := models.CountDown{
			Date:    date,
			Minutes: 1,
			Seconds: 00,
			Message: message + " time started",
		}
		startLog := models.Log{
			Message:  message + " time started",
			Filepath: logFile,
			Date:     date,
		}
		utils.Notify(startNotification)
		utils.Log(startLog)
		utils.Count(workCountDown)
		now = time.Now()
		year, month, day = now.Date()
		hour, minn, sec = now.Clock()
		date = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minn, sec)
		endNotification := models.Notification{
			Message:  message + " time ended",
			IconPath: icon,
		}
		endLog := models.Log{
			Message:  message + " time ended",
			Filepath: logFile,
			Date:     date,
		}
		startBreakNotification := models.Notification{
			Message:  "Break time started",
			IconPath: icon,
		}
		breakCountDown := models.CountDown{
			Date:    date,
			Minutes: 1,
			Seconds: 0,
			Message: "Break time",
		}
		endBreakNotification := models.Notification{
			Message:  "Break time ended",
			IconPath: icon,
		}
		utils.Notify(endNotification)
		utils.Log(endLog)
		utils.Notify(startBreakNotification)
		utils.Count(breakCountDown)
		utils.Notify(endBreakNotification)
		fmt.Printf("\n\n")
		var response string
		fmt.Println("Do you want to start another session?(y/n)")
		_, err := fmt.Scan(&response)
		if err != nil {
			fmt.Println(err)
			break
		}
		if response != "y" {
			break
		}
		fmt.Printf("\n\n")
	}
}
