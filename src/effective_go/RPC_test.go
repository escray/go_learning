package effective_go

type Buffer struct {

}

func load(b *Buffer)  {

}

func processs(b *Buffer) {

}

var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
	for {
		var b *Buffer
		select {
		case b = <-freeList:
			// Got one; nothing more todo
		default:
			// None free, so allocate a new one
			b = new(Buffer)
		}
		load(b)
		serverChan <- b
	}
}

func server()  {
	for {
		b := <-serverChan
		processs(b)
		select {
		case freeList <- b:
			// Buffer on free list; nothing more to do
		default:
			// Free list full, just carry on
		}
	}
}
