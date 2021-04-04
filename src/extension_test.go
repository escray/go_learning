package go_learning

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type NewDog struct {
	// p *Pet
	Pet
}

// func (d *NewDog) Speak() {
//   // d.p.Speak()
// 	fmt.Println("Wang!")

// }

// func (d *NewDog) SpeakTo(host string) {
// 	// d.p.SpeakTo(host)
// 	d.Speak()
// 	fmt.Println(" ", host)
// }

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

func (d *Dog) SpeakTo(host string) {
	fmt.Print(" dog ", host)
}

func TestDog(t *testing.T) {
	dog := new(NewDog)
	// cannot support LSP
	// var p = (*Pet)(dog)
	// Cannot support Override
	dog.SpeakTo("Chao")
}
