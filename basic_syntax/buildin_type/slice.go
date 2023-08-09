package main

import "fmt"

func Slice() {
	s1 := []int{1, 2, 3, 4} //直接初始化了4个元素切片
	fmt.Printf("s1: %v, len=%d, cap=%d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) //通过make初始化了3个元素，容量为4的切片
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 7) //追加一个元素，没有扩容
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 8) //追加一个元素,扩容了
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 4) // make只传入一个参数，表示创建一个4个元素的切片
	fmt.Printf("s3: %v, len=%d, cap=%d \n", s3, len(s3), cap(s3))
}

func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[2:3] //[start下标:end下标],
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	//容量等于start开始往后，包括原本s1的底层数组的元素个数
	//左闭右开 https://blog.csdn.net/huluhulu957/article/details/131487441
	s3 := s1[2:]
	fmt.Printf("s3: %v, len=%d, cap=%d \n", s3, len(s3), cap(s3))
	s4 := s1[:4]
	fmt.Printf("s4: %v, len=%d, cap=%d \n", s4, len(s4), cap(s4))
}

func ShareSlice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[2:]
	//切片和子切片内存共享
	//一旦产生扩容操作，则底层数组变为2个不同的，即不共享内存
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	s2[0] = 99
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v, len=%d, cap=%d \n", s1, len(s1), cap(s1))

	s2 = append(s2, 199)
	fmt.Printf("s2: %v, len=%d, cap=%d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v, len=%d, cap=%d \n", s1, len(s1), cap(s1))
}
