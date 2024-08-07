#!/usr/bin/bash

mkdir -p "$HOME/.config/pomodoro"
touch "$HOME/.config/pomodoro/logs.log"
cat <<EOF >"$HOME/.config/pomodoro/custom.json"
[]
EOF
go build -o pomo main.go
sudo ln -sf "$(pwd)/pomo" /usr/local/bin/pomo
