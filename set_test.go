package utils

import "testing"

func TestExistFunctionForSet(t *testing.T) {
	set := NewSet[string]()
	set.Add("hellofromtheotherside")
	if !set.Exist("hellofromtheotherside") {
		t.Errorf("The exist function is not validating an existing item")
	}
}

func TestToSliceFunction(t *testing.T) {
	set := NewSet[string]()
	set.Add("hello")
	setSlice := set.ToSlice()
	if len(setSlice) < 1 || setSlice[0] != "hello" {
		t.Errorf("ToSlice Is not transfering adequetely the data from the set")
	}
}

