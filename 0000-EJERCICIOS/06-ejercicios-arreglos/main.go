package main

import (
	"fmt"
)

func main()  {

	//ejercicios 
	//como recorrer un arreglo 
	//hacemos un arreglo 

	//arrego de edades . descompisicion 
	//5 enteros = edades de diferentes personas y muestre una en cada consola 
	// var, variables, cantidad, tipo
	fmt.Println("===========EJERCICIO #1====")

	//diferentes edades 
	var edades_2 [5]int
	edades_2[0] = 18
	edades_2[1] = 21
	edades_2[2] = 25
	edades_2[3] = 28
	edades_2[4] = 90
	fmt.Println("Edades:", edades_2[0])

	var cilindrada [5]int
	cilindrada[0] = 125
	cilindrada[1] = 150
	cilindrada[2] = 190
	cilindrada[3] = 250
	cilindrada[4] = 300
	fmt.Println("Clindradas mas bajas: ", cilindrada[0])

	//con inizialiacion directa 
	var marca_motos = [4]string{"KAWASAKI", "HONDA", "DUCATI", "APRILIA"}
	fmt.Println("Marca Italiana de motos:", marca_motos[3])

	var marca_celulares = [3]string{"SAMSUNG", "APPLE", "XIAOMI"}
	fmt.Println("Marca mas vendida en estados unidos: ", marca_celulares[1])

	var salarios_ti = [3]float64{900.21, 1000.21, 1220.21}
	fmt.Println("Salario mas alto: ", salarios_ti[2])

	//arreglo inferido 
	lista1 := [...]int{1,2,3,4,5,6,7}
	fmt.Println("Arreglo inferido: ", lista1)

	nombres := [...]string{"Ana", "Carlos", "Maria"}
	// aqui go detecta automaticamente que el arreglo tiene 3 elementos 
	fmt.Println(nombres)
    

	//aqui podemos contar cuantos elementos hay 
	motosDisponibles := [...]string{"Yamaha", "Honda", "Suzuki", "Kawasaki"}
	fmt.Println("Cantidad de motos registradas: ", len(motosDisponibles))
	for indice, moto := range motosDisponibles {
		 fmt.Println(indice,moto)
	}

	pokemonesInciales := [...]string{"Bulbasaur", "Charmander", "Squirtle"}
	fmt.Println("Pokemones inciales: ", len(pokemonesInciales))
	fmt.Println("Pokemon 1:", pokemonesInciales[0])














	


}