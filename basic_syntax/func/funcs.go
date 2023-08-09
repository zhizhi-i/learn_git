package main

//四个部分：关键字func，方法名称首字母大小写决定作用域 参数列表[<name type>] 返回列表[<name type>]

//名字+参数列表+返回值 也叫做方法签名

// Func1 没有任何参数
func Func1() {

}

// Func2 只有一个参数
func Func2(a int) {

}

// Func3 多个参数
func Func3(a int, b string) {

}

// Func4 多个参数一个类型
func Func4(a, b string) {

}

// Func5 多个返回值
func Func5(a, b string) string {
	//有返回值要保证一定返回
	return "hello"
}

// Func6 多个返回值
func Func6(a, b string) (string, string) {
	//有返回值要保证一定返回
	return "hello", "world"
}

func Func7(a, b string) (name string, age int) {
	return "hello", 18
}

func Func8(a, b string) (name string, age int) {
	name = "hello"
	age = 18
	return
}

func Func9() (name string, age int) {
	// 等价于 "",0
	// 使用对应数据类型的默认值
	return
}
