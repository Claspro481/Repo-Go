package main

import "fmt"

func main() {
	fmt.Println("=== Operadores aritméticos ===")

	//Operadores aritmeticos 
	cpu1 := 10
	cpu2 := 5
	//Imprimos resutados 
	fmt.Println("Suma: ",cpu1 + cpu2)
	fmt.Println("Resta: ", cpu1 - cpu2)
	fmt.Println("Multiplicacion: ", cpu1 * cpu2)
	fmt.Println("Division: ", cpu1 / cpu2)
	fmt.Println("Residuo: ", cpu1 % cpu2)


	fmt.Println("=== Operadores Comparacion ===")
	//Verdadero o falso 
	ram1 := 8
	ram2 := 16
	fmt.Println("Igual:", ram1 == ram2) 
	fmt.Println("Diferente:", ram1 != ram2) 
	fmt.Println("Mayor: ", ram1 > ram2)
	fmt.Println("Menor: ", ram1 < ram2)
	fmt.Println("Mayor o igual: ", ram1 >= ram2)
	fmt.Println("Menor o igual: ", ram1 <= ram2)

	fmt.Println("=== Operadores Condicionales ===")

	estaActivo := true
	tieneSsh := false

	fmt.Println(estaActivo && tieneSsh) // mentira porque los dos son diferentes 
	fmt.Println(estaActivo || tieneSsh) // uno basta 
	fmt.Println(!tieneSsh) //Invierte al valor contrario


	fmt.Println("=== Operadores Asigancion  ===")

	temperatura := 50

	//temperatura += 5 
	//temperatura -=10
	//temperatura *= 2
	//temperatura /= 2
	fmt.Println(temperatura)

	fmt.Println("=== Operadores incremento y decremento ===")

	conexiones := 1
	//conexiones++   //suma 1
	//conexiones++ //suma otro mas 
	conexiones--

	fmt.Println("Conexiones: ",conexiones)


	//OPERADOR ||
	servidorA := false
	servidorB := false
	fmt.Println(servidorA || servidorB) //con que funcione 1 es suficiente para que nos lance true 
	//si los dos son false entonces flase 

}
