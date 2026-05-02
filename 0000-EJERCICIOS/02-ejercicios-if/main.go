package main
import(
	"fmt"
)
func main(){
	//dado un numero escriba si es par o impar 
	/*numero := 2
	if numero % 2 == 0{
		fmt.Println("Es par")
	}else {
		fmt.Println("Es impar")
	}*/

	//dado un numero de temperatura 
	/*temperatura := 33
	if temperatura < 15 {
		fmt.Println("Frio")
	}else if temperatura >= 15{
		fmt.Println("Templado")
	}else if temperatura < 28{
		fmt.Println("Calor")
	}*/

	contra := "123"
	activo := true

	if contra == "123" && activo {
		fmt.Println("Puede pasar")
	}else {
		fmt.Println("Acceso denegado")
	}


	edad := 18

	if edad <= 12 {
		fmt.Println("Es kid")
	}else if edad >=  17{
		fmt.Println("Adolescente")
	}else if edad >= 18 {
		fmt.Println("Adulto")
	}else if edad >= 65{
		fmt.Println("Adulto mayor")
	}else{
		fmt.Println("Edad invalida")
	}


	




	









}