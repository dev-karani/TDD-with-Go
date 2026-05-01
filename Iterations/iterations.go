package iteration

import "strings"

func Repeat(charachter string) string {

	var repeated strings.Builder
	for range 5 {
		//concantenation is memory write heavy
		repeated.WriteString(charachter)
	}
	return repeated.String()
}

// func Repeat(charachter string) string {

// 	var repeated strings.Builder
// 	for range 5 {
// 		//concantenation is memory write heavy
// 		repeated += charachter
// 	}
// 	return repeated
// }

