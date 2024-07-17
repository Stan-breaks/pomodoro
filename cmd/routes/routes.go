package routes

import (
	"os"
	"pomo/cmd/models"
	"pomo/cmd/utils"
	"time"
)

func HandleRoutes(opt int) {
	switch opt {
	case 0:
		currentTime := time.Now()
		countDown := models.CountDown{
			Message: "Demo",
			Minutes: 1,
			Seconds: 0,
			Date:    currentTime.Format("2006-01-02 15:04:05"),
		}
		utils.Count(countDown)
	default:

		os.Stdout.Write([]byte("\033[H\033[2J\r\n")) // clear the screen
  os.Stdout.Write([]byte("Under Construstion...\r\n"))
	}
}
