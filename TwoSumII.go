package gopractice
func twoSum(numbers []int, target int) []int {
	res := make([]int,0)
	for i := 0; i < len(numbers); i++ {
		for j := i+1; j < len(numbers); j++ {
			checkNum := numbers[i]+numbers[j]
			if checkNum == target{
				res = append(res, i+1,j+1)
			}
		}
	}  
	return res
}