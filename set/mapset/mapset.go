package mapset

import (
	"fmt"
)

type set struct {
	data map[interface{}]bool
}

func New() *set {

	return &set{
		data: make(map[interface{}]bool, 10),
	}
}

func (s *set) Add(v interface{}) {

	s.data[v] = false
}

func (s *set) Remove(v interface{}) {

	delete(s.data, v)
}

func (s *set) Contains(v interface{}) bool {

	_, ok := s.data[v]

	return ok
}

func (s *set) Equals(ss *set) bool {

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

func (s *set) Subset(ss *set) bool {

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

func (s *set) Diff(ss *set) *set {

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

func (s *set) Inter(ss *set) *set {

	sss := New()

	for k, _ := range s.data {

		_, ok := ss.data[k]

		if ok {
			sss.data[k] = false
		}
	}

	return sss
}

func (s *set) Union(ss *set) *set {

	sss := New()

	for k, _ := range s.data {

		sss.data[k] = false
	}

	for k, _ := range ss.data {

		sss.data[k] = false
	}

	return sss
}

func (s *set) Clear() {

	for k, _ := range s.data {

		delete(s.data, k)
	}

}

func (s *set) Size() int {

	return len(s.data)
}

func (s *set) Empty() bool {

	return len(s.data) == 0
}

func (s *set) String() string {

	str := "( "

	for k, _ := range s.data {

		str = str + fmt.Sprintf("%v ", k)
	}

	str = str + ")"

	return str
}
