package main

import (
	"fmt"
)

//constante 

func main(){

	const (
		nombre = "Jostyn"
	)
	fmt.Println("Accedemos a valores:", nombre)

    const hola = "Cadena"
	fmt.Println(hola)
	const hola2 string = "Cadena3"
	fmt.Println(hola2)

	//variables 
	nombre2 := "Jostyn"
	apellido := "Iraheta"
	//concatenamos cadena e imprimimos 
	fmt.Println(nombre2 + " " + apellido)

	//operadores 
	servidor := true
	es_turnon := false
	fmt.Println(servidor || es_turnon)

	server_is_on := true
	fmt.Println(!server_is_on)

	number1 := 14
	number2 := 14
	fmt.Println(number1 != number2)

	//haceme una conversion de tipos 
	precio := 12 
	//lo quiero flotante 
	precio2 := float32(precio)
	fmt.Println(precio2)

	//hacemos o calculamos un promedio 
	note1 := 8
	note2 := 8
	note3 := 10
	promedio := float64(note1 + note2 + note3) / 3
	fmt.Println(promedio)

	//creame un if 
	//usemos las mismas variables 
	if promedio >= 9{
		fmt.Println("Excelente")
	}else if promedio >= 7 && promedio < 9{
		fmt.Println("Regular")
	}else {
		fmt.Println("Reprobado")
	}

	//hagamos un switch un ejemplo con las motos

	//creamos una variable de string 
	marca_celular := "Samsung"
	//ponemos switch y llamamos la varibale cuando es de tipo strig
	switch marca_celular{
	case "Samsung":
		fmt.Println("Especificaciones")
	case "Iphone":
		fmt.Println("Especificaciones")
	case "Xiaomi":
		fmt.Println("Especificaciones")
	}

	//diferente manera , siempre llamamos a la vaiable
	teclado := 1
	switch teclado{
	case 1:
		fmt.Println("Mecanico")

	case 2:
		fmt.Println("Convencional")
	}

	




	marca_moto := "Yamaha"
	switch marca_moto {
	case "Honda":
		fmt.Println("Honda: 150cc........")
	case "Yamaha":
		fmt.Println("Yamaha: 300cc R3")
	default:
		fmt.Println("Marca desconocida")
	}











	



}