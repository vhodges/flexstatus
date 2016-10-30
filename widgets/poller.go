package widgets

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

type Poller struct {
	BaseWidget

	Filename string

	Exec    string
	Command string   // Deprecated
	Args    []string // Deprecated

	Millis time.Duration

	output bytes.Buffer
}

func (widget *Poller) Value() string {
	return widget.output.String()
}

func (widget *Poller) Start() {

	if widget.Millis == 0 {
		widget.Millis = 1000
	}

	// Get the inital value
	if widget.Filename != "" {
		widget.readFile(false)
	} else {
		widget.runCommand(false)
	}

	ticker := time.NewTicker(time.Millisecond * widget.Millis)
	go func() {
		for _ = range ticker.C {
			if widget.Filename != "" {
				widget.readFile(true)
			} else {
				widget.runCommand(true)
			}
		}
	}()
}

func (widget *Poller) runCommand(notify bool) {

	var cmd *exec.Cmd

	if widget.Exec != "" {
		cmd = exec.Command("/bin/bash", "-c", widget.Exec)
	} else {
		cmd = exec.Command(widget.Command, widget.Args...)
	}

	out, err := cmd.Output()
	if err != nil {
		return
	}

	widget.processOutput(out, notify)
}

func (widget *Poller) readFile(notify bool) {

	out, err := ioutil.ReadFile(widget.Filename)

	if err != nil {
		return
	}

	widget.processOutput(out, notify)
}

func (widget *Poller) processOutput(out []byte, notify bool) {
	old := widget.output.String()

	s := string(out)
	s = strings.Replace(s, "\n", "", -1)

	widget.output.Reset()
	widget.output.WriteString(s)

	if notify && old != widget.output.String() {
		widget.Updated(widget)
	}
}
