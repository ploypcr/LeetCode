package gopractice

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp ListNode
	res := &tmp
	carry := 0
	for (l1 != nil || l2 != nil) || carry == 1 {
		val1 := 0
		val2 := 0

		if(l1 != nil){
			val1 = l1.Val
			
		}
		if(l2 != nil){
			val2 = l2.Val
		}

		sum := val1 + val2 + carry
		res.Val = sum % 10
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		//fmt.Printf("%v %v", l1.Val, l2.Val)
		if(l1 != nil){
            l1 = l1.Next
        }
        if(l2 != nil){
		    l2 = l2.Next
        }
        if((l1 != nil || l2 != nil) || carry == 1){
		    res.Next = &ListNode{}
		    res = res.Next
        }

	}
	return &tmp
}
