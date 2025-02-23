package interface_test

import (
	"reflect"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

	var prog Programmer = &GoProgrammer{}
	t.Log(prog.WriteHelloWorld())
	t.Logf("%T", reflect.TypeOf(prog))
}
