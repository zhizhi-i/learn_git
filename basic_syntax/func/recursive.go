package main

// 递归使用不当，就有可能造成stackoverflow
// goroutine 为一个2KB的堆栈
// 指标：栈溢出可以考虑增加堆栈的大小
// 治本：找到对应递归的代码进行修改
func Recursive(n int) {
	if n > 10 {
		return
	}
	Recursive(n + 1)
}

func A() {
	B()
}

func B() {
	C()
}

func C() {
	A()
}
