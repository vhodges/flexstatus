package widgets

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Streamer struct {
	BaseWidget

	Filename string

	Command string
	Args    []string

	output bytes.Buffer
}

func (widget *Streamer) Value() string {
	return widget.output.String()
}

func (widget *Streamer) Start() {

	var wg sync.WaitGroup
	waiting := true
	wg.Add(1)

	go func() {

		var stream io.Reader
		var err error

		if widget.Filename != "" {
			stream, err = os.Open(widget.Filename)

			if err != nil {
				fmt.Println(err)
				return
			}

		} else {
			spyCmd := exec.Command(widget.Command, widget.Args...)
			stream, _ = spyCmd.StdoutPipe()
			spyCmd.Start()
		}

		scanner := bufio.NewScanner(stream)

		for scanner.Scan() {

			s := scanner.Text()

			old := widget.output.String()

			s = strings.Replace(s, "\n", "", -1)

			widget.output.Reset()
			widget.output.WriteString(s)

			// Wait for the first line of data before letting Start() return
			if waiting {
				waiting = false
				wg.Done()
			} else {
				if old != widget.output.String() {
					widget.Updated(widget)
				}
			}
		}
	}()

	// Block here.
	wg.Wait()
}
