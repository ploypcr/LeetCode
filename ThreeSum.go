package gopractice

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; {
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}
		start := i + 1
		end := len(nums) - 1
		sumArr := make([]int, 0)
		for start < end {
			if nums[start]+nums[end] == nums[i]*-1 {
				sumArr = append(sumArr, nums[i], nums[start], nums[end])
				if len(sumArr) != 0 {
					res = append(res, sumArr)
				}
				sumArr = make([]int, 0)
				start++
				end--
			} else {
				if nums[start]+nums[end] > nums[i]*-1 {
					end--
				} else {
					start++
				}
			}
			if start > i+1 {
				for nums[start] == nums[start-1] && start < end {
					start++
				}
			}
			if end < len(nums)-1 {
				for nums[end] == nums[end+1] && end > start {
					end--
				}
			}
			//fmt.Printf("%v %v\n",start,end)
		}
		i++

	}
	return res
}