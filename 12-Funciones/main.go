package main

import(
	"fmt"
)

//AQUI VAN LAS FUNCIONES QUE CREAREMOS 

//FUNCION #1 
//Funcion que no recibe parametros ni devuelve valores 
func saludar(){
	fmt.Println("Hola Coders!!")
}

// Funcion #2
//Funcion con parametros 
//ADENTRO DE LOS () VAN LOS PARAMETROS 
func mostrarAuto(marca string, modelo string){
	fmt.Println("Marca:", marca)
	fmt.Println("Modelo:", modelo)
}
//funcion #3 
//funcion que devuelve un valor 
func calcularDanio(ataque int, bono int) int{

	return ataque + bono
}

//FUncion #4 
//funcion que devuelve multiples retornos 
func datosAuto()(string, int) {
	marca2 := "Ford"
	anio := 2026

	//luego le pedimos que retorne 
	return marca2, anio
}


func main (){

   fmt.Println("===RESULTADO DE FUNCIONES")
   //Imprimimos la funcion #1 o mas bien la ejecutamos 
   saludar()

   fmt.Println("==FUNCION CON PARAMETROS")
   mostrarAuto("Toyota", "Corola")

   fmt.Println("==FUNCION DEVUELVE UN VALOR")
   danioTotal := calcularDanio(80, 20)
   fmt.Println("Danio total:", danioTotal)

   fmt.Println("==FUNCION CON MUTIPLES RETORNOS")
   marca2, anio := datosAuto()

   fmt.Println("Marca", marca2)
   fmt.Println("Anio:", anio)
   












}