package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// All global variables are inserted here.
var (
	// Manipulate the window settings here.
	width             float32 = 480
	height            float32 = 240
	consoleEntry      int     = 0
	directoryLocation string
	stopServer        = false
	consoleOutput     = make([]string, consoleEntry)
)

// Manipulate the window settings here.

// Global so other functions can access this as well.
var a = app.New()
var win = a.NewWindow("GoUIWebServer")

var exitBtn = widget.NewButton("Exit Application", func() { os.Exit(0) }) //Fyne Button to shut app down.

func main() {
	SetConsolePanel()
	SetControlPanel()
	a.Run()
}

// TODO: Ensure the port input is only numbers. (Done)
// Ensure that text of start server changes to stop server if server is already running.
// Ensure that port panel cannot be left empty. (Done)
// Allow user to change directory to one they wish. Using GUI and commandline arguments.

/*
	Ensure that a link is generated for each html file found in static folder.
	The urls should be automatically created without any additional modifications to the script itself.
	Ensure that if the link doesn't exist, it goes straight to 404.
*/

/*
	Current Issues;
	#2 The application cannot recover from panic, if the port given is in use the application closes.
	Potential Fix: Use go's version of try catch: 'defer'

	#3 Logs do not get sent out to inform the user on success.
	Potential Fix:

	#4 Create a channel to control and stop the goroutines of the program.
*/

//TODO : Here is what we need to do to implement CONSOLE.
/*
# Create a new window to contain Console,
# Insert data into the new console via slice,
*/
func SetControlPanel() { //GUI
	win.SetFixedSize(true)
	win.Resize(fyne.NewSize(width, height)) //When app launches the window will be in this dimensions.
	//Fyne boilerplate code.

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter a port i.e 8080. Use numbers only.")

	openFolder := widget.NewButton("Select HTML Directory", func() {
		folder_Dialog := dialog.NewFolderOpen(
			func(lu fyne.ListableURI, _ error) {
				if lu != nil {
					directoryLocation = lu.String()
				}
			}, win)
		folder_Dialog.Show()
	})

	//Multiple buttons can be and should be added to a single "content"
	content := container.NewVBox(
		input, widget.NewButton("Start Server", func() { receivePort(input.Text) }), openFolder, //Start Server button.
		exitBtn)

	win.SetContent(content)
	win.Show()
}

//Control panel when the server is started.

func ControlPanelStarted(port string) {
	win.SetTitle("GoUIWebServer - Running at: " + port)
	content := container.NewVBox(
		widget.NewButton("Stop Server", func() { /*Insert Stop Action Controller here.*/ stopServer = true }), //Start Server button.
		exitBtn)
	win.SetContent(content)
}

// CONSOLE
func SetConsolePanel() {
	consoleWin := a.NewWindow("GoServer - Console")
	consoleWin.Resize(fyne.NewSize(width, height))
	list := widget.NewList(
		func() int {
			return len(consoleOutput)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(consoleOutput[i])
		})
	consoleWin.SetContent(list)
	consoleWin.Show()
}

// Perhaps this should be decoupled as well.
func receivePort(port string) {
	if port != "" {
		if _, err := strconv.Atoi(port); err == nil { //Second if checks whether input is a number.
			consoleOutput = append(consoleOutput, "Starting server at port: "+port) //Move this to its own method.                                                  //Instead of running console everytime we should make a  method that updates the console.
			//go startServer(port, directoryLocation)
			serverController(port, directoryLocation)
		} else {
			consoleOutput = append(consoleOutput, port+" is not a valid port\n")
		}
	} else {
		log.Println("Port cannot be null!")
	}
}

// Removes 'file://' from the user given directory
func filterInput(dir string) string {
	res := strings.Split(dir, "file://")
	userSelectDir := string(res[len(res)-1])
	return userSelectDir
}

// Start server should also receive path string to determine location of directory.
func startServer(port string, directory string) {
	if directory == "" {
		fileServer := http.FileServer(http.Dir("./static"))
		http.Handle("/", fileServer)
	} else {
		userSelectDir := filterInput(directory)
		fileServer := http.FileServer(http.Dir(userSelectDir))
		http.Handle("/", fileServer)
		log.Println(userSelectDir)
	}
	ControlPanelStarted(port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
func serverController(port string, directory string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		go startServer(port, directory)
		log.Println("Server Started")
		if stopServer {
			log.Println("Server Stopped")
			wg.Done()
		}
	}()
}

//Create a server controller that can control the flow of work of program.
//Server controller should have a waitground. Wg Done should be called when the user clicks Stop Server in UI.
//Once that is triggered, it should use a bool method to stop server.
