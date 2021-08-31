package main

import "strconv"

func main() {

}

// 编写函数，返回其（两个）参数正确的（自然）数字顺序：
func Sequence(a, b int) (int, int) {
	if a >= b {
		return a, b
	}
	return b, a
}

//创建一个固定大小保存整数的栈。它无须超出限制的增长。定义 push 函数—— 将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出（LIFO） 的。
//更进一步。编写一个 String 方法将栈转化为字符串形式的表达。可以这样的 方式打印整个栈： fmt.Printf("My stack %v\n", stack) 栈可以被输出成这样的形式： [0:m] [1:l] [2:k]

type stack struct {
	node int
	data [10]int
}

func (stack *stack) Push(number int) {
	if stack.node < len(stack.data) {
		stack.data[stack.node] = number
		stack.node++
	}

}
func (stack *stack) pop() int {
	if stack.node > -1 {
		stack.node--

	}
	return stack.data[stack.node]
}

// 更进一步。编写一个 String 方法将栈转化为字符串形式的表达。可以这样的
//方式打印整个栈： fmt.Printf("My stack %v\n", stack)
//栈可以被输出成这样的形式： [0:m] [1:l] [2:k]
func (stack *stack) String() string {
	var s string
	for i := 0; i < len(stack.data); i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(stack.data[i])
		s = "[" + k + ":" + v + "]"
	}
	return s
}
