package main

import "fmt"

//TV Button invoker
type tvOnOffButton struct {
	command command
}

func (b *tvOnOffButton) press() {
	b.command.execute()
}

//Remote button invoker
type remoteOnOffButton struct {
	command command
}

func (b *remoteOnOffButton) press() {
	b.command.execute()
}

//Command interface
type command interface {
	execute()
}

//OFF Concrete command
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

//ON Concrete command
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

//Interface for recievers
type device interface {
	on()
	off()
}

//Command reciever
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	//Create reviever
	tv := &tv{}
	//Pass reciever to commands
	onCommand := &onCommand{
		device: tv,
	}
	offCommand := &offCommand{
		device: tv,
	}
	//Use tv button invoker
	onButton := &tvOnOffButton{
		command: onCommand,
	}
	onButton.press()
	//Use remote button invoker
	offButton := &remoteOnOffButton{
		command: offCommand,
	}
	offButton.press()
}
