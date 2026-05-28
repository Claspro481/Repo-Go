package main

import (
	"fmt"
	"slices"
)

func main(){

	usuarioA := []string{"Ana", "Carlos", "Maria"}
	usuarioB := []string{"Ana", "Carlos", "Maria"}
	usuarioC := []string{"Ana","Maria", "Carlos"}

	fmt.Println("A == B:", slices.Equal(usuarioA, usuarioB)) // true 
	fmt.Println("A == C:", slices.Equal(usuarioA, usuarioC)) //falso
	//util para verificar si dos listas son completamente iguales 


	inventario := []string{"Laptop", "Mouse", "Teclado"}

	if slices.Contains(inventario, "Mou"){
		fmt.Println("El producto esta en inventario")
	}else{
		fmt.Println("Producto no encontrado")
	}
	//evitas escribir bucles para buscar elementos en las listas 

	precios := []int{1200,500,3000,800}
	slices.Sort(precios)

	fmt.Println("Precio ordenados:", precios)


	fmt.Println("=====SLICES 3")

	pokemones := []string{"Pikachu", "Charmander"}
	pokemones = append(pokemones, "Bulsaur")
	fmt.Println(pokemones)


	// slices desde un array 
	var servidores = [3]string{"Server-01", "Server-02", "Server-03"}
	//creamos una nueva variable
	
	activos := servidores[0:2] // slices las primeras 2 posiciones dame 0 y 1
	fmt.Println(activos)

	//hacer un slice de motos , donde solo nos arroje las japonesas de la lista 
	motos := []string{"Honda", "Susuki", "Yamaha", "Aprilia", "Ducati"}
	//DAME SOLO LAS JAPONESAS 
	japonesas := motos[0:3]
	fmt.Println("Japonesas", japonesas)

	//slices con make
	computadoras := make([]string, 0, 5)
	computadoras = append(computadoras, "Dell", "Lenovo")
	fmt.Println(computadoras)

	//slices con equal 
	servidoresA := []string{"Server-01", "Server-02"}
	servidoresB := []string{"Server-01", "Server-02"}
	//servidoresC := []string{"Server-02", "Server-01"}

	fmt.Println(slices.Equal(servidoresA, servidoresB))

	

















}