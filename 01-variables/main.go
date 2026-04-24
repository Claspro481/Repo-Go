package main

import (
	"fmt"
)

func main()  {

	//var basico -go infiere que es string 
	var modelo = "Toyota Supra"
	fmt.Println(modelo)

	//var 
	var teclado = "Regradon"
	fmt.Println(teclado)
	var variable3 = 21
	fmt.Println(variable3)



	//Declarar varias del mismo tipo 
	var velocidad, potencia int = 280, 320
	fmt.Println(velocidad, potencia)
	//variables del mismo tipo 
	var torque, velocidad2 int = 120, 300
	fmt.Println(torque, velocidad2)
	var precio, gramos float32 = 3 , 120.3
	fmt.Println(precio , gramos)





	// Go infiere el tipo automatico 
	var tuboActivo = true
	fmt.Println(tuboActivo)



	

	//Forma corta, solo dentro de funciones 
	pokemon := "Chizard"
	nivel := 78
	hp := 266.5
	esFuego := true
	fmt.Println(pokemon, nivel, hp, esFuego)


    //go por defecto le da el valor de 0 y si es bool le da el que el quiere osea false 
	var kilometro int
	// Si no le das el valor, go le pone el valor cero del tipo 
	var kilometraje int
	var marca string
	var encendido bool
	var on bool

	fmt.Println(kilometraje)
	fmt.Println(marca)
	fmt.Println(encendido)
	fmt.Println(kilometro)
	fmt.Println(on)







}