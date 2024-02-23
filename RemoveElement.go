package gopractice

func removeElement(nums []int, val int) int {
	var res int = 0
	if(len(nums) == 0){
		return 0
	}

	for i := 0; i < len(nums)-1; {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
			res++
		}
	}
	if(nums[len(nums)-1] == val){
		nums = nums[:len(nums)-1]
	}else{
		res++
	}
	//fmt.Printf("%v", nums)
	return res
}