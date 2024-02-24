package gopractice

func majorityElement(nums []int) int {
	element := make(map[int]int)
	for _, v := range nums {
		element[v]++
	}
    majorElement := 0
	count := 0
    for k, v := range element{
		if v > count{
            //fmt.Printf("%v",v)
            count = v
			majorElement = k
		}
	}
	
	return majorElement
}