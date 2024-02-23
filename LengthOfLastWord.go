package gopractice

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	arr := strings.Split(s," ")
	fmt.Printf("%v", arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.TrimSpace(arr[i])
	}
	return len(arr[len(arr)-1])
}