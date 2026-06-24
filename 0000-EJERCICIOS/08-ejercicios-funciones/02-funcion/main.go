package main

import(
	"fmt"
)


func plus(a int, b int) int{
	return a + b
}

func plusPlus(a, b, c int) int{
	return a + b + c
}




func main(){

	fmt.Println("Suma:", plusPlus(10, 10 ,10))


	fmt.Println(plus(10, 10))


	fmt.Println()
}