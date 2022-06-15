package main

import "fmt"

//General item produced by the factory
type iBird interface {
	setName(name string)
	getName() string
	tryToFly()
}

type bird struct {
	name   string
	weight float64
	canFly bool
}

func (b *bird) setName(name string) {
	b.name = name
}

func (b *bird) getName() string {
	return b.name
}

func (b *bird) tryToFly() {
	if b.canFly {
		fmt.Printf("The %s is flying!\n", b.name)
	} else {
		fmt.Printf("The %s cannot fly.\n", b.name)
	}
}

//Factory item 1
type ostrich struct {
	bird
}

func newOstrich() iBird {
	return &ostrich{
		bird: bird{
			name:   "Common Ostrich",
			weight: 250,
			canFly: false,
		},
	}
}

//Factory item 2
type osprey struct {
	bird
}

func newOsprey() iBird {
	return &osprey{
		bird: bird{
			name:   "Western Osprey",
			weight: 3.25,
			canFly: true,
		},
	}
}

//Bird Factory
func makeBird(birdSpecies string) (iBird, error) {
	if birdSpecies == "ostrich" {
		return newOstrich(), nil
	}
	if birdSpecies == "osprey" {
		return newOsprey(), nil
	}
	return nil, fmt.Errorf("what kind of bird is that: %s?", birdSpecies)
}

func main() {
	ostrich, _ := makeBird("ostrich")
	osprey, _ := makeBird("osprey")
	_, err := makeBird("dodo")
	if err != nil {
		fmt.Println(err)
	}

	ostrich.tryToFly()
	osprey.tryToFly()
}
