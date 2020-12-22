package ds

import (
	"fmt"
)

//Set is a set
type Set map[interface{}]interface{}

//NewSet creates a set
func NewSet() Set {
	return make(Set)
}

// Size returns the size of the set
func (s Set) Size() int {
	return len(s)
}

//Add adds an element to the set and returns true if at least one element has been added
func (s Set) Add(e ...interface{}) bool {
	done := false
	for _, f := range e {
		if _, ok := s[f]; ok {
			continue
		}
		s[f] = nil
		_, ok := s[f]
		done = ok
	}
	return done
}

//Contains returns true if the set contains e
func (s Set) Contains(e interface{}) bool {
	_, ok := s[e]
	return ok
}

// Delete delete the given element from the set if present and returns true if element is absent after.
func (s Set) Delete(e interface{}) bool {
	delete(s, e)
	if _, ok := s[e]; !ok {
		return !ok
	}
	return false
}

// Pop returns and deletes a random element
func (s Set) Pop() interface{} {
	for k := range s {
		delete(s, k)
		return k
	}
	return nil
}

// Peek returns a random element from the set
func (s Set) Peek() interface{} {
	for k := range s {
		return k
	}
	return nil
}

func (s Set) String() string {
	str := "{"
	i := 0
	for k := range s {
		str += fmt.Sprintf("%v", k)
		if i != len(s)-1 {
			str += ", "
		}
		i++
	}
	str += "}"
	return str
}
