package main
import "fmt"
// longest prefix 

// ["flow","flower","flowing"]



// func min_string(s []string) string {

// 	min := len(s[0])
// 	minstring := s[0]

// 	for i := 1; i < len(s); i++ {
// 		length := len(s[i])

// 		if length < min {
// 			min = length
// 			minstring = s[i]
// 		}
// 	}

// 	return minstring
// }

// func max_prefix(s []string, prefix string) string {

// 	for len(prefix) > 0 {

// 		allMatch := true

// 		for i := 1; i < len(s); i++ {

// 			if s[i][:len(prefix)] != prefix {
// 				allMatch = false
// 				break
// 			}
// 		}

// 		if allMatch {
// 			return prefix
// 		}

// 		prefix = prefix[:len(prefix)-1]
// 	}

// 	return ""
// }

// func main() {

// 	s := []string{"flow", "flower", "flowing"}

// 	prefix := min_string(s)

// 	result := max_prefix(s, prefix)

// 	fmt.Println("Longest Prefix:", result)
// }

// optimized way

func min_string(s []string) string {

	min := len(s[0])
	minstring := s[0]

	for i := 1; i < len(s); i++ {
		if len(s[i]) < min {
			min = len(s[i])
			minstring = s[i]
		}
	}

	return minstring
}

func longest_prefix(s []string) string {

	if len(s) == 0 {
		return ""
	}

	minstring := min_string(s)

	for i := 0; i < len(minstring); i++ {

		ch := minstring[i]

		for j := 0; j < len(s); j++ {

			if s[j][i] != ch {
				return minstring[:i]
			}
		}
	}

	return minstring
}

func main() {

	s := []string{"flow", "flower", "flowing"}

	result := longest_prefix(s)

	fmt.Println("Longest Common Prefix:", result)
}
