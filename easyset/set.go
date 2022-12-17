package easyset

import (
	"errors"
	"fmt"
)

type Set struct {
	m map[interface{}]struct{}
}

func (s *Set) InitSet() {
	s.m = make(map[interface{}]struct{})
}

func (s *Set) Add(val interface{}) {
	s.m[val] = struct{}{}
}

func (s *Set) Delete(val interface{}) (err error) {
	_, ok := s.m[val]
	if !ok {
		return errors.New("the value doesn't exist in the set")
	} else {
		delete(s.m, val)
		return nil
	}
}

func (s *Set) Length() int {
	return len(s.m)
}

func (s *Set) Show() {
	for k := range s.m {
		fmt.Print(k, " ")
	}
	fmt.Println()
}

func (s *Set) Contain(val interface{}) bool {
	for k := range s.m {
		if k == val {
			return true
		}
	}
	return false
}

func (s *Set) Clear() {
	if s.Length() == 0 {
		return
	}
	for k := range s.m {
		delete(s.m, k)
	}
}
