package gopractice

import "fmt"

func romanToInt(s string) int {
	roman := map[string]int{
		"I":  1,
        "IV":4,
		"V":  5,
        "IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	res := 0
    i := 0
	for ;i < len(s)-2; {
		if _, ok := roman[s[i:i+2]]; !ok {
			fmt.Printf("%v \n", roman[string(s[i])])
			res += roman[string(s[i])]
			i++
		} else {
			fmt.Printf("%v 2\n", s[i:i+2])
			res += roman[s[i:i+2]]
			i += 2
		}
	}
	
	if _, ok := roman[s[i:]]; !ok {
		for _, v := range s[i:]{
            res += roman[string(v)]
        }
	} else {
		res += roman[s[i:]]
	}
	return res
}