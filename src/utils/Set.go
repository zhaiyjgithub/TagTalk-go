package utils

var Exists = struct {}{}
type Set struct {
	M map[interface{}]struct{}
}

func NewSet() *Set {
	s := &Set{}
	s.M = make(map[interface{}]struct{})
	return s
}

func (s *Set)Add(item interface{})  {
	s.M[item] = Exists
}

func (s *Set)AddArray(items...interface{})  {
	for _, item := range items {
		s.M[item] = Exists
	}
}

func (s *Set)Contains(item interface{}) bool {
	_, ok := s.M[item]
	return ok
}

func (s *Set)Size() int {
	return len(s.M)
}

func (s *Set)Delete(item interface{}) {
	delete(s.M, item)
}

func (s *Set)Clear()  {
	s.M = make(map[interface{}]struct{})
}

func (s *Set) Values() []interface{} {
	var values[]interface{}
	for key, _ := range s.M {
		values = append(values, key)
	}
	return values
}
