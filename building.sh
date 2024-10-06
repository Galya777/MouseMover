#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o MouseMoverMine

# Variables
APP_NAME="MouseMover"                   # Name of your application
VERSION="1.0.0"                         # Version of your application
ARCHITECTURE="amd64"                    # Architecture
APP_DIR="./$APP_NAME"                   # Application directory
BUILD_DIR="build"                       # Build directory
DEB_DIR="$BUILD_DIR/DEBIAN"             # Directory for DEBIAN control files
BIN_DIR="$BUILD_DIR//usr/local/bin"      # Directory for the binary                 # Path to your icon file

# Create necessary directories
mkdir -p "$BIN_DIR"
mkdir -p "$DEB_DIR"

# Build the Fyne application for Linux
echo "Building the Fyne application..."
GOOS=linux GOARCH=amd64 go build -o "$APP_DIR/$APP_NAME" .

# Check if the binary exists
if [ ! -f "$APP_DIR/$APP_NAME" ]; then
    echo "Error: Application binary '$APP_DIR/$APP_NAME' not found."
    exit 1
fi

# Move the binary to the bin directory
echo "Copying binary to $BIN_DIR..."
cp "sleeping.png" "$BIN_DIR/"
cp "wakeUP.png"  "$BIN_DIR/"
cp "icon.png" "$BIN_DIR/"
cp "$APP_DIR/$APP_NAME" "$BIN_DIR/"

# Copy any other images used by the application
mkdir -p "$BUILD_DIR/usr/share/$APP_NAME/"
cp "sleeping.png" "$BUILD_DIR/usr/share/$APP_NAME/" # Add your image path here
cp "wakeUP.png" "$BUILD_DIR/usr/share/$APP_NAME/" # Add your image path here
cp "icon.png" "$BUILD_DIR/usr/share/$APP_NAME/" # Add your image path here

# Create control file
cat <<EOF > "$DEB_DIR/control"
Package: $APP_NAME
Version: $VERSION
Section: base
Priority: optional
Architecture: $ARCHITECTURE
Depends:
Maintainer: Your Name <your.email@example.com>
Description: A simple Fyne application
 A simple Fyne application built in Go.
EOF

# Create the DEBIAN control file (previous example)
mkdir -p "$BUILD_DIR/usr/share/applications/"

# Create the .desktop file
cat <<EOF >> "$BUILD_DIR/usr/share/applications/$APP_NAME.desktop"
[Desktop Entry]
Version=1.0
Name=$APP_NAME
Comment=A simple Fyne application
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
