#!/usr/bin/bash

mkdir "~/.config/pomodoro"
touch "~/.config/pomodoro/logs.log"
touch "~/.config/pomodoro/custom.json"
go build -o exec main.go
sudo ln -s exec /usr/local/bin/pomo
