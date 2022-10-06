package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var serverStarted = false
var width float32 = 480
var height float32 = 480

func main() {
	SetControlPanel()
}

// TODO: Ensure the port input is only numbers. (Done)
// Ensure that text of start server changes to stop server if server is already running.
// Ensure that port panel cannot be left empty. (Done)
// Allow user to change directory to one they wish. Using GUI.

// Current Issues; The GUI part of application freezes up when the server is running.
// Logs do not get sent out to inform the user on success.

func SetControlPanel() { //GUI
	a := app.New()
	win := a.NewWindow("GoUIWebServer")
	win.SetFixedSize(true)
	win.Resize(fyne.NewSize(width, height)) //When app launches the window will be in this dimensions.
	//Fyne boilerplate code.

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter a port i.e :8080")

	//Multiple buttons can be and should be added to a single "content"
	content := container.NewVBox(
		input, widget.NewButton("Start Server", func() { receivePort(input.Text) }), //Start Server button.
		widget.NewButton("Exit Application", func() { os.Exit(0) })) //Exit button.
	win.SetContent(content)
	win.ShowAndRun()
}

func receivePort(port string) {
	if port != "" { //First if checks if port isn't empty.
		if _, err := strconv.Atoi(port); err == nil { //Second if checks whether input is a number.
			fmt.Printf("Trying port %q.\n", port)
			startServer(port)
			serverStarted = true
		} else {
			fmt.Printf("%s is not a valid port.\n", port)
		}
	} else {
		fmt.Println("Port cannot be empty, use 8080 if you wish to use default.")
	}
}

func receiveDirectory() {

}

// Start server should also receive path string to determine lcoation of directory.
func startServer(port string) {
	fileServer := http.FileServer(http.Dir("./static")) //Directory that holds html files.
	http.Handle("/", fileServer)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	if serverStarted {
		fmt.Println("Success, server is active in port: %s" + port)
	}
}

func determineArgs(arg string) {
	// Determine arguments given to the user. First one should be directory.
	switch arg {
	}
}
