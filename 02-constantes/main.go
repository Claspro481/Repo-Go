package main
import (
	"fmt"
	
)

//A dioferencia de las variables, las constantes no cambian nunca . el compilador te impide modificarlas 
func main(){
	
	//constante con const en vez de var 
	fmt.Println("-----------------Contantes Normales--------")
	const velocidadLuz = 29999999
	fmt.Println("Velicidad de la luz: ", velocidadLuz)
	const constante1 = "constant1"
	fmt.Println("Constante normal: ", constante1)
	const contante2 = "Constante 3"
	fmt.Println("Constante 3: ", contante2)

	


    fmt.Println("-----------------Contantes TIPADAS--------")
	//Constantes tipadas nosotros elegimos el tipo int 66
	const ram int64 = 16
	fmt.Println("Ram: ", ram)
	const conexiones int32 = 200
	fmt.Println("Conexiones: ", conexiones)
	const puerto2 int32 = 5050
	fmt.Println("Puerto: ", puerto2)



	//Multiples constantes - la forma profesional 
	fmt.Println("-----------------Constantes de forma profesional --------")
	const(
		nombrePokemon = "Pikachu"
		nivel = 25
		hp = 35
		esLengendario = true

	)
	fmt.Println(nombrePokemon, nivel, hp, esLengendario)

	//Constantes en devops 
	const (
		puerto = 8080
		timeout = 30
		MaxConexiones = 100
		Ambiente = "produccion"
	)
	fmt.Println("Maximo Conexiones: ", MaxConexiones)

	//Constante de celular
	const(
		modelo = "S22 Ultra"
		procesador = "Redragon"
		ram2 = 16
		almacanmiento = 120
	)
	fmt.Println("Informacion del telefono: ", modelo, procesador, ram2, almacanmiento)

	
	



}

