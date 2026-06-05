// Package dim provides a function to dim the screen on macOS by simulating key presses.
package dim

import (
	"fmt"
	"os/exec"
)

func Dim() error {
	// We pass each '-e' flag and its script snippet as separate arguments to osascript
	cmd := exec.Command("osascript",
		"-e", "tell application \"System Events\"",
		"-e", "repeat 16 times", "-e", "key code 145", "-e", "end repeat",
		"-e", "repeat 3 times", "-e", "key code 144", "-e", "end repeat",
		"-e", "end tell",
	)

	// Run executes the command and waits for it to complete
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to dim screen: %w", err)
	}
	return nil
}
