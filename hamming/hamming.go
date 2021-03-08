package hamming

import "errors"

// Distance calculates hamming distance between
// two strings
func Distance(a, b string) (int, error) {

	switch {
	case len(a) != len(b):
		return 0, errors.New("lengths are not equal")
	case len(a) == 0 || len(b) == 0:
		return 0, errors.New("lengths are not equal")
	case a == b:
		return 0, errors.New("lengths are not equal")
	}
	var ham int
	for i := range a {
		var str1 string
		str := string(a[i])
		for j := range b {
			str1 = string(b[j])
		}
		if str != str1 {
			ham++
		}
	}

	return ham, nil
}

// Distanze seemed to be a more effective approach from the
// community, one loop can iterate same byte-length types.
func Distanze(a, b string) (int, error) {

	switch {
	case len(a) != len(b):
		return 0, errors.New("lengths are not equal")
	}
	var ham int
	for i := range a {
		if a[i] != b[i] {
			ham++
		}
	}

	return ham, nil
}
