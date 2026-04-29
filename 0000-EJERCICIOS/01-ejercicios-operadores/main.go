package main

import (
	"fmt"
)

func main(){

	//Declaramos variables 
	fmt.Println("=======EJERCICIO 1=======")
	ram_total := 16
	ram_usada := 10
	cpu_totales := 4
	cpu_usos := 3
	esta_activo := true
	tiene_backup := false

	fmt.Println("Ram disponible: ", ram_total - ram_usada)
	fmt.Println("Ram usada mayor al 50%?", ram_usada > (ram_total / 2))
	fmt.Println("Esta activo y tiene backup: ", esta_activo && tiene_backup)
	fmt.Println("Esta activo y tiene backup: ", esta_activo || tiene_backup)
	fmt.Println("Cpus disponibles: ", cpu_totales - cpu_usos)

	fmt.Println("=======EJERCICIO 2=======")
	const nombre = "Charizard"
	hp_actual := 150
	da_racibido := 45
	
	ataque := 80
	defensa_rival := 30
	es_legendario := false
	esta_vivo := true

	fmt.Println(nombre)
	fmt.Println("da real al rival: ", ataque - defensa_rival)
	fmt.Println("Hp despues de recibir un golpe: ", hp_actual - da_racibido)
	fmt.Println("Hp acutal mayor a 100: ", (hp_actual - da_racibido) > 100)
	fmt.Println("Es legendario y esta vivo?: ", es_legendario && esta_vivo)
	fmt.Println("Es legendario o esta vivo?", es_legendario || esta_vivo)

	//Ejercicios 
	fmt.Println("Operadores aritmeticos")  //  + , - , * , / , % 
	variable1 := 10
	variable2 := 10
	fmt.Println(variable1 + variable2)

	fmt.Println("Operadores de comparacion") // >, <, >= , <= , != , == , 
	variable3 := 50 
	variable4 := 50 
	fmt.Println(variable3 >= variable4)

	fmt.Println("Operadores condicionales ") // ! , && , ||
	variable5 := true
	variable6 := true
	fmt.Println(variable6 && variable5)

	fmt.Println("Operadores de asignacioon")//asigale el valor que te digo
	// += , -= , *= , /= 
	temperatura := 100

	temperatura *= 2
	fmt.Println(temperatura)

	fmt.Println("Operadores de incremento y decremento") //variable-- , variable++ , aumenta 1 o mas 
	conexiones := 10 
	conexiones--
	fmt.Println(conexiones)


















}
