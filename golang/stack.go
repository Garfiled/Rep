package stack

import (
	"errors"
	"fmt"
)

type TStack struct {
	data []interface{}
}

func NewStack() *TStack {
	return &TStack{data: make([]interface{}, 0)}
}

func (s *TStack) Push(val interface{}) {
	s.data = append(s.data, val)
}

func (s *TStack) Pop() (interface{}, error) {
	slen := len(s.data)
	if slen <= 0 {
		return nil, errors.New("no data pop")
	}
	d := s.data[slen-1]
	s.data = s.data[:slen-1]
	return d, nil
}

func (s *TStack) Ls() {
	for i, v := range s.data {
		fmt.Println(i, v)
	}
}
