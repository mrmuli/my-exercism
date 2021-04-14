package sublist

import (
	"fmt"
	"reflect"
)


func Sublist(la []int, lb []int) string {
	var result string
	if len(la) > len(lb) {
		fmt.Println("Seems la is greater")
	} else if len(la) < len(lb) {
		fmt.Println("Seems lb is greater")

		for i := 0; i < len(lb)-len(la)+1; i++ {
			if reflect.DeepEqual(la, lb[i:i+len(la)]) {
				return "sublist"
			}
		}
	} else if len(la) == 0 || len(lb) == 0 {
		if len(la) == 0 {
			result = "sublist"
		} else {
			result = "superlist"
		}
	} else if len(la) == len(lb) {
		if equal(la, lb) == true {
			result = "equal"
		} else {
			result = "unequal"
		}
	} else {
		fmt.Println("special case for la, lb", len(la), len(lb))
	}
	return result
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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