package main

import (
	"unicode/utf8"
)

func string() {
	println("Hello,\" GO!")
	println(len("你好"))
	println(utf8.RuneCountInString("你好"))
	//string长度很特殊：字节长度和编码无关，用len(str)获取
	//字符数量和编码有关，用编码库来计算，默认情况下使用utf8库

}
