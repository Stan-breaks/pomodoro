package cmd

import (
	"github.com/Stan-breaks/pomodoro/cmd/models"
	"github.com/Stan-breaks/pomodoro/cmd/utils"
	"time"
)

func HandleRoutes(opt int) {
	switch opt {
	case 0:
		currentTime := time.Now()
		countDown := models.CountDown{
			Message: "Demo",
			Minutes: 25,
			Seconds: 0,
			Date:    currentTime.Format("2006-01-02 15:04:05"),
		}
		utils.Count(countDown)
	case 1:
		utils.LoadSchedule()
	case 2:
		utils.Edit()
	}
}
