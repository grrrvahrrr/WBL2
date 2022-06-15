package main

import "fmt"

//Structs we don't want to change
type Employee interface {
	FullName()
	//Allow to accept visiors
	Accept(Visitor)
}
type Developer struct {
	FirstName string
	LastName  string
	Income    int
	Age       int
}

func (d Developer) FullName() {
	fmt.Println("Developer :", d.FirstName, " ", d.LastName)
}

//Allows to use visitor methods
func (d Developer) Accept(v Visitor) {
	v.VisitDeveloper(d)
}

type Director struct {
	FirstName string
	LastName  string
	Income    int
	Age       int
}

func (d Director) FullName() {
	fmt.Println("Director: ", d.FirstName, " ", d.LastName)
}

//Allows to use visitor methods
func (d Director) Accept(v Visitor) {
	v.VisitDirector(d)
}

//Visitor interface with structs we want to visit
type Visitor interface {
	VisitDeveloper(d Developer)
	VisitDirector(d Director)
}

//Adding calc income using visitor
type CalculIncome struct {
	bonusRate int
}

func (c CalculIncome) VisitDeveloper(d Developer) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100)
}
func (c CalculIncome) VisitDirector(d Director) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100)
}

//Adding age with visitor
type AddingCaptainAge struct {
	captainAge int
}

func (c AddingCaptainAge) VisitDeveloper(d Developer) {
	fmt.Println(d.Age + c.captainAge)
}
func (c AddingCaptainAge) VisitDirector(d Director) {
	fmt.Println(d.Age + c.captainAge)
}

func main() {
	backend := Developer{"Bob", "Bilbo", 1000, 32}
	boss := Director{"Bob", "Baggins", 2000, 40}

	backend.FullName()
	backend.Accept(CalculIncome{20})
	backend.Accept(AddingCaptainAge{42})

	boss.FullName()
	boss.Accept(CalculIncome{10})
	boss.Accept(AddingCaptainAge{42})
}
