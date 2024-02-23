package gopractice

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	res := make([]string,0)
	arr := strings.Split(s, " ")
	for i := len(arr)-1; i >= 0; i--{
		arr[i] = strings.TrimSpace(arr[i])
        
		if(!strings.Contains(arr[i]," ") && arr[i] != ""){
			res = append(res,arr[i])
		}
	}
    //fmt.Printf("%v", len(res))
	return strings.Join(res, " ")
}