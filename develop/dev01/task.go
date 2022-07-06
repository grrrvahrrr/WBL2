package main

import (
	"context"
	"fmt"
	"os"

	"github.com/beevik/ntp"
	"github.com/eiannone/keyboard"
)

func main() {
	//Creating context
	ctx, cancel := context.WithCancel(context.Background())

	//Creating keybard object
	if err := keyboard.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "error openning keboard: %v\n", err)
		os.Exit(1)
	}

	//Openning channel for keyboard keys
	keychan := make(chan keyboard.Key)

	//Listening to keyboard events
	go func() {
		err := listenToKeyboardEvents(keychan, cancel)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error listening to keyboard: %v\n", err)
			os.Exit(1)
		}
	}()

	//Printing time
	go func() {
		err := printTime(keychan)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error printing time: %v\n", err)
			os.Exit(1)
		}
	}()

	//Shutdown
	<-ctx.Done()
	keyboard.Close()
	fmt.Println("Shutting Down")
}

//Printing current time when recieving Ctrl+T to keyboard keys channel
func printTime(keychan chan keyboard.Key) error {
	for range keychan {
		timeNow, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
		if err != nil {
			return err
		}
		fmt.Println(timeNow)
	}
	return nil
}

func listenToKeyboardEvents(keychan chan keyboard.Key, cancel context.CancelFunc) error {
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			return err
		}
		//Ctrl+C for canceling context, Ctrl+T for printing time
		switch key {
		case keyboard.KeyCtrlT:
			keychan <- key
		case keyboard.KeyCtrlC:
			cancel()
		default:
			continue
		}
	}
}
