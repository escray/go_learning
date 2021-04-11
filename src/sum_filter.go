package go_learning

import "errors"

var ErrorSumFilterWrongFormat = errors.New("input data should be []int")

type SumFilter struct {

}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int)
	if !ok {
		return nil, ErrorSumFilterWrongFormat
	}
	ret := 0
	for _, v := range elems {
		ret += v
	}
	return ret, nil
}
