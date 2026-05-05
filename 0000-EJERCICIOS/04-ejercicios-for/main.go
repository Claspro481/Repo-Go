package main

import (
	"fmt"
)

//func checkConnection (){
//	return false
//}

func main (){
	
	fmt.Println("============FOR1")

	intentoConexion := 1
	for intentoConexion <= 3 {
		fmt.Println("Intento conexion:", intentoConexion)
		intentoConexion++
	}

	fmt.Println("============FOR2")
	for archivoProcess := 1; archivoProcess <= 5; archivoProcess++ {
		fmt.Println("Procesando archivo numero:", archivoProcess)
	}

	fmt.Println("============FOR3")
	servicios := []string{"ngix", "mysql", "redis"}
	for indiceServicio, nombreServicio := range servicios {
		fmt.Println(indiceServicio, nombreServicio)
	}







}