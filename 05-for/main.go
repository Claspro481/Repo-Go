//for = para 
//estructura simple  for inicio ; condicion , incremento
package main
import (
	"fmt"
)

func checkConnection() bool{
	return false
}

func main(){
	
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++ // incrementamos i
	}
	fmt.Println("-------------------")

	
	for numero := 0; numero < 3; numero++{
		fmt.Println(numero)
	}
	fmt.Println("-------------------")

	for rango := range 3 {
		fmt.Println("rango:", rango)
	}

	for {
		fmt.Println("loop")
		break
	}

	for valor := range 6 {
		if valor%2 == 0 {
			continue
		}
		fmt.Println(valor)
	}

	fmt.Println("----------EJEMPLOS")
	//imaginemos que estamos monitoreando un servidor 
	// y queremos intentar reconectarte mientras el servidro no responde
	//for1 , si es menor a lo que vale recorre lo que falta 
	connected := false
	attempts := 0

	for !connected && attempts < 2 {
		fmt.Println("Intentando conectar al servidor....")
		attempts++
	}
	fmt.Println("Conexion establecida o limite alcanzado")

	

	//creamos la variable desde adentro
	for a := 1; a < 10; a++{
		fmt.Println(a)
	}

	for intentoConexiones := 1; intentoConexiones <= 3; intentoConexiones++ {
		fmt.Println("Intento de conexion numero: ", intentoConexiones)
	}

	for archivoProcesado := 1; archivoProcesado <= 5; archivoProcesado++ {
		fmt.Println("Archivo procesado: ", archivoProcesado)
	}

	// for que usa rande para recorrer listas 
	nombre2 := []string{"Juan", "Maria", "Pedro"}

	for indice, nombre2 := range nombre2 {
		fmt.Println("Indice", indice, "Nombres:", nombre2)

	}

	celulares := []string{"Samsunf", "Apple", "Xiaomi"}
	for _, celulares := range celulares {
		fmt.Println("Celulares:", celulares)
	}

	//El otro tipo de for inifinito
	intentos := 0
	for{

	
		fmt.Println("Esperando conexion entrante...")
		intentos++

		if checkConnection(){
			fmt.Println("Conexion establecida")
			break
		}
	}




	


}