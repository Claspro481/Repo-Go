//este es un archivo para practicar nuestra resolucion de problemas con switch en go 

//primero importamos le paquete main para nuestra funcion principal main
package main
//luego importamos la libreria estandar de fmt 
import (
	"fmt"
)
//creamos nuestra funcion principal main 
func main(){

	//creamos el primer ejemplo de swrich 
	fmt.Println("======EJERCICIO #1=======")
	//programa que reciva el nombre de una computadora , usando switch y muestre el pais de origen
	marca_computadora := "DELL"

	switch marca_computadora{
	case "Lenovo":
		fmt.Println("Pais de origen: CHINA")
	case "DELL":
		fmt.Println("Pais De Origen: ESTADOS UNIDOS ")
	case "ASUS":
		fmt.Println("Pais de origen: TAIWAN")
	default:
		fmt.Println("Origen desconocido")
	}

	fmt.Println("======EJERCICIO #2=======")

	//creamos un programa que evalue la edad de una persona y clasificar . en este solo ponemos switch 
	edad := 12

	switch{
	case edad <= 12:
		fmt.Println("KID")
	case edad >= 13 && edad < 17:
		fmt.Println("Adolescente")
	case edad >= 18 && edad < 64:
		fmt.Println("Adulto")
	case edad >= 65:
		fmt.Println("Adulto mayor")
	default:
		fmt.Println("Edad Desconocida")
	}

	fmt.Println("======EJERCICIO #3=======")
	//creamnos un programa que reciba un numero de nivel de acceso , 123 , y muiestre los permisos 

	acceso_usuario := 10

	switch acceso_usuario{
	case 1:
		fmt.Println("Acceso 1 = R")
	case 2:
		fmt.Println("Acceso 2 = RW ")
	case 3:
		fmt.Println("Acceso: RWX")
	default:
		fmt.Println("Accesos no otorgados")
	}






}