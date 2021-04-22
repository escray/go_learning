package go_paradigms

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"testing"
)

type Point struct {
	Logitude string
	Latitude string
	Distance string
	ElevationGain string
	ElevationLoss string
}

// lose it
// Panic attack
// reaching breaking point
// nervous/mental breakdown
func parseOriginFreakingOut(r io.Reader) (*Point, error) {
	var p Point
	if err := binary.Read(r, binary.BigEndian, &p.Logitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
		return nil, err
	}
	return &p, nil
}

func parseFunctional(r io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Logitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}

	return &p, nil
}

type Reader struct {
	r io.Reader
	err error
}

func (r *Reader) read(data interface{})  {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

func parse(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Logitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}

	return &p, nil
}

var b = []byte {0x48, 0x61, 0x6f, 0x20, 0x43, 0x68, 0x65, 0x6e, 0x00, 0x00, 0x2c}
var r  = bytes.NewReader(b)

type PersonAgain struct {
	Name [10]byte
	Age uint8
	Weight uint8
	err error
}

func (p *PersonAgain) read(data interface{})  {
	if p.err == nil {
		p.err = binary.Read(r, binary.BigEndian, data)
	}
}

func (p *PersonAgain) ReadName() *PersonAgain {
	p.read(&p.Name)
	return p
}

func (p *PersonAgain) ReadAge() *PersonAgain {
	p.read(&p.Age)
	return p
}

func (p *PersonAgain) ReadWeight() *PersonAgain {
	p.read(&p.Weight)
	return p
}

func (p *PersonAgain) Print() *PersonAgain {
	if p.err == nil {
		fmt.Printf("Name=%s, Age=%d, Weight=%d\n", p.Name, p.Age, p.Weight)
	}
	return p
}

func TestErrorHandle(t *testing.T) {
	p := PersonAgain{}
	p.ReadName().ReadAge().ReadWeight().Print()
	fmt.Println(p.err)

	// _, _ := fmt.Println("print debug")
	// _, _ := fmt.Println("here")
	// _, _ := fmt.Println("")
}

type errWriter struct {
	w io.Writer
	err error
}

func (e *errWriter) Write(p []byte)  {
	if e.err != nil {
		return
	}
	_, e.err = e.w.Write(p)
}

func (e *errWriter) Err() error {
	return e.err
}



func do() error {
	ew := &errWriter {
		// w:
	}
	buf := []byte {}
	ew.Write(buf)
	ew.Write(buf)

	if ew.Err() != nil {
		return ew.Err()
	}
	return nil
}
