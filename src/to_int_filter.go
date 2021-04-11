package go_learning

import (
	"errors"
	"strconv"
)

var ErrorToIntFilterWrongFormat = errors.New("ipnut data should be []string")

type ToIntFilter struct {

}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tif *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ErrorToIntFilterWrongFormat
	}
	ret := []int{}
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}


