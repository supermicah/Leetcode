package main

import "fmt"

func main() {
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	println(fmt.Sprintf("%+v", a[len(a)-1]))
	println(fmt.Sprintf("%+v", a[:len(a)-1]))
}
