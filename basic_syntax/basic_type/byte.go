package main

func ByteExample() {
	var a byte = 'a'
	// 输出的是a的ASCII表达 97
	println(a)

	var str = "this is string"
	var bs []byte = []byte(str)
	bs[0] = 'T'
	println(str)
	println(bs)

}
