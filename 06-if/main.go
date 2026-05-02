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

	usuario := "admin"
	logueado := true

	if usuario == "admin" && logueado {
		fmt.Println("Bienvenido admin")
	}else if logueado {
		fmt.Println("Bienvenido usuario ")
	}else {
		fmt.Println("Porfavor inicie sesion")

	}

	usuario2 := "desbloqueado"
	user_des := "desbloqueado"
	
	if usuario2 == "Bloqueado" {
		fmt.Println("Su usuario esta bloqueado")
	}else if user_des == "desbloqueado"{
		fmt.Println("Puede acceder al perfil")
	}else{
		fmt.Println("denegado")
	}

	name := "Amin"
	age := 40

	if name == "Amin"{
		fmt.Println("Hola amin")
	}else{
		fmt.Println("Hola desconocido")
	}

	if age >= 18 {
		fmt.Println("Ya puedes votar")
	}

	if 8%2 == 0{
		fmt.Println("El 8 es par")
	}else {
		fmt.Println("Es impar")
	}

	//seguimos practicando el else 
	note1 := 10
	note2 := 9
	note3 := 10
	promedio :=  float64(note1+note2+note3) / 3

	if promedio >= 9 {
		fmt.Println("Excelente", promedio)
	}else if promedio >= 7{
		fmt.Println("Regular", promedio)
	}else{
		fmt.Println("Malo:", promedio)
	}


	//operadores relacionales para ver si cierta variable cumple una condicion
	age2 := 20 
	if age2 >= 18 {
		fmt.Println("Eres mayor de edad")
	}else {
		fmt.Println("Eres menor de edad")
	}

	//operador logico 
	age3 := 25
	tieneLicencia := true

	if age3 >= 18 && tieneLicencia {
		fmt.Println("Puedes conducir")
	}else{
		fmt.Println("No puedes conducir")
	}









	






}