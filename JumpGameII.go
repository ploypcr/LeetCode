package gopractice

func findMaxIndex(arr []int) int {
	maxIndex := 0
	maxValue := -999
	for i, v := range arr {
		if v >= maxValue{
			maxIndex = i
			maxValue = v
		}
	}
	return maxIndex
}
func jump(nums []int) int {
    res := 0
    if(len(nums) == 1){
        return 0
    }
	for i := 0; i < len(nums); {
		diff := make([]int,0)
        if(i+nums[i] >= len(nums)-1){
            res++
			return res
		}
		for j := i+1; j <= i+nums[i]; j++ {
			diff = append(diff, nums[j]+j)
		}
		maxIndex := findMaxIndex(diff)
		i = i+1+maxIndex
		res++
	}
	return res
}