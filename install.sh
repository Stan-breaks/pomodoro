#!/usr/bin/bash

# Create directory and files
mkdir -p "$HOME/.config/pomodoro"
touch "$HOME/.config/pomodoro/logs.log"
cat <<EOF >"$HOME/.config/pomodoro/custom.json"
[]
EOF

# Set correct permissions
sudo chmod 755 "$HOME/.config/pomodoro"
sudo chmod 644 "$HOME/.config/pomodoro/logs.log"
sudo chmod 644 "$HOME/.config/pomodoro/custom.json"

# Build the Go program
go build -o pomo main.go

# Install the program
sudo install -o root -g root -m 755 pomo /usr/local/bin/pomo

# Clean up
rm pomo
