package gopractice

import (
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var res float64
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	medIndex := (float64(len(nums1))+1)/2

	if(len(nums1)%2 == 1){
		res = float64(nums1[int(medIndex)-1])
	}else{
		first := int(math.Floor(medIndex))-1
		second := first+1
		res = (float64(nums1[first])+float64(nums1[second]))/2
	}
	return res
}