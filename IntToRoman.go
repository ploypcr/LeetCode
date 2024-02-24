package gopractice

import "strings"

func intToRoman(num int) string {
	res := ""
	thousand, left := num/1000, num%1000
	hunred, left := left/100, left%100
	ten, left := left/10, left%10

	res += strings.Repeat("M",thousand)

	if(hunred >= 9){
		res += "CM"
	}else{
		if(hunred >= 5){
			res += "D"
			res += strings.Repeat("C",hunred-5)
		}else{
			if(hunred >= 4){
				res += "CD"
			}else{
				res += strings.Repeat("C",hunred)
			}
		}
	}


	if(ten >= 9){
		res += "XC"
	}else{
		if(ten >= 5){
			res += "L"
			res += strings.Repeat("X",ten-5)

		}else{
			if(ten >= 4){
				res += "XL"
			}else{
				res += strings.Repeat("X",ten)
			}
		}
	}

	if(left >= 9){
		res += "IX"
	}else{
		if(left >= 5){
			res += "V"
			res += strings.Repeat("I",left-5)

		}else{
			if(left >= 4){
				res += "IV"
			}else{
				res += strings.Repeat("I",left)
			}
		}
	}
	return res
}