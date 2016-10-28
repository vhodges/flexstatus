package widgets

import (
	"bufio"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type Desktop struct {
	BaseWidget

	DesktopNames []string
	Count        int64
	Current      int64
}

func (widget *Desktop) Title() string {
	return widget.DesktopNames[widget.Current-1]
}

func (widget *Desktop) Start() {

	// We need to block until we have all of the data we need that might be rendered/
	var wg sync.WaitGroup
	waiting := true
	wg.Add(1)

	go func() {

		// Warning: xprop -spy is known to leak memory, but has been supposedly fixed upstream (and in Arch?)
		spyCmd := exec.Command("xprop", "-root", "-notype", "-spy", "_NET_NUMBER_OF_DESKTOPS", "_NET_DESKTOP_NAMES", "_NET_CURRENT_DESKTOP")
		spyOut, _ := spyCmd.StdoutPipe()
		spyCmd.Start()

		scanner := bufio.NewScanner(spyOut)

		for scanner.Scan() {

			widget.processLine(scanner.Text())

			// Do we have all the info?
			if len(widget.DesktopNames) > 0 && widget.Count != 0 && widget.Current != 0 {

				// Were we still gathering info for the initial render?
				if waiting {
					waiting = false
					wg.Done()
				} else {
					widget.Updated(widget)
				}
			}
		}
	}()

	// Block here.
	wg.Wait()
}

func (widget *Desktop) processLine(s string) {

	strs := strings.Split(s, " = ")

	// Sometimes lines come in that are malformed (error messages etc)
	if len(strs) < 2 {
		return
	}

	switch strs[0] {
	case "_NET_NUMBER_OF_DESKTOPS":
		if s, err := strconv.ParseInt(strs[1], 10, 64); err == nil {
			widget.Count = s
		}
	case "_NET_DESKTOP_NAMES":
		widget.DesktopNames = strings.Split(strs[1], ", ")

		// Strip the quotes from around the desktop names
		for i := range widget.DesktopNames {
			widget.DesktopNames[i] = strings.Replace(widget.DesktopNames[i], "\"", "", -1)
		}
	case "_NET_CURRENT_DESKTOP":
		if s, err := strconv.ParseInt(strs[1], 10, 64); err == nil {
			widget.Current = s + 1 // Make them 1 based
		}
	}
}
