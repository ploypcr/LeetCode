package gopractice

import "math"

func findMinStr(arr []string) string {
	minLen := math.MaxInt64
	minStr := ""
	for _, v := range arr {
		if len(v) < minLen {
			minLen = len(v)
			minStr = v
		}

	}
	return minStr
}
func longestCommonPrefix(strs []string) string {
	str := findMinStr(strs)

	for str != "" {
		check := true
		for _, v := range strs {
			if v[:len(str)] != str {
				check = false
				str = str[:len(str)-1]
			}
		}
		if check == true {
			return str
		}
	}
	return str
}