package utils

import (
	"fmt"
	"os/exec"
	"pomo/models"
)

func Notify(notification models.Notification) {
	cmd := exec.Command("notify-send", "Pomo", notification.Message, "-i", notification.IconPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
