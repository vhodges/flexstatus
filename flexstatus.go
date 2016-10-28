package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/vhodges/flexstatus/widgets"
)

func main() {

	updates := make(chan widgets.Widget, 10) // Buffered

	// Get filename for the set of widgets
	home := os.Getenv("HOME")
	defaultFilename := filepath.Join(home, ".flexstatus_widgets.toml")

	configFilePtr := flag.String("config", defaultFilename,
		"pathname for toml file listing the widgets to display")

	flag.Parse()

	// Load widgets from config file.if none found we'll put two in for
	// them to see something
	active_widgets := get_widgets(*configFilePtr)

	// Set them up
	for i := range active_widgets {
		// Channel to send paint events on
		active_widgets[i].SetUpdateChannel(updates)
		active_widgets[i].Start() // start them up.
	}

	// Render the first string for the widgets
	// (Start() will have filled in the initial value for display)
	renderWidgets(active_widgets)

	// When widgets change, we get a message on this channel to repaint
	for _ = range updates {
		renderWidgets(active_widgets)
	}
}

func renderWidgets(active_widgets []widgets.Widget) {
	var buffer bytes.Buffer

	for i := range active_widgets {
		buffer.WriteString(widgets.Render(active_widgets[i]))
	}

	fmt.Println(buffer.String())
}
