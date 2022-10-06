package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var serverStarted = false

func main() {
	SetControlPanel()

	//fmt.Printf("Starting server at port: ",port)
}

// TODO: Ensure the port input is only numbers. (Done)
// Ensure that text of start server changes to stop server if server is already running.
// Ensure that port panel cannot be left empty. (Done)
// Allow user to change directory to one they wish. Using GUI.

func SetControlPanel() { //GUI
	a := app.New()
	win := a.NewWindow("GoUIWebServer") //Creates a new window and gives it Title.
	win.SetFixedSize(true)              //Ensures you cant resize.
	win.Resize(fyne.NewSize(480, 480))  //When app launches the window will be in this dimensions.
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
	if port != "" { //First if checks if port is empty.
		if _, err := strconv.Atoi(port); err == nil { //Second if checks whether input is a number.
			fmt.Printf("Trying port %q.\n", port)
			//serverStarted = true
		} else { //If it is this message will be given out.
			fmt.Printf("%s is not a valid port.\n", port)
		}
	} else { //If it is this message will be given out.
		fmt.Println("Port cannot be empty, use 8080 if you wish to use default.")
	}
}
func receiveDirectory() {

}
func startServer(path string) {
	path = "./static"
	//fileServer := http.FileServer(http.Dir("./static")) //Directory that holds html files.
	fileServer := http.FileServer(http.Dir("./%s" + path))
	http.Handle("/", fileServer)
	//http.HandleFunc("/form",formHandler)
	//http.HandleFunc("/Readme",readmeHandler)
}
