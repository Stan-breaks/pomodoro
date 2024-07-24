package routes

import (
	"encoding/json"
	"os"
	"pomo/cmd/models"
)

func DeleteItem(pointer int, file string, items []models.ScheduleItem) {
	newItems := []models.ScheduleItem{}
	for i, item := range items {
		if i != pointer {
			newItems = append(newItems, item)
		}
	}

	byteData, err := json.Marshal(newItems)
	if err != nil {
		os.Stdout.Write([]byte("error with json to byte convertion\r\n"))
		return
	}
	err = os.WriteFile(file, byteData, 0644)
	if err != nil {
		os.Stdout.Write([]byte("Error with writing to file\r\n"))
		return
	}
}
