package main

import "fmt"

//var stack []int
//
//func Push(stack *[]int, element ...int) {
//
//	*stack = append(*stack, element...)
//
//}
//
//func Pop(stack *[]int) {
//	*stack = (*stack)[:len(*stack)-1]
//}

//func main() {
//	Push(&stack, 1)
//	Push(&stack, 2, 3, 4)
//
//	Pop(&stack)
//
//	fmt.Println("stack:", stack)
//}

//with structs

type Stack struct {
	values []int
	top    int
}

func NewStack() *Stack {
	return &Stack{
		values: make([]int, 0),
		top:    -1,
	}
}

func (st *Stack) Push(elements ...int) {

	st.top++
	st.values = append(st.values, elements...)
}

func (st *Stack) Pop() int {

	if st.IsEmpty() {
		return 0
	}

	value := st.values[st.top]

	st.values = st.values[:len(st.values)-1]

	st.top--

	return value

}

func (st *Stack) IsEmpty() bool {
	return st.top == -1
}

func main() {

	stack1 := NewStack()

	stack1.Push(1, 2, 3, 4)

	fmt.Println("stack", stack1)

	stack1.Pop()

	fmt.Println("stack", stack1)

}
