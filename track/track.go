// Package track provides functionality to track the active window and its application on macOS.
package track

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// TrackWindow Returns the title of the frontmost window and its application name.
func TrackWindow() (app string, window string, err error) {
	// AppleScript to get the frontmost application and its active window title
	script := `
		tell application "System Events"
			set frontApp to name of first application process whose frontmost is true
			tell process frontApp
				if (count of windows) > 0 then
					set windowTitle to name of window 1
				else
					set windowTitle to "No Active Window"
				end if
			end tell
		end tell
		return frontApp & "||" & windowTitle
	`

	// Execute the AppleScript via osascript
	cmd := exec.Command("osascript", "-e", script)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return "", "", fmt.Errorf("failed to run apple script: %v (stderr: %s)", err, stderr.String())
	}

	// Clean up and parse the output
	result := strings.TrimSpace(out.String())
	parts := strings.Split(result, "||")

	if len(parts) == 2 {
		return parts[0], parts[1], nil
	}

	return "", "", fmt.Errorf("unexpected output format: %s", result)
}
