package utils

import (
	"fmt"
	"pomo/models"
	"time"
)

func Count(countDown models.CountDown) {
	fmt.Println(countDown.Message + " at " + countDown.Date)
	for i := (countDown.Minutes * 60) + countDown.Seconds; i >= 0; i-- {
		minutes := i / 60
		seconds := i % 60
		fmt.Printf("\r %02d:%02d", minutes, seconds)
		time.Sleep(1 * time.Second)
	}
}
