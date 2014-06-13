package stack

type Stack struct {
    leng int
    data [10]int
}

func (stack *Stack) Push (ele int) {
    stack.data[stack.leng] = ele
    stack.leng++
}

func (stack *Stack) Pop() int {
    stack.leng--
    return stack.data[stack.leng]
}


