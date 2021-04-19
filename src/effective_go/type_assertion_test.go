package effective_go

import "fmt"

type Stringer interface {
	String() string
}

func TypeAssertion() string{
	var value interface{}

	switch str := value.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}
	return ""
}

func StringAssertion(value interface{}) {
	str, ok := value.(string)
	if ok {
		fmt.Printf("string value is %q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}

func StringAssertionAgain(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	} else if str, ok := value.(Stringer); ok {
		return str.String()
	}
	return ""
}
