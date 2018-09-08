package mapset

import (
	"fmt"
)

type Set struct {
	data map[interface{}]bool
}

func New() *Set {

	return &Set{
		data: make(map[interface{}]bool, 10),
	}
}

func (s *Set) Add(v interface{}) {

	s.data[v] = false
}

func (s *Set) Remove(v interface{}) {

	delete(s.data, v)
}

func (s *Set) Contains(v interface{}) bool {

	_, ok := s.data[v]

	return ok
}

func (s *Set) Equals(ss *Set) bool {

	len1, len2 := len(s.data), len(ss.data)

	if len1 != len2 {
		return false
	}

	for k, _ := range s.data {

		_, ok := ss.data[k]

		if !ok {
			return false
		}
	}

	return true
}

func (s *Set) Subset(ss *Set) bool {

	len1, len2 := len(s.data), len(ss.data)

	if len1 < len2 {
		return false
	}

	for k, _ := range ss.data {

		_, ok := s.data[k]

		if !ok {
			return false
		}
	}

	return true

}

func (s *Set) Diff(ss *Set) *Set {

	sss := New()

	for k, _ := range s.data {

		_, ok := ss.data[k]

		if !ok {
			sss.data[k] = false
		}
	}

	for k, _ := range ss.data {

		_, ok := s.data[k]

		if !ok {
			sss.data[k] = false
		}
	}

	return sss
}

func (s *Set) Inter(ss *Set) *Set {

	sss := New()

	for k, _ := range s.data {

		_, ok := ss.data[k]

		if ok {
			sss.data[k] = false
		}
	}

	return sss
}

func (s *Set) Union(ss *Set) *Set {

	sss := New()

	for k, _ := range s.data {

		sss.data[k] = false
	}

	for k, _ := range ss.data {

		sss.data[k] = false
	}

	return sss
}

func (s *Set) Clear() {

	for k, _ := range s.data {

		delete(s.data, k)
	}

}

func (s *Set) Size() int {

	return len(s.data)
}

func (s *Set) Empty() bool {

	return len(s.data) == 0
}

func (s *Set) String() string {

	str := "( "

	for k, _ := range s.data {

		str = str + fmt.Sprintf("%v ", k)
	}

	str = str + ")"

	return str
}
