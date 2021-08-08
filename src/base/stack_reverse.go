//实现栈的所有元素翻转
package main

import (
	st "../structs/stack"
	"fmt"
)

//实现将栈顶元素移到栈底
func MoveStack(stack *st.LNode) {
	if stack.IsEmpty() {
		return
	}

	head := stack.Pop()
	if !stack.IsEmpty() {
		MoveStack(stack)
		end := stack.Pop()
		stack.Push(head)
		stack.Push(end)
	} else {
		stack.Push(head)
	}
}

//实现栈元素翻转
func ReverseStack(stack *st.LNode) {
	for !stack.Next.IsEmpty() {
		MoveStack(stack)
		stack = stack.Next
	}
}

func main() {
	stack := &st.LNode{0, nil}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	fmt.Println("Before Reverse:")
	fmt.Println("栈大小为:", stack.Size(), "栈顶元素为:", stack.Top())
	stack.PrintStack()
	fmt.Println("After Reverse")
	ReverseStack(stack)
	fmt.Println("栈大小为:", stack.Size(), "栈顶元素为:", stack.Top())
	stack.PrintStack()
}
