package go_learning

import (
	"fmt"
	"testing"
)

// 入参是一个空接口的实例
func DoSomething(p interface{}) {
  // if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("String", s)
	// 	return
	// }
	// fmt.Println("Unknow Type")
	switch v:=p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

type FileAttrs interface {
	getPath() string
}

type File struct {
	Name string
}

type BlobFile struct {
	File
}

type TextFile struct {
	File
}

func (file *File) getPath() string {
	return fmt.Sprintf("%s%T", file.Name, file)
}

func printFileAttrs(file FileAttrs) {
	fmt.Printf("file name: %s ? %T\n", file.getPath(), file)
}

func TestFile(t *testing.T) {
	file := &File{Name: "0"}
	bf := &BlobFile{File{Name: "1"}}
	tf := &TextFile{File{Name: "2"}}

	printFileAttrs(file)
	printFileAttrs(bf)
	printFileAttrs(tf)
}
