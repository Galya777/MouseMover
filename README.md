# Creating a Mouse Mover GUI Application in Linux Using Go and Fyne

This document outlines the steps to create a simple mouse mover GUI application in Linux using the Go programming language and the Fyne framework. The application will feature a graphical interface that allows users to start and stop a mouse mover function while displaying an icon.
Prerequisites

    Go: Ensure you have Go installed on your Linux system. You can download it from golang.org.

    Fyne Framework: Install the Fyne framework by running the following command:

    bash

    go get fyne.io/fyne/v2

# Step 1: Create the Fyne Application

    Create a new directory for your project:

    bash

    mkdir MouseMoverApp
    cd MouseMoverApp

    Create a Go file named main.go in this directory. This file will contain the main application logic.

    Write the application code:

go

package main

import (
    "log"
    "os/exec"
    "time"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/image"
    "fyne.io/fyne/v2/widget"
)

var moverProcess *exec.Cmd

// Function to start the mouse mover command
func startMouseMover() {
    moverProcess = exec.Command("keep-presence", "--seconds", "30")
    if err := moverProcess.Start(); err != nil {
        log.Fatalf("Failed to start mouse mover: %v", err)
    }
    log.Println("Mouse mover started")
}

// Function to stop the mouse mover command
func stopMouseMover() {
    if moverProcess != nil {
        if err := moverProcess.Process.Kill(); err != nil {
            log.Fatalf("Failed to stop mouse mover: %v", err)
        }
        log.Println("Mouse mover stopped")
    }
}

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Mouse Mover")

    // Load the application icon
    iconPath := "icon.png" // Update to your icon's path
    resource, err := fyne.LoadResourceFromPath(iconPath)
    if err != nil {
        log.Fatalf("Failed to load icon: %v", err)
    }
    myApp.SetIcon(resource)

    // Create start and stop buttons
    startButton := widget.NewButton("Start Mover", func() {
        startMouseMover()
    })
    stopButton := widget.NewButton("Stop Mover", func() {
        stopMouseMover()
    })

    // Set window content
    myWindow.SetContent(container.NewVBox(
        widget.NewLabel("Mouse Mover Application"),
        startButton,
        stopButton,
        widget.NewButton("Quit", func() {
            myApp.Quit()
        }),
    ))

    myWindow.ShowAndRun()
}

Code Explanation

    Start and Stop Functions: The startMouseMover function runs the keep-presence command, which keeps the mouse active for a specified duration, while stopMouseMover kills the process when requested.
    UI Components: The application contains buttons to start and stop the mouse mover, along with a quit button.
    Icon Setup: The application icon is loaded and set in the application window.

Step 2: Create a Debian Package

To package your application as a Debian executable file (DEB), you can use a shell script.
Example Build Script

Create a shell script named build_deb.sh in the same directory:

bash

#!/bin/bash

# Variables
APP_NAME="MouseMoverApp"                # Name of your application
VERSION="1.0.0"                         # Version of your application
ARCHITECTURE="amd64"                    # Architecture
BUILD_DIR="build"                       # Build directory
DEB_DIR="$BUILD_DIR/DEBIAN"             # Directory for DEBIAN control files
BIN_DIR="$BUILD_DIR/usr/local/bin"      # Directory for the binary
ICON_PATH="icon.png"                    # Path to your icon file

# Create necessary directories
mkdir -p "$BIN_DIR"
mkdir -p "$DEB_DIR"

# Build the Fyne application for Linux
echo "Building the Fyne application..."
GOOS=linux GOARCH=amd64 go build -o "$APP_NAME" .

# Check if the binary exists
if [ ! -f "$APP_NAME" ]; then
    echo "Error: Application binary '$APP_NAME' not found."
    exit 1
fi

# Move the binary to the bin directory
echo "Copying binary to $BIN_DIR..."
cp "$APP_NAME" "$BIN_DIR/"

# Create control file
cat <<EOF > "$DEB_DIR/control"
Package: $APP_NAME
Version: $VERSION
Section: base
Priority: optional
Architecture: $ARCHITECTURE
Maintainer: Your Name <your.email@example.com>
Description: A simple mouse mover application
 A simple application that moves the mouse to prevent sleep.
EOF

# Copy the icon to the appropriate directory
mkdir -p "$BUILD_DIR/usr/share/icons/hicolor/48x48/apps/"
cp "$ICON_PATH" "$BUILD_DIR/usr/share/icons/hicolor/48x48/apps/$APP_NAME.png"

# Create the .desktop file
cat <<EOF > "$BUILD_DIR/usr/share/applications/$APP_NAME.desktop"
[Desktop Entry]
Version=1.0
Name=$APP_NAME
Comment=A simple mouse mover application
Exec=/usr/local/bin/$APP_NAME
Icon=$APP_NAME
Terminal=false
Type=Application
Categories=Utility;
EOF

# Build the DEB package
echo "Building the DEB package..."
dpkg-deb --build "$BUILD_DIR"

# Move the DEB package to the current directory
mv "$BUILD_DIR.deb" "./$APP_NAME-$VERSION.deb"

# Clean up
rm -rf "$BUILD_DIR"

echo "DEB package created: ./$APP_NAME-$VERSION.deb"

Usage of the Build Script

    Make the Script Executable:

    bash

chmod +x build_deb.sh

Run the Script:

bash

    ./build_deb.sh

This script will build your application and create a .deb package that includes your application binary and any required icons and desktop entries.
Step 3: Install the DEB Package

Once the DEB package is created, you can install it on any Debian-based system using:

bash

sudo dpkg -i MouseMoverApp-1.0.0.deb

Step 4: Run Your Application

After installing the DEB package, you should be able to find your application in your application menu, complete with its icon. You can launch it from there or run it directly from the terminal.


Conclusion

You have successfully created a simple mouse mover GUI application in Linux using Go and Fyne. You also learned how to package it as a DEB file for easy distribution and installation. This documentation should serve as a reference for creating similar applications in the future. 


![Снимка на екрана на 2024-10-06 14-35-14](https://github.com/user-attachments/assets/ab145a99-3259-44f2-bb03-435d24371c8d)

![изображение](https://github.com/user-attachments/assets/b79c8f88-c4ba-4529-a1ae-7eb772f01532)



