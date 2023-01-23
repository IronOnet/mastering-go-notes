package addtwonumbers


import "testing"

type ArrayTuple struct{
	A []int 
	B []int
	Result []int 
}

// TestNewList test the ListNode initialization 
// from an array of integers
func TestNewList(t *testing.T){
	rawArray := []int{1, 2, 3} 
	expectedList := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	actualList := &ListNode{} 
	actualList.NewList(rawArray) 
	if EqualList(expectedList, actualList) != true{
		t.Errorf("Lists are not equal expected %v got %v", expectedList, actualList)
	}
}

func TestAddTwoNumbers(t *testing.T){
	arrayTuples := []ArrayTuple{
		ArrayTuple{[]int{1, 2, 3}, []int{1, 2, 3}, []int{6, 4, 2}}, 
		ArrayTuple{[]int{5, 0}, []int{0}, []int{0, 5}},
	}
	for _, array := range arrayTuples{
		listA, listB := new(ListNode), new(ListNode) 
		listA.NewList(array.A) 
		listB.NewList(array.B) 
		tempResult := AddTwoNumbers(listA, listB) 
		actualResult := tempResult.ToArray()
		if Equal(actualResult, array.Result) != true{
			t.Errorf("Error occured in addTwoNumbers test, expected %v but got %v", array.Result, tempResult)
		}
	}

}


