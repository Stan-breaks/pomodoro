package utils

import (
	"os"
	"pomo/cmd/models"
	"strconv"
)

func EditMenu(items []models.ScheduleItem, pointer int) {
	os.Stdout.Write([]byte("\033[2J\033[H\r\n"))
	os.Stdout.Write([]byte("  === This is your Schedule === \r\n"))
	os.Stdout.Write([]byte("(hint: a to add in current position, d to delete,e to edit,l add to the list ,q to quit )\r\n"))
	for i, item := range items {
		if i == pointer {
			os.Stdout.Write([]byte("  >" + item.Activity + " (" + strconv.Itoa(item.Minutes) + ":" + strconv.Itoa(item.Seconds) + ") \r\n"))
		} else {

			os.Stdout.Write([]byte("   " + item.Activity + " (" + strconv.Itoa(item.Minutes) + ":" + strconv.Itoa(item.Seconds) + ") \r\n"))

		}
	}
}