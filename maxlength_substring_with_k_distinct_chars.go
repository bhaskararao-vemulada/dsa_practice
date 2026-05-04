package main
import "fmt"

func longestsubstring_with_k_distinct_chars(s string, k int) int {
	left := 0
	maxlength := 0
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m[s[i]]++

		// 🔥 shrink only when invalid
		for len(m) > k {
			m[s[left]]--
			if m[s[left]] == 0 {
				delete(m, s[left])
			}
			left++
		}

		length := i - left + 1
		if length > maxlength {
			maxlength = length
		}
	}

	return maxlength
}

func main() {
	s := "aabbcc"
	k := 2
	fmt.Println(longestsubstring_with_k_distinct_chars(s, k)) // 4
}