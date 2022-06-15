package main

import "fmt"

type CellPhone struct {
	Camera       bool
	DualSim      bool
	Torch        bool
	ColorDisplay bool
}

//Builder Interface for all concrete builders
type BuildProcess interface {
	SetCamera() BuildProcess
	SetDualSim() BuildProcess
	SetTorch() BuildProcess
	SetColorDisplay() BuildProcess
	GetCellPhone() CellPhone
}

//Nokia Concrete builder
type Nokia struct {
	Phone CellPhone
}

func (n *Nokia) SetCamera() BuildProcess {
	n.Phone.Camera = false
	return n
}

func (n *Nokia) SetDualSim() BuildProcess {
	n.Phone.DualSim = false
	return n
}

func (n *Nokia) SetTorch() BuildProcess {
	n.Phone.Torch = true
	return n
}

func (n *Nokia) SetColorDisplay() BuildProcess {
	n.Phone.ColorDisplay = false
	return n
}

func (n *Nokia) GetCellPhone() CellPhone {
	return n.Phone
}

//Samsung concrete builder
type Samsung struct {
	Phone CellPhone
}

func (s *Samsung) SetCamera() BuildProcess {
	s.Phone.Camera = true
	return s
}

func (s *Samsung) SetDualSim() BuildProcess {
	s.Phone.DualSim = true
	return s
}

func (s *Samsung) SetTorch() BuildProcess {
	s.Phone.Torch = false
	return s
}

func (s *Samsung) SetColorDisplay() BuildProcess {
	s.Phone.ColorDisplay = true
	return s
}

func (s *Samsung) GetCellPhone() CellPhone {
	return s.Phone
}

//Build Director
type Director struct {
	builder BuildProcess
}

func (d *Director) Construct(b BuildProcess) CellPhone {
	d.builder = b
	d.builder.SetCamera().SetDualSim().SetTorch().SetColorDisplay()
	return d.builder.GetCellPhone()
}

func main() {
	diro := Director{}
	nokiaPhone := &Nokia{}
	phone := diro.Construct(nokiaPhone)
	fmt.Println(phone)

	samsungPhone := &Samsung{}
	phone = diro.Construct(samsungPhone)
	fmt.Println(phone)
}
