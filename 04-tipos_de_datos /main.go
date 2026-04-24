package main

import (
	"fmt"
	
)

func main(){

	//Uso real int numero enteros
	fmt.Println("------------NUMERO ENTERO-------")
	puerto := 8080
	cpus := 4
	peticiones := 1500
	tiempoRespuesta := 200 //milisegundos

	fmt.Println("Datos:", puerto, cpus, peticiones, tiempoRespuesta)

	//Float64 con numeros decimales 
	fmt.Println("------------NUMEROS DECIMALES-------")
	temperatura := 75.5 
	porcentajeCPU := 23.8
	memoriUsada := 3.2 // GB
	latencia := 0.5

	fmt.Println("Datos:", temperatura, porcentajeCPU, memoriUsada, latencia)


	//STRING TEXTO
	//uso real 
	fmt.Println("------------STRINGS-------")
    servidor := "web-server-01"
	ip := "192.168.1.10"
	sistema := "Ubuntu 22.04"
	estado := "corriendo"

	fmt.Println("Informacion del servidor: ", servidor, ip, sistema, estado)

	//bool verdadero o falso 
	//uso real 
	fmt.Println("------------BOOL-------")
	estaActivo := true
	tieneSsh := true
	necesitaReinicio := false
	estaEnProduccion := true

	fmt.Println("Informacion de sistema: ", estaActivo, tieneSsh, necesitaReinicio, estaEnProduccion)

	//Byte un solo caracter 
	//Uso real - leer archivos , procesar texto - las letras valen ciertos bytes 
	fmt.Println("----------BYTE un solo caracter--------")
	inicial := byte('A')
	grado := byte('B')
	modelo := byte('S')
	fmt.Println(inicial, grado, modelo)
	//ejemplo  
	contenido := []byte("Hola")
	fmt.Println("Valor de contenido: ", contenido)

     
	fmt.Println("-----ENTEROS SiN NEGATIVOS--") //algo que nunca puede ser negativo 
	puertoSevidor := uint(8080)
	maximoUsuarios := uint(1000)
	bytesRecibidos := uint(5120)
	fmt.Println(puertoSevidor, maximoUsuarios, bytesRecibidos)





    fmt.Println("------------------------------------------")
	//Diferentes comparaciones , tama;os 
	var conexiones int32 = 200000

	var bytesTotal int64 = 90000000000
	fmt.Println(conexiones, bytesTotal)

	fmt.Println("------------------------------------------")

	//diferencia de float 
	var temperatura2 float32 = 75.5

	var latencia2 float64 = 0.00000212444
	fmt.Println(temperatura2, latencia2)

	//ruine caracteres unicode , simbolos unicodes os





}