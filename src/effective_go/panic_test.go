package effective_go

import (
	"fmt"
	"log"
)

func CubeRoot(x float64) float64 {
	z := x/3
	for i := 0; i < ie6; i++ {
		prevz := z
		z -= (z*z*z-x) / (3*z*z)
		if veryClose(z, prevz) {
			return z
		}
	}
	panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}

func veryClose(x, y float64) bool {
	return false
}

type Work struct {

}

func serverWork(workChan <- chan *Work)  {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work)  {
	defer func(){
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work)
}

func do(work *Work) {

}

type Regexp struct {

}

// Error is the type of a parse error; it satisfied the error interface
type Error string

func (e Error) Error() string {
	return string(e)
}

// error is a method of *Regexp that reports parsing erros by
// panicking with an Error.
func (regexp *Regexp) error(err string)  {
	panic(Error(err))
}

// Compile returns a parsed representation of the regular expression
func Compile(str string) (regexp *Regexp, err error) {
	regexp = new(Regexp)
	defer func() {
		if e := recover(); err != nil {
			regexp = nil
			err = e.(Error)
		}
	}()
	return regexp.doParse(str), nil
}

func (regexp *Regexp) doParse(str string) (r *Regexp) {
	return nil
}
