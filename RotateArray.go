package gopractice

func rotate(nums []int, k int)  {
	newArr := make([]int, len(nums))
	for i, v := range nums {
		newIndex := ((i + k) % len(nums))
		newArr[newIndex] = v
	}
	
	for i := 0; i < len(nums); i++ {
		nums[i] = newArr[i]
	}
}