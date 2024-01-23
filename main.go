package main

import "fmt"

// 实现删除切片特定下标元素的方法。
//
// 要求一：能够实现删除操作就可以。
// 要求二：考虑使用比较高性能的实现。
// 要求三：改造为泛型方法
// 要求四：支持缩容，并旦设计缩容机制。
func main() {
	println("Hello")
	vals := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v \n", vals)

	//delete(3, vals)
	//fmt.Printf("%v \n", vals)

	// 直接截取
	fmt.Printf("%v \n", deleteV2(1, vals))
	vals1 := []string{"Herry", "Jerry", "John"}

	// 泛型
	fmt.Printf("%v \n", deleteV3[string](1, vals1))
	strings := deleteV3[string](1, vals1)

	// 缩容 3->2
	fmt.Printf("缩容前：len: %v, cap: %v, vals: %v \n", len(strings), cap(strings), strings)
	strings = shrinking[string](strings)
	fmt.Printf("缩容后：len: %v, cap: %v, vals: %v \n", len(strings), cap(strings), strings)

}
