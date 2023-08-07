package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	// Create a new app instance
	app := app.New()

	// Create a new window with a greeting message
	window := app.NewWindow("Наша первая программа!")

	// Create the UI components
	output, entry, btnText := myApp.makeUI()

	// Set the content of the window to a vertical box containing the UI components
	window.SetContent(container.NewVBox(output, entry, btnText))

	// Show the window and start the application event loop
	window.ShowAndRun()
}

// makeUI creates and initializes the UI components.
// It returns the output label, entry field, and button.
func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Привет, мир!") // create output label with initial text
	entry := widget.NewEntry() // create entry field
	btnText := widget.NewButton("Click me!", func() { // create button with click event handler
		app.output.SetText(entry.Text) // set the output label text to the entry field text
	})
	

	app.output = output // assign the output label to the app's output field

	return output, entry, btnText // return the UI components
}
