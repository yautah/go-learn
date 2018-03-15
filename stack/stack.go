package stack

import (
	"strconv"
)

/** 实现一个栈，push, pop方法，后进先出 **/
type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(k int) {
	if s.i > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *Stack) Pop() int {
	s.i--
	return s.data[s.i]
}

func (s Stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
