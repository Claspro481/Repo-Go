package main

import (
	"fmt"
)

func main (){

	//if normalito con else 
	edad := 12
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	}else {
		fmt.Println("No puedes pasar")
	}

	fmt.Println("=======Varios casos=======")
	nota := 90

	if nota >= 90{
		fmt.Println("Execelente")
	}else if nota >= 70{
		fmt.Println("Aprobado")
	}else if nota >= 50 {
		fmt.Println("Regular")
	}else {
		fmt.Println("Reprobado")
	}





	






}