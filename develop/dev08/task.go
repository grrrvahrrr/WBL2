package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"WBL2/develop/dev08/errors"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// 'cd' to home with empty path not yet supported.
		if len(args) < 2 {
			return errors.ErrNoPath
		}
		if len(args) > 2 {
			return errors.ErrTooManyArgs
		}
		// Change the directory and return the error.
		return os.Chdir(args[1])
	case "pwd":
		//Print current directory
		if len(args) > 2 {
			return errors.ErrTooManyArgs
		}
		path, err := os.Getwd()
		fmt.Println(path)

		return err
	case "echo":
		//Print arguments
		if len(args) < 2 {
			return errors.ErrNoEcho
		}
		for i := 1; i < len(args); i++ {
			fmt.Printf("%s ", args[i])
			//when run out of arguments go to new line
			if i == len(args)-1 {
				fmt.Println("")
			}
		}
		return nil
	case "kill":
		//Kill process by Pid
		if len(args) < 2 {
			return errors.ErrNoProcessToKill
		}
		if len(args) > 2 {
			return errors.ErrTooManyArgs
		}

		p := os.Process{}
		var err error
		p.Pid, err = strconv.Atoi(args[1])
		if err != nil {
			return errors.ErrInvalidPid
		}
		return p.Kill()
	case "exit":
		os.Exit(0)
	}

	// Pass the program and the arguments separately.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
