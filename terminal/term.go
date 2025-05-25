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

type Terminal struct{}

func welcomeMessage() {
	fmt.Println("Welcome to the termevent")
}

func createCommand() *exec.Cmd {
	return exec.Command(SHELL)
}

func listenForInput(ptmx *os.File) {
	rules_object := []rules.Rule{rules.Create([]byte("hello"))}
	tracker := rules.Tracker{Rules: rules_object}
	for {
		buffer := make([]byte, 1)
		os.Stdin.Read(buffer)
		ptmx.Write(buffer)
		tracker.Track(buffer)
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
	go listenForInput(ptmx)
	_, _ = io.Copy(os.Stdout, ptmx)
}
