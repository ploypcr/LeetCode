package gopractice

func removeDuplicates(nums []int) int {
    numElements := make(map[int]int)
	res := 0
	
	if(len(nums)==0){
		return res
	}

	for i := 0; i < len(nums)-1; {
		numElements[nums[i]]++
		if(numElements[nums[i]] > 2){
			nums = append(nums[:i],nums[i+1:]...)
		}else{
			i++;
			res++;
		}
	}

	if((numElements[nums[len(nums)-1]] + 1) <= 2){
		res++;
	}
	return res
}
