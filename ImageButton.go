package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// ImageButton is a custom widget that includes an image and a label
type ImageButton struct {
	widget.BaseWidget
	Image    *canvas.Image
	OnTapped func()
}

// NewCustomButton creates a new ImageButton
func NewCustomButton(imagePath string, labelText string, onTapped func()) *ImageButton {
	btn := &ImageButton{
		Image:    canvas.NewImageFromFile(imagePath),
		OnTapped: onTapped,
	}

	// Set the initial size of the image
	btn.Image.FillMode = canvas.ImageFillContain
	btn.Image.FillMode = canvas.ImageFillStretch // Stretch the image to fill the container
	btn.Image.SetMinSize(fyne.NewSize(300, 300)) // Change this size as needed

	btn.ExtendBaseWidget(btn)
	return btn
}

// Tapped is triggered when the widget is tapped (for mobile and regular clicks)
func (b *ImageButton) Tapped(_ *fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped()
	}
}

// MouseIn, MouseOut and MouseMoved implement hover effects for desktop (optional)
func (b *ImageButton) MouseIn(_ *desktop.MouseEvent) {
	b.Image.FillMode = canvas.ImageFillOriginal // Example hover effect (optional)
}

func (b *ImageButton) MouseOut() {
	b.Image.FillMode = canvas.ImageFillContain
}

func (b *ImageButton) MouseMoved(_ *desktop.MouseEvent) {}

// CreateRenderer defines how the widget is rendered
func (b *ImageButton) CreateRenderer() fyne.WidgetRenderer {
	// Layout the button vertically with image on top and label below
	container := container.NewVBox(b.Image)

	// Optional: Add a background (for visual distinction)
	background := canvas.NewRectangle(color.NRGBA{R: 0, G: 228, B: 255, A: 255})

	// Overlay the button content on top of the background
	objects := []fyne.CanvasObject{background, container}

	return &customButtonRenderer{
		button:  b,
		objects: objects,
		layout:  container,
		bg:      background,
	}
}

// customButtonRenderer handles the layout and rendering
type customButtonRenderer struct {
	button  *ImageButton
	objects []fyne.CanvasObject
	layout  *fyne.Container
	bg      *canvas.Rectangle
}

// Layout arranges the objects in the custom button
func (r *customButtonRenderer) Layout(size fyne.Size) {
	r.bg.Resize(size)     // Resize the background
	r.layout.Resize(size) // Resize the VBox layout
}

// MinSize returns the minimum size of the custom button
func (r *customButtonRenderer) MinSize() fyne.Size {
	return r.layout.MinSize() // Get the min size from VBox layout
}

// Refresh is used to update the UI if the button changes
func (r *customButtonRenderer) Refresh() {
	r.bg.FillColor = color.NRGBA{R: 0, G: 128, B: 255, A: 255} // Optional refresh of background color
	canvas.Refresh(r.button)
}

// Objects returns all the objects that the renderer manages
func (r *customButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

// Destroy cleans up any resources (not needed for this example)
func (r *customButtonRenderer) Destroy() {}
