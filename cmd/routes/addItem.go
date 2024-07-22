package routes

import (
	"os"
	"pomo/cmd/models"
)

func AddItem(pointer int, file string, items []models.ScheduleItem) {
	name := " "
	minutes := " "
	seconds := " "
	for i := range 3 {
		os.Stdout.Write([]byte("\033[H\033[2J\r\n"))
		for {
			b := make([]byte, 1)
			_, err := os.Stdin.Read(b)
			if err != nil {
				os.Stdout.Write([]byte("Error reading input \r\n"))
				return
			}
			switch b[0] {
			case 13, 10:
				break
			default:
				switch i {
				case 0:
					os.Stdout.Write([]byte("  ===Please Enter Name of task=== \r\n"))
					name += string(b)
					os.Stdout.Write([]byte(name))
				case 1:
					os.Stdout.Write([]byte("  ===Please Enter Minutes of task===\r\n"))
					minutes += string(b)
					os.Stdout.Write([]byte(minutes))
				case 2:
					os.Stdout.Write([]byte("  ===Please Enter Seconds of task===\r\n"))
					seconds += string(b)
					os.Stdout.Write([]byte(seconds))
				}
			}
		}
	}
}

func AddToList(file string) {
}
