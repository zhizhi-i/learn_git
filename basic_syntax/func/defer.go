package main

import "fmt"

func DeferClosure() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("i的地址是 %p,值是 %d \n", &i, i)
		}()
	}
}
