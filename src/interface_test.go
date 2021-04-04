package go_learning

import (
	"fmt"
	"testing"
	"time"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func write_sth(p Programmer) string {
	return p.WriteHelloWorld()
}

func TestClient(t *testing.T) {
	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
	t.Log(write_sth(p))
}

type IntConv func(op int) int

// func timeSpent(inner func(op int) int) func(op int) int {
func timeSpentAgain(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFuncAgain(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestTimeSpent(t *testing.T) {
	tsSF := timeSpentAgain(slowFuncAgain)
	t.Log(tsSF(10))
}

type IAnimal interface {
	Say(msg string) string
}

type Duck struct {
	Name string
}

type Dog struct {
	Name string
}

func (duck *Duck) Say(msg string) string {
	return fmt.Sprintf("duck duck! my name is %s, %s", duck.Name, msg)
}

func (dog *Dog) Say(msg string) string {
	return fmt.Sprintf("wang wang! my name is %s, %s", dog.Name, msg)
}

func TestAnimal(t *testing.T) {
	var myDuck, myDog IAnimal
	myDuck = new(Duck)
	if _, ok := myDuck.(*Duck); ok {
		t.Log("Duck is implements IAnimal")
	}

	myDog = new(Dog)
	zoo := []IAnimal{myDuck, myDog}
	for _, animal := range zoo {
		t.Log(animal.Say("hello"))
	}
}
