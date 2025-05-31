package terminal

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"termevent/rules"

	"github.com/creack/pty"
	"golang.org/x/term"
)

const SHELL = "sh"

type Terminal struct {
	inputBuffer []byte
}

func welcomeMessage() {
	fmt.Println("Welcome to the termevent")
}

func createCommand() *exec.Cmd {
	return exec.Command(SHELL)
}

func listenForInput(ptmx *os.File, t *Terminal) {
	rules_object := rules.CollectRules()
	tracker := rules.Tracker{Rules: rules_object}
	for {
		buffer := make([]byte, 1)
		os.Stdin.Read(buffer)
		ptmx.Write(buffer)
		submittedInput := t.collectInput(buffer)
		tracker.Track(submittedInput)
	}

}

func handlePtySize(ch chan os.Signal, ptmx *os.File) {
	for range ch {
		if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
			log.Printf("error resizing pty: %s", err)
		}
	}
}

func (t *Terminal) CreateTerminal() {
	welcomeMessage()
	cmd := createCommand()
	ptmx, err := pty.Start(cmd)
	if err != nil {
		panic(err)
	}
	defer func() { _ = ptmx.Close() }()
	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go handlePtySize(ch, ptmx)
	ch <- syscall.SIGWINCH                        // Initial resize.
	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	// Set stdin in raw mode.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	// NOTE: The goroutine will keep reading until the next keystroke before returning.
	go listenForInput(ptmx, t)
	_, _ = io.Copy(os.Stdout, ptmx)
}

func (t *Terminal) collectInput(input []byte) []byte {
	submittedInput := []byte{}
	if input[0] == ENTER {
		submittedInput = t.inputBuffer
		t.resetInputBuffer()
	} else {
		t.inputBuffer = append(t.inputBuffer, input...)
	}
	return submittedInput
}

func (t *Terminal) resetInputBuffer() {
	t.inputBuffer = []byte{}
}
