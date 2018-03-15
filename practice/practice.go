package practice

import (
	"fmt"
	"strconv"
)

/** 实现一个栈，push, pop方法，后进先出 **/
type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

/** 实现一个栈，push, pop方法，后进先出 **/

//计算一个slice的平均值
func Avarage(arr []float64) (result float64) {
	sum := 0.0
	switch len(arr) {
	case 0:
		result = 0
	default:
		for _, v := range arr {
			sum += v
		}
		result = sum / float64(len(arr))
	}
	return result
}

//返回两个参数的正确顺序
func Order(x, y int) (int, int) {
	switch {
	case x > y:
		return y, x
	default:
		return x, y
	}
}

//编写函数接受不定整数参，每行打印一个
func randArgs(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

//菲薄拉起数列
func fabnaq(num int) (result []int) {
	result = make([]int, num)
	if num == 1 {
		result[0] = 1
		return
	}
	result[0], result[1] = 1, 1
	for i := 2; i < num; i++ {
		result[i] = result[i-2] + result[i-1]
	}
	return
}

//定义一个函数，接受一个函数和不定长度的int做参数
//对每一个int执行参数中的函数，返回计算结果的列表
func Map(f func(int) int, args ...int) []int {
	result := make([]int, len(args))
	for k, v := range args {
		result[k] = f(v)
	}
	return result
}

//计算slice最小值
func min(s []int) (min int) {
	min = s[0]
	for _, v := range s {
		if min > v {
			min = v
		}
	}
	return min
}

//冒泡排序
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

//函数返回函数
func PlusTwo() func(int) int {
	return func(x int) int {
		return x + 2
	}
}
