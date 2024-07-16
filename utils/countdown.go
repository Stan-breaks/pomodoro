package utils

import (
	"fmt"
	"pomo/models"
	"strconv"
	"strings"
	"time"
)

var number = Number()

func Count(countDown models.CountDown) {
	fmt.Print("\033[H\033[2J") // clear the screen
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	totalSeconds := (countDown.Minutes * 60) + countDown.Seconds
	for range ticker.C {
	fmt.Println(countDown.Message + " at " + countDown.Date)
		if totalSeconds < 0 {
			break
		}
		minutes := totalSeconds / 60
		seconds := totalSeconds % 60
		stringMinutes := strings.Split(fmt.Sprintf("%02d", minutes), "")
		stringSeconds := strings.Split(fmt.Sprintf("%02d", seconds), "")
		intMinutes := make([]int, len(stringMinutes))
		intSeconds := make([]int, len(intMinutes))
		for i, s := range stringMinutes {
			intMinutes[i], _ = strconv.Atoi(s)//converting the strings to int
		}
		for i, s := range stringSeconds {
			intSeconds[i], _ = strconv.Atoi(s)
		}
		// Split each number into lines
		var lines [][]string
		for i, digit := range append(intMinutes, intSeconds...) {
			if i == 2 {
				lines = append(lines, strings.Split(Colon, "\n"))
			}
			lines = append(lines, strings.Split(number[digit], "\n"))
		}
		// Concatenate corresponding lines
		for j := 0; j < len(lines[0]); j++ {
			for _, line := range lines {
				fmt.Print(line[j] + " ") // adding some space btwn the no
			}
			fmt.Println()
		}
		time.Sleep(1 * time.Second)
		fmt.Print("\033[H\033[2J") // clear the screen
		totalSeconds--
	}
	fmt.Printf("\n\n")
}
