package go_learning

import (
	"fmt"
	"testing"
)

type Code string

type ProgrammerAgain interface {
	WriteHelloWorld() Code
}

type GoProgrammerAgain struct {
}

func (p *GoProgrammerAgain) WriteHelloWorld() Code {
  return "fmt.Println(\"Hello World!\") -- Go"
}

type JavaProgrammerAgain struct {
}

func (p *JavaProgrammerAgain) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\") -- Java"
}

// 这里传入的参数 p ProgrammerAgain 是一个 interface
// 所以在调用的时候，也只能传递过来一个指针/引用
func writeFirstProgram(p ProgrammerAgain) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := &GoProgrammerAgain{}
  // goProg := new(GoProgrammerAgain)
	javaProg := new(JavaProgrammerAgain)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
