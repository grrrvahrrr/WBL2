package main

import (
	"context"
	"testing"
	"time"

	"github.com/eiannone/keyboard"
)

func TestPrintTime(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	keychan := make(chan keyboard.Key)

	go func() {
		err := printTime(keychan)
		if err != nil {
			t.Error("error connecting to ntp server: ", err)
		}
	}()

	go func() {
		for {
			keychan <- keyboard.KeyCtrlT
			time.Sleep(time.Second)
		}
	}()

	<-ctx.Done()

}
