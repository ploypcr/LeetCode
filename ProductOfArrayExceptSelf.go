package gopractice

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] , res[len(nums)-1] = 1, 1
	lsum, rsum := 1, 1
	for i := 1; i < len(nums); i++ {
		lsum *= nums[i-1]
		res[i] = lsum
	}	
	for i := len(nums)-2;i >= 0;i-- {
		rsum *= nums[i+1]
		res[i] *= rsum
	}
	return res
}