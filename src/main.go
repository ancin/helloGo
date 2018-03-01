package main

import "fmt"
import "time"
import "log"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func sum(a, b int) int {
	fmt.Printf("sum()函数中 a= %d \n", a)
	fmt.Printf("sum()函数中b=%d \n", b)
	return a + b
}
func swap(x, y string) (string, string) {
	return y, x
}
func swap2(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}

func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

type Book struct {
	title   string
	author  string
	subject string
	id      int
}

type MyStr struct {
	name string
}

func main() {
	var start = time.Now()
	var str string = "HELLO world GO."
	var i, j int
	var arry [10]int
	fmt.Println(str)
	fmt.Println(max(11, 99))

	for c := 0; c < 6; c++ {
		fmt.Println(c)
	}
	for i = 0; i < 10; i++ {
		arry[i] = i + 100
	}
	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, arry[j])
	}

	x, y := swap("xyz", "abc")
	fmt.Println(x, y)
	//go say("world")
	//say("hello")
	var a int = 10
	var b int = 20
	var ip *int
	var ptr *int

	ip = &a
	fmt.Println(sum(a, b))
	fmt.Printf("变量a的地址: %x\n", &a)
	fmt.Printf("变量b的地址: %x\n", &b)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	if ptr == nil {
		fmt.Printf("空指针 %x \n", ptr)
	}
	var book Book
	book.subject = "bk_subject"
	book.author = "LeeChao"
	book.title = "Redis development"
	fmt.Println(book)
	slicetest()
	rangeTest()
	maptest()
	fmt.Println(time.Now().Unix())
	var fact int = 15
	fmt.Printf("%d 的阶乘是 %d\n", fact, Factorial(uint64(fact)))
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}

	//http server
	//runServer()
	//
	//type_name(expression)
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()
	//writetest()
	readFile()
	outputerror()
	Tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}
	Tom.SayHi()
	Tom.Sing("very good things")
	//runClient()
	myStr1 := new(MyStr)
	myStr2 := new(MyStr)
	(*myStr1).name = "1111"
	(*myStr2).name = "222222"
	var lst List
	var plst *List = &lst
	plst.Append(myStr1)
	plst.Append(myStr2)
	fmt.Printf("##List size= %d \n", plst.GetSize())

	opmysql()

	fmt.Println("执行时间:", time.Since(start))
	log.Println(" main run end.")

}
