package main

import (
	"fmt"
	"strconv"
)

func main(){

	fmt.Println("========PRIMEROS EJERCICIOS===========")

	//convertir un numero entero a decimal 
	var entero = 42 
	//convertirlo a float 
	var decimal float64 = float64(entero)
	fmt.Println("Entero:", entero)
    fmt.Println( "Decimal: ",decimal)

	fmt.Println("========SEGUNDO EJERCICIOS===========")

	nota1 := 85
	nota2 := 90
	nota3 := 78

	//quiero el promedi en float 
	//cremoa una variable con float y calculamos promedio 
	promedio := float64(nota1+nota2+nota3) / 3.0
	fmt.Println("Promedio:", promedio)


	fmt.Println("========TERCER EJERCICIOS===========")
	//Concatenar cadenas 
    nombre := "Jostyn"
	apellido := "Iraheta"
	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)

	//concatenamos mas cadenas 
	modelo := "3R"
	marca := "Raizen"
	nombre_mouse := marca + " " + modelo
	fmt.Println("Informacion: ", nombre_mouse)



    fmt.Println("========CUARTO STRCONV===========")
	texto := "21"
	numero1, err := strconv.Atoi(texto) //devuelve int y error 

	if err != nil {
		fmt.Println("Hubo un error", err)

	}else {
		fmt.Println("Conversion exitosa", numero1)
	}

	//ITOA 
	//integer ASCI
	numero2 := 21
	texto2 := strconv.Itoa(numero2) //int - string
	fmt.Println(texto2)
	//es la forma mas directa de transformar un numero entero en texto 

	//NIl nulo o vacio 
	
	number, err := strconv.Atoi("21")
	if err == nil {
		fmt.Println("Conversion correcta", number)
	}else{
		fmt.Println("Error", err)
	}

	
	


}