package main

import (
	"fmt"
)

// Creamos una constantes
func main() {

	const (
		nombreServidor = "servidor1"
		ip             = "192.168.1.10"
		puerto         = 29901901
	)
	fmt.Println("Datos: ", nombreServidor, ip, puerto)

	ram_disponible := 8
	esta_activo := true
	fmt.Println(ram_disponible, esta_activo)

	const (
		nombre = "Pikachu"
		tipo   = "electrico"
		hp_max = 50
	)
	fmt.Println(nombre, tipo, hp_max)

	hp_actual := 40
	esta_vivo := true
	fmt.Println(hp_actual, esta_vivo)

	const (
		marca  = "NVIDIA"
		modelo = "RTX 4090"
		v_Ram  = 16
	)
	fmt.Println(marca, modelo, v_Ram)
	temperatura_actual := 200
	estado_uso := true
	fps_actual := 300

	fmt.Println(temperatura_actual, estado_uso, fps_actual)

    
	fmt.Println("------EJERCICIO DE TIPOS DE DATOS")

	const (
		ip_puerto = "8080"
		puerto2 = 8080

	)
	latencia_ms := 200.0 //milisegundos
	esta_conectado := true
	bytes_recibidos := uint(600)
	proces_name := "Conexion"

	fmt.Println("Info puertos:", ip_puerto, puerto2)
	fmt.Println(latencia_ms,esta_conectado,bytes_recibidos,proces_name)

	//en la vida real la latencia tiene decimales 

	//CREAMOS UN POKEMON POKEDEK 
	const(
		numero_pokedex = uint(1)
		name = "Pikachu"
		tipo1 = "Electrico"
		peso = 180.0



	)
	fmt.Println(numero_pokedex, name, tipo1, peso)
	hp := 200
	nivel := uint(9)
	experiencia := 100.5
	en_equipo := true
	fmt.Println(hp,nivel,experiencia,en_equipo)


	

}
