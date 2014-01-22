package main

import (
	"fmt"
	"github.com/codegangsta/envy/lib"
	"io"
	"os"
	"os/exec"
)

func main() {
	// Bootstrap the environment
	envy.Bootstrap()

	// execute a command
	args := os.Args[1:]
	if len(args) == 0 {
		Exit("too few arguments")
	}

	command := exec.Command(args[0], args[1:]...)
	stdout, err := command.StdoutPipe()
	ExitIfErr(err)

	stderr, err := command.StderrPipe()
	ExitIfErr(err)

	// use goroutine to output
	err = command.Start()
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)
	command.Wait()

	ExitIfErr(err)
}

func ExitIfErr(err error) {
	if err != nil {
		Exit(err)
	}
}

func Exit(message interface{}) {
	fmt.Fprint(os.Stderr, message)
	os.Exit(1)
}
