package factory

import "fmt"

type Op interface {
	GetName() string
}

type Factory struct {
}

type A struct {
}
type B struct {
}

func (a *A) GetName() string {
	return "A"
}
func (b *B) GetName() string {
	return "B"
}

func (f *Factory) Create(name string) Op {
	switch name {
	case "A":
		return new(A)
	case "B":
		return new(B)
	default:
		fmt.Printf("no such struct : %s ", name)
	}

	return nil
}
