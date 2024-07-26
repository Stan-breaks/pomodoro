package routes

import (
	"encoding/json"
	"os"
	"pomo/cmd/models"
	"strconv"
)

func EditItem(pointer int, items []models.ScheduleItem, file string) {
	name := ""
	minutes := ""
	seconds := ""
	for i := range 3 {
		os.Stdout.Write([]byte("\033[H\033[2J\r\n"))
		switch i {
		case 0:
			os.Stdout.Write([]byte("  ===Please Enter Name of task=== \r\n"))
		case 1:
			os.Stdout.Write([]byte("  ===Please Enter Minutes of task===\r\n"))
		case 2:
			os.Stdout.Write([]byte("  ===Please Enter Seconds of task===\r\n"))
		}
		for {

			b := make([]byte, 1)
			_, err := os.Stdin.Read(b)
			if err != nil {
				os.Stdout.Write([]byte("Error reading input \r\n"))
				return
			}
			switch b[0] {
			case 13, 10:
				goto nextField
			case 27:
				return
			case 127, 8:
				switch i {
				case 0:
					if len(name) > 0 {
						name = name[:len(name)-1]
						os.Stdout.Write([]byte("\b \b"))
					}
				case 1:
					if len(minutes) > 0 {
						minutes = minutes[:len(minutes)-1]
						os.Stdout.Write([]byte("\b \b"))
					}
				case 2:
					if len(seconds) > 0 {
						seconds = seconds[:len(seconds)-1]
						os.Stdout.Write([]byte("\b \b"))
					}
				}
			default:
				switch i {
				case 0:
					name += string(b)
					os.Stdout.Write(b)
				case 1:
					minutes += string(b)
					os.Stdout.Write(b)
				case 2:
					seconds += string(b)
					os.Stdout.Write(b)
				}

			}

		}
	nextField:
	}
	intMinutes, err := strconv.Atoi(minutes)
	if err != nil {
		os.Stdout.Write([]byte("Error with string convertion \r\n"))
		return
	}
	intSeconds, err := strconv.Atoi(seconds)
	if err != nil {
		os.Stdout.Write([]byte("Error with string convertion \r\n"))
		return
	}
	newItem := models.ScheduleItem{
		Activity: name,
		Minutes:  intMinutes,
		Seconds:  intSeconds,
	}
	newItems := []models.ScheduleItem{}
	for i, item := range items {
		if i == pointer {
			newItems = append(newItems, newItem)
		} else {
			newItems = append(newItems, item)
		}
	}
	byteData, err := json.Marshal(newItems)
	if err != nil {
		os.Stdout.Write([]byte("Error with json to byte convertion"))
		return
	}
	err = os.WriteFile(file, byteData, 0644)
	if err != nil {
		os.Stdout.Write([]byte("Error with writing to file"))
		return
	}

}
