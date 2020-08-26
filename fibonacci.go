package main

import (
	"fmt"
)

// 闭包
func fibonacci() int {
	x,y :=0,1
	return func ()int{
		x,y =y,x+y
		return x
	}
}

func main() {
	f := fibonacci()
	for i:=0;i<10;i++{
		fmt.Println(f())
	}
}