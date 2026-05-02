package main

import (
	"fmt"
	"strconv"
)

func main(){

	//Primer EJERCICIO - convertimos numero entero a texto 
	edad := 21
	texto := strconv.Itoa(edad)
	fmt.Println("Edad como texto:", texto)

	nombre := "Jostyn"
	apellido := "Iraheta"
	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)

	var nota1 int = 5
	var nota2 int = 10
	var nota3 int = 7

	//covertimos a decimales = variable - tipo de dato- operacion 
	promedio := float64(nota1+nota2+nota3) / 3
	fmt.Println(promedio)

	var edad2 int = 21
	new_age := float32(edad2)
	fmt.Println(new_age)
	


	




}
