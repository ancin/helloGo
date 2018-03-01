package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func slicetest() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
	num2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	println(num2)
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("num2[1:4] ==", num2[1:4])
	/* 默认下限为 0*/
	fmt.Println("num2[:3] ==", num2[:3])
	/* 默认上限为 len(s)*/
	fmt.Println("num2[4:] ==", num2[4:])
	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	num3 := num2[:2]
	printSlice(num3)
	num3 = append(num3, 1)
	printSlice(num3)
	/* 同时添加多个元素 */
	fmt.Println("num3添加多个")
	num3 = append(num3, 2, 3, 4)
	printSlice(num3)
	fmt.Println("创建切片 num4 是之前切片的两倍容量")
	/* 创建切片 numbers1 是之前切片的两倍容量*/
	num4 := make([]int, len(num3), (cap(num3))*2)
	printSlice(num4)
	copy(num4, num3)
	fmt.Println("copy num3 to num4")
	fmt.Println(num4)

}
