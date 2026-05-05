package main
import (
	"fmt"
	
)
func main(){
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
	
}