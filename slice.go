package main

import (
	"fmt"
	"math"
)

//实现删除切片特定下标元素的方法。
//
//要求一：能够实现删除操作就可以。
//要求二：考虑使用比较高性能的实现。
//要求三：改造为泛型方法
//要求四：支持缩容，并旦设计缩容机制。

// 要求一：能够实现删除操作就可以。
func delete(idx int, vals []int) []int {
	if idx < 0 || idx > len(vals)-1 {
		panic("Illegal idx")
	}
	for i := idx; i < len(vals)-1; i++ {
		if i > 0 {
			vals[i] = vals[i+1]
		}
	}
	vals[len(vals)-1] = math.MaxInt
	return vals
}

// 要求二：考虑使用比较高性能的实现。
func deleteV2(idx int, vals []int) []int {
	if idx < 0 || idx > len(vals)-1 {
		panic("Illegal idx")
	}

	final := append(vals[0:idx], vals[idx+1:]...)
	return final
}

// 要求三：改造为泛型方法
func deleteV3[T any](idx int, vals []T) []T {
	if idx < 0 || idx > len(vals)-1 {
		panic("Illegal idx")
	}
	final := append(vals[0:idx], vals[idx+1:]...)
	return final
}

// 要求四：支持缩容，并旦设计缩容机制。
func shrinking[T any](vals []T) []T {
	shrinking := make([]T, len(vals))
	fmt.Printf("赋值前：len: %v, cap: %v, vals: %v \n", len(shrinking), cap(shrinking), shrinking)
	for i := 0; i < len(vals); i++ {
		shrinking[i] = vals[i]
	}
	fmt.Printf("赋值后：len: %v, cap: %v, vals: %v \n", len(shrinking), cap(shrinking), shrinking)

	return shrinking

}
