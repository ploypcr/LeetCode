package gopractice

import "sort"

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	maxHIndex := 0
	for i, v := range citations {
		//mt.Printf("%v %v\n", len(citations[i:]), v)
		if len(citations[:i+1]) <= v {
			maxHIndex++
		}
	}
	return maxHIndex
}