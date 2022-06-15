package main

import "fmt"

func main() {
	newToy := NewToy(SpiderMan{})
	// This performs the spider man dialogue.
	newToy.PerformDialogue()
	// Change the behaviour at runtime.
	newToy.SetSuperHero(SuperMan{})
	// This performs the super man dialogue.
	newToy.PerformDialogue()
}

// DialogueReciter known how to recite a dialogue
type DialogueReciter interface {
	// Concrete types should implement this method.
	Recite()
}

// SpiderMan is a concrete type that implements Recite
type SpiderMan struct{}

// Recite -- SpiderMan says the dialogue
func (spm SpiderMan) Recite() {
	fmt.Println("No Man Can Win Every Battle, " +
		"But No Man Should Fall Without A Struggle")
}

// SuperMan is a concrete type that implements Recite
type SuperMan struct{}

// Recite -- SuperMan says the dialogue
func (sum SuperMan) Recite() {
	fmt.Println("There is a superhero in all of us, " +
		"we just need the courage to put on the cape")
}

// BatMan is a concrete type that implements Recite
type BatMan struct{}

// Recite -- BatMan says the dialogue
func (sum BatMan) Recite() {
	fmt.Println("It's not who I am underneath, " +
		"but what I do that defines me!")
}

type toy struct {
	// DialogueReciter is the behaviour that is encapsulated
	// Now this DialogueReciter is of interface type and hence
	// can hold any concrete type.
	// Now the concrete type implements the methods of the
	// DialogueReciter interface.
	DialogueReciter DialogueReciter
}

// NewToy returns a toy object
func NewToy(dr DialogueReciter) *toy {
	return &toy{
		DialogueReciter: dr,
	}
}

// PerformDialogue performs the dialogue
func (t *toy) PerformDialogue() {
	t.DialogueReciter.Recite()
}

// SetSuperHero sets the superhero for the toy
func (t *toy) SetSuperHero(dr DialogueReciter) {
	t.DialogueReciter = dr
}
