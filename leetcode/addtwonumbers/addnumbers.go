package addtwonumbers

type LinkedList struct{
	Head *ListNode 
	Tail *ListNode
}

type ListNode struct{
	Val int 
	Next *ListNode
}

func (l *ListNode) NewList(data []int) *ListNode{
	current := l 
	for _, value := range data{
		current.Val = value 
		current = current.Next
	}

	return l
}

func (l *ListNode) ToArray() []int{
	var array []int
	for node := l; node != nil; node = node.Next{
		array = append(array, node.Val)
	}
	return array
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	if(l1 == nil && l2 == nil){
		return nil
	}
	if(l1 == nil){
		return l2
	}
	if(l2 == nil){
		return l1
	}
	carry := 0 
	dummy := &ListNode{Val: 0} 
	temp := dummy 
	for l1 != nil && l2 != nil{
		x, y := 0, 0 
		if l1 != nil{
			x = l1.Val
		} 
		if l2 != nil{
			y = l2.Val
		}
		sum := x + y + carry 
		carry = int(sum/10) 
		temp.Next = &ListNode{Val: int(sum/10)} 
		temp = temp.Next 
		if l1 != nil{
			l1 = l1.Next
		}
		if l2 != nil{
			l2 = l2.Next
		}
	}
	if carry > 0{
		temp.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

// Equal function to compare two arrays
func Equal(a, b []int) bool{
	if len(a) != len(b){
		return false 
	}
	for i, v := range a{
		if v != b[i]{
			return false
		}
	}
	return true 
}

// EqualList function to compare two ListNode objects 
func EqualList(l1, l2 *ListNode) bool{
	// Implementation goes here 
	// ... 
	for list_a, list_b := l1, l2; list_a != nil && list_b != nil; list_a, list_b = list_a.Next, list_b.Next{
		if list_a.Val != list_b.Val{
			return false 
		}
	}
	return true
}



