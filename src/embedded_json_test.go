package go_learning

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
  "basic_info": {
		"name":"Mike",
		"age":30
	},
	"job_info": {
		"skills":["Java","Go","C"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(EmployeeJSON)
	fmt.Println([]byte(jsonStr))
	err := json.Unmarshal([]byte(jsonStr), e)
  if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	fmt.Println(e)
	fmt.Println(&e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
		fmt.Println(v)
	} else {
		t.Error(err)
	}
}

func TestEasyJSON(t *testing.T) {
	e := EmployeeJSON{}
	e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(e)
	if v, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(v))
	}
}

func TestEasyJSONList(t *testing.T) {
	e1 := BasicInfo{"Mike", 10}
	e2 := BasicInfo{"Rose", 11}
	eList := BasicInfoList{e1, e2}
	v, err := eList.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(v))
}

func BenchmarkEmbeddedJSON(b *testing.B) {
	b.ResetTimer()
	e := new(EmployeeJSON)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err = json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyJSON(b *testing.B) {
	b.ResetTimer()
	e := EmployeeJSON{}
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			b.Error(err)
		}
		if _, err = e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}
