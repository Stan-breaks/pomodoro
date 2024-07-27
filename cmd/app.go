package cmd

import (
	"fmt"
	"github.com/Stan-breaks/pomo/cmd/utils"
	"golang.org/x/term"
	"os"
)

func App() {
	oldstate, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldstate)
	homePointer := 0
	for {
		utils.DisplayMenu(homePointer)
		b := make([]byte, 1)
		_, err = os.Stdin.Read(b)
		if err != nil {
			os.Stdout.Write([]byte("Error reading input \r\n"))
			return
		}
		switch b[0] {
		case 'w', 'W':
			if homePointer > 0 {
				homePointer--
			}
		case 's', 'S':
			if homePointer < 5 {
				homePointer++
			}
		case 'q', 'Q':
			return
		case 65:
			if homePointer > 0 {
				homePointer--
			}
		case 66:
			if homePointer < 5 {
				homePointer++
			}
		case 13, 10:
			if homePointer == 5 {
				return
			}

			os.Stdout.Write([]byte("\033[H\033[2J\r\n"))
			// clear the screen
			HandleRoutes(homePointer)
		}
	}
}
