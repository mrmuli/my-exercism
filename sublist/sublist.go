package sublist

import (
	"reflect"
)

type Relation string

func Sublist(la []int, lb []int) Relation {
	// var result string
	if len(la) > len(lb) {
		for i := 0; i < len(la)-len(lb)+1; i++ {
			if reflect.DeepEqual(lb, la[i:i+len(lb)]) {
				return "superlist"
			}
		}
	} else if len(la) < len(lb) {
		for i := 0; i < len(lb)-len(la)+1; i++ {
			if reflect.DeepEqual(la, lb[i:i+len(la)]) {
				return "sublist"
			}
		}
	} else if len(la) == 0 || len(lb) == 0 {
		if len(la) == 0 && len(lb) == 0 {
			return "equal"
		}
	} else if len(la) == len(lb) {
		if equal(la, lb) == true {
			return "equal"
		} else {
			return "unequal"
		}
	}
	return "unequal"
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}