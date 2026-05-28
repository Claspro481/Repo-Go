package main

import (
	"fmt"
<<<<<<< HEAD
=======

>>>>>>> 8800f21 (recap and new excercises with slices and new examples)
)

func main()  {

<<<<<<< HEAD
	var arreglos[5]int
	fmt.Println("Arreglo completo", arreglos)
=======
	fmt.Println("======VAR ARREGLOS==========")

	//declaracion de arreglos con var 
	var edades [3]int
	edades[0] = 25
	edades[1] = 30
	edades[2] = 35
	fmt.Println("Edades permitidas = ", edades)

	var opciones_ram [5]int
	opciones_ram[0] = 4
	opciones_ram[1] = 8
	opciones_ram[2] = 16
	opciones_ram[3] = 24
	opciones_ram[4]= 36
	fmt.Println("Ram minima para equipo: ",opciones_ram[1])


	//areglos sin var 
	temperaturaSemana := [7]float64{30.5, 31.0, 29.0, 28.7, 32.1, 33.0, 31.5}
	fmt.Println("Temperatura Lunes:", temperaturaSemana[0])
	

	fmt.Println("======CON INICIALIZACION DIRECTA==========")
	
	//ejemplo 
	var ciudades = [3]string{"San salvador", "Santa Tecla", "La libertad"}
	fmt.Println("Ciudad mas poblada: ", ciudades[0])

	//var variable = [cantidad]tipodedatos{arreglo}
	var tipos_motos = [3]string{"Sport", "Adventure", "Trabajo"}
	fmt.Println("Tu mejor tipo de moto:", tipos_motos[0])

	//tu mejor tipo de teclado
	var tipo_teclado = [2]string{"Mecanico", "Convencional"}
	fmt.Println("Tipo de teclado:", tipo_teclado[0])


	// con inferencia de longitud
	primos := [...]int{2 ,3, 5, 7, 11}
	fmt.Println(primos[0])

	







	



>>>>>>> 8800f21 (recap and new excercises with slices and new examples)
}