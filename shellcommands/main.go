package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	sleepTimeout()
}

func basicCommand() {
	// create a new *Cmd instance
	// here we pass the command as the first argument and the arguments to pass to the command as the
	// remaining arguments in the function
	cmd := exec.Command("ping", "./")

	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command
	fmt.Println("Output: ", string(out))
}

func longRunningCommand() {
	cmd := exec.Command("ping", "google.com")

	// pipe the commands output to the applications
	// standard output
	cmd.Stdout = os.Stdout

	// Run still runs the command and waits for completion
	// but the output is instantly piped to Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}
}

func longRunningCommandCustomOutput() {
	cmd := exec.Command("ping", "google.com")

	cmd.Stdout = customOutput{}

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}
}

func stdinGrep() {
	cmd := exec.Command("grep", "apple")

	// Create a new pipe, which gives us a reader/writer pair
	reader, writer := io.Pipe()
	// assign the reader to Stdin for the command
	cmd.Stdin = reader
	// the output is printed to the console
	cmd.Stdout = os.Stdout

	go func() {
		defer writer.Close()
		// the writer is connected to the reader via the pipe
		// so all data written here is passed on to the commands
		// standard input
		writer.Write([]byte("1. pear\n"))
		writer.Write([]byte("2. grapes\n"))
		writer.Write([]byte("3. apple\n"))
		writer.Write([]byte("4. banana\n"))
	}()

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}
}

type customOutput struct{}

func (c customOutput) Write(p []byte) (int, error) {
	fmt.Println("received output: ", string(p))
	return len(p), nil
}

func sleepTimeout() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	cmd := exec.CommandContext(ctx, "sleep", "100")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Output: ", string(out))
}
