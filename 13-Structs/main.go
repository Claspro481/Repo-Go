package main

import(
	"fmt"
)

//declaracion + instancia por separado
type Auto struct{
	Modelo string
	Marca string
}

type Pokemon struct{
	Nombre string
	Tipo string
	Nivel int
}

func main(){

	pikachu := Pokemon{
		Nombre: "Pikachu",
		Tipo: "Electrico",
		Nivel: 22,
	}
	fmt.Println(pikachu)

   
	var miAuto Auto
	miAuto.Modelo = "Supra"
	miAuto.Marca = "Toyota"
	fmt.Println("Marca:", miAuto.Marca, miAuto.Modelo)

	

	












/*	type NombreDelStruct struct{
		Campo1 Tipo1
		Campo2 Tipo2
	}*/


  



}