package utils

import (
	"encoding/json"
	"os"
	"pomo/cmd/models"
	"pomo/cmd/routes"
)

func Edit() {
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
	editPointer := 0
	for {
		EditMenu(items, editPointer)
		b := make([]byte, 1)
		_, err = os.Stdin.Read(b)
		if err != nil {
			os.Stdout.Write([]byte("Error reading input \r\n"))
			return
		}
		switch b[0] {
		case 'q', 'Q':
			return
		case 65:
			if editPointer > 0 {
				editPointer--
			}
		case 66:
			if editPointer < len(items)-1 {
				editPointer++
			}
		case 13, 10:
			if editPointer == 5 {
				return
			}
			os.Stdout.Write([]byte("\033[H\033[2J\r\n"))
			routes.HandleEdit(editPointer)
		}
	}

}
