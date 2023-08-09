package main

import "fmt"

func Array() {
	// 直接初始化一个三个元素的数组，大括号里面只能少不能多
	a1 := [3]int{9, 8, 7}
	fmt.Printf("a1: %v,len: %d,cap: %d \n", a1, len(a1), cap(a1))
	// 少的部分会默认为0 等价9，8，0
	a2 := [3]int{9, 8}
	fmt.Printf("a2: %v,len: %d,cap: %d \n", a2, len(a2), cap(a2))
	//数组不支持append操作
}
