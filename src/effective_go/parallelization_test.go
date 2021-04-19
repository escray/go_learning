package effective_go

import "runtime"

type Vector []float64

func (v Vector) DoSome(i, n int, u Vector, c chan int)  {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1
}

func (v Vector) Op(f float64) float64  {
	return 0.0
}

// const numCPU = 4

func (v Vector) DoAll(u Vector)  {
	// var numCPU = runtime.NumCPU()
	var numCPU = runtime.GOMAXPROCS(0)
	c := make(chan int, numCPU)
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}
	for i := 0; i < numCPU; i++ {
		<- c
	}
}
