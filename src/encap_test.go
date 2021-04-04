package go_learning

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
}

func TestEncap(t *testing.T) {

}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	// 返回指针
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	// e is go_learning.Employee
	t.Logf("e is %T", e)
	// e2 is *go_learning.Employee
	t.Logf("e2 is %T", e2)
}

func (e Employee) String() string {
	fmt.Printf("(String)Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 类型的指针
func (e *Employee) StringByPoint() string {
	fmt.Printf("(StringByPoint)Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	// e := Employee{"0", "Bob", 20}
	// // ID:0-Name:Bob-Age:20
	// t.Log(e.String())
	// // ID:21-Name:Tom-Age:41
	// e1 := &Employee{"21", "Tom", 41}
	// t.Log(e1.String())

	e2 := Employee{"1", "Jack", 21}
	fmt.Printf("(Employee)Address is %x\n", unsafe.Pointer(&e2.Name))
	t.Log(e2.String())
	fmt.Println("should not Same")
	// 指向类型的指针
	e3 := &Employee{"22", "Mike", 42}
	fmt.Printf("(&Employee)Address is %x\n", unsafe.Pointer(&e3.Name))
	t.Log((e3.StringByPoint()))
	fmt.Println("should Same")
}

type ListNode struct {
	Val int
	Next *ListNode
}

func TestListNode(t *testing.T) {
	temp := new(ListNode)
	fmt.Printf("temp address is %x\n", &temp)

	current := temp
	fmt.Printf("current address is %x\n", &current)

	current.Next = &ListNode{Val: 10, Next: nil}
	fmt.Printf("current.Next address is %x and Val is %d\n", &current.Next, current.Next.Val)
	fmt.Printf("temp.Next address is %x and Val is %d\n", &temp.Next, temp.Next.Val)

	temp.Next = new(ListNode)
	fmt.Printf("current.Next address is %x and Val is %d\n", &current.Next, current.Next.Val)
	fmt.Printf("temp.Next address is %x and Val is %d\n", &temp.Next, temp.Next.Val)
}
