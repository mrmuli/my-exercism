package piglatin

import (
	"strings"
)

/*
https://play.golang.org/p/M5f24xRRCJI

Check if a word begins with a vowel sound
add an "ay" sound to the end of the word

Check if word begins with a constant followed by "qu"
move constant and "qu" to end
add an "ay" to the end

Check if a word begins with a consonant sound
move it to the end of the word and then add an "ay"

Check if a word contains a "y" after a consonant cluster
Check whether input is a 2 letter word with y as second cluster.
Both equate to vowel sounds.
*/

// func Sentence(s string) string {

// 	vowels := []string{"a", "e", "i", "o", "u", "xr", "yt"}
// 	var str string

// 	for _, i := range vowels {
// 		if strings.HasPrefix(s, i) {
// 			str = s + "ay"
// 			return str
// 		}
// 	}
// 	// capture sound by slicing by the next vowel sound.
// 	var firstVowel int
// 	runeSlice := []int{}
// 	for i, j := range s {
// 		for _, k := range vowels {
// 			if string(j) == k {
// 				runeSlice = append(runeSlice, i)
// 			}
// 		}
// 	}
// 	// Use firstVowel to slice the first sound.
// 	firstVowel = runeSlice[0]

// 	con := s[:firstVowel]
// 	soundSlicePos := firstVowel + 1
// 	sound := s[:soundSlicePos]

// 	fmt.Println("########")

// 	if len(s) == 2 {
// 		str = string(s[1]) + string(s[0]) + "ay"
// 	} else if strings.Contains(s, "qu") {
// 		str = s[soundSlicePos:] + sound + "ay"
// 	} else {
// 		str = s[firstVowel:] + con + "ay"
// 	}
// 	return str
// }
func Sentence(s string) string {

	vowels := []string{"a", "e", "i", "o", "u", "xr", "yt"}
	var str string

	for _, i := range vowels {
		if strings.HasPrefix(s, i) {
			str = s + "ay"
			return str
		}
	}

	var con string
	var sound string
	var soundSlicePos int
	runeSlice := []int{}
	for i, j := range s {
		for _, k := range vowels {
			if string(j) == k {
				runeSlice = append(runeSlice, i)
			}
		}
	}

	if len(s) == 2 {
		str = string(s[1]) + string(s[0]) + "ay"
	} else if strings.Contains(s, "qu") {
		soundSlicePos = runeSlice[0] + 1
		sound = s[:soundSlicePos]
		str = s[soundSlicePos:] + sound + "ay"
	} else {
		con = s[:runeSlice[0]]
		str = s[runeSlice[0]:] + con + "ay"
	}
	return str
}
