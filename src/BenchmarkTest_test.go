package main_test

import "testing"
import "fmt"

func BenchmarkHello(b *testing.B) {
	fmt.Println("性能测试")
}
