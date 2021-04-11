package go_learning

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

// 空接口可以代表任意类型
func CheckType(v interface{})  {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
	CheckType(&f)
}

func TestDeepEqual(t *testing.T) {
	// map
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 4: "three"}
	// invalid operation: a == b (map can only be compared to nil)
	// t.Log(a == b)
	t.Log("a == b?", reflect.DeepEqual(a, b))

	// slice
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3", reflect.DeepEqual(s1, s3))
}

type Employeee struct {
	EmployeeID string
	Name string `format:"normal"`
	Age int
}

func (e *Employeee) UpdateAge(newVal int)  {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name string
	Age int
}

func TestInvokeByName(t *testing.T) {
	e := &Employeee{"1", "Mike", 30}
	// get member by name
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Update Age:", e)
}

func fillBySettings(st interface{}, settings map[string]interface{}) error  {
	// func(v value) Elem() Value
	// Elem returns the value that the interface v contains or that the pointer
	// It panics if v's kind is not Interface or Ptr
	// It returns the zero Value if v is nil
	// 因为要改变 st 的值，所以必须是一个指针，也必须是一个结构
	// Elem() 获取指针指向的
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}

	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok bool
	)

	for k,v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			fmt.Println("--", (reflect.ValueOf(st)).Elem())
			fmt.Println("**", (reflect.ValueOf(st)).Elem().Type())
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 40}
	e := Employeee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
