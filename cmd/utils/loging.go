package utils

import (
	"fmt"
	"os"
	"pomo/cmd/models"
)

func Log(log models.Log) {
	file, err := os.OpenFile(log.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	_, err = file.WriteString(log.Date + " " + log.Message + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
