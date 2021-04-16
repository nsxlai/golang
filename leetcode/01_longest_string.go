package main

import "fmt"

func main() {
	var window string
	var longestStr string
	var str string = "this is a test"

	for i := 0; len(str); i++ {
		fmt.Println(str[i])
	}
}

// def lengthOfLongestSubstring(s: str) -> int:
//     window = ''
//     longest_str = ''
//     for i_key, i_value in enumerate(s):
//         window = i_value  # starting character
//         for j in s[i_key+1:]:
//             if j not in window:
//                 window += j
//                 # print(f'window = {window}')
//             else:
//                 break
//         if len(window) > len(longest_str):
//             longest_str = window
//     return longest_str, len(longest_str)
