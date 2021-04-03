package my_map

import (
	"sort"
	"testing"
)

func Test(t *testing.T) {
	m1 := map[int]int{1:1, 2:4, 3:9}
	t.Log(m1[2], len(m1))
	m2 := map[int]int{}
	t.Log(len(m2))
	m2[4] = 17
	t.Log(len(m2))
	m3 := make(map[int]int, 10)
	t.Log(len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // 0
	m1[2] = 0
	t.Log(m1[2]) // 0
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else	{
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1:1, 2:2, 3:3}
	for k, v:=range m1{
		t.Log(k, v)
	}
}

func TestSortByKey(t *testing.T) {
	m1 := map[string]int{"aa": 1, "zz": 2, "bb": 3, "yy": 4, "cc": 5, "xx": 6}
	s1 := make([]string, 0, len(m1))
	for k, _ := range m1 {
		s1 = append(s1, k)
	}
	sort.Strings(s1)
	for _, k := range s1 {
		t.Log(k, m1[k])
	}

}
