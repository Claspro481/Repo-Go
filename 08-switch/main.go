package main

import (
	"fmt"
	"time"
)

//hacemos una funcion para agregar con switches 
func agregarUsuario(){
	fmt.Println("Usuario agregado")
}

func eliminarUsuario(){
	fmt.Println("Usuario eliminado")
}

func obtenerUsuario(){
	fmt.Println("Usuario obtenido")
}
func crearUsuario(){
	fmt.Println("Usuario creado")
}
func main(){

	//agregamos switch de metodos
	metodo := "POST"
	switch metodo {
	case "GET":
		obtenerUsuario()
	case "POST":
		crearUsuario()
	}

	opcion := 3
	switch opcion{
	case 1:
		agregarUsuario()
	case 2:
		eliminarUsuario()
	default:
		fmt.Println("Opcion invalida")
	}



    //imaginemos que queremos imprimir el nombre de un dia segun un numero 
	diaDeLaSemana := 7
	switch diaDeLaSemana{
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Numero invalido")
	}

	temperatura := 14
	switch{
	case temperatura < 15 :
		fmt.Println("Hace frio")
	case temperatura >= 15 && temperatura < 25:
		fmt.Println("Clima agradable")
	case temperatura >= 25:
		fmt.Println("Hace calor")
	}

	//usar sentencia switch para determinar el tipo de vehiculo segun un numero ingresado por el usuario 
	auto := 1

	switch auto{
	case 1:
		fmt.Println("Automovil")
	case 2:
		fmt.Println("Motocicleta")
	case 3:
		fmt.Println("Camion")
	default:
		fmt.Println("Veihiculo desconocido")
	}

	velocidad := 50
	switch{
	case velocidad == 0:
		fmt.Println("Vehiculo detenido")
	case velocidad > 0 && velocidad <= 50:
		fmt.Println("Vehiculo en movimiento (velocidad normal)")
	case velocidad > 51:
		fmt.Println("Alta velocidad")
	}


	switch time.Now().Weekday(){
	case time.Monday:
		fmt.Println("Inicio de semana")
	case time.Friday:
		fmt.Println("Ultimo dia laboral")
	case time.Saturday, time.Sunday:
		fmt.Println("Fin de semana")
	default:
		fmt.Println("Dia normal")


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

	//nuevos ejemplos de switch 



	


	
}