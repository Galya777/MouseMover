package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"os/exec"
)

var imageButton, imageButton2 *ImageButton

func moveMouse() (*exec.Cmd, error) {
	// Create the command
	cmd := exec.Command("keep-presence", "--seconds", "30")

	// Start the command asynchronously
	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start keep-presence: %v", err)
	}

	fmt.Println("keep-presence started.")
	return cmd, nil
}
func killMoveMouse(cmd *exec.Cmd) error {
	// Check if the command is not nil and the process is running
	if cmd != nil && cmd.Process != nil {
		// Kill the process
		err := cmd.Process.Kill()
		if err != nil {
			return fmt.Errorf("failed to stop keep-presence: %v", err)
		}
		fmt.Println("keep-presence stopped.")
	} else {
		return fmt.Errorf("no running process to stop")
	}

	return nil
}

func main() {

	var cmd *exec.Cmd
	var err error
	clicked := false
	// Create the application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("Keep Linux Awake!")

	// Create a custom button with an image
	imageButton = NewCustomButton("wakeUP.png", "pictures/wakeUP.png", func() {
		clicked = true
		err = killMoveMouse(cmd)
		if err != nil {
			fmt.Println("Error stopping command:", err)
		}
		myWindow.SetContent(container.NewCenter(imageButton2))
	})
	imageButton2 = NewCustomButton("sleeping.png", "pictures/sleeping.png", func() {
		clicked = true
		myWindow.SetContent(container.NewCenter(imageButton))
		cmd, err = moveMouse()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Command executed successfully.")
		}
	})

	// Set the window content and show it
	if !clicked {
		myWindow.SetContent(container.NewCenter(imageButton2))
	} else {
		myWindow.SetContent(container.NewCenter(imageButton))
	}
	myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.ShowAndRun()

	err = killMoveMouse(cmd)
}
