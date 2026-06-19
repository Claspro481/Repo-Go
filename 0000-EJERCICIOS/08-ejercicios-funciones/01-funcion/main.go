package main

import(
	"fmt"
)

//funcion sin parametros 
func age(){
	fmt.Println("Edad")
}
//funcion que recibe parametros 

func mostarPokemon(nombre string, nivel int){
	fmt.Println("Pokemon:", nombre)
	fmt.Println("Nivel:", nivel)
}

func datosMoto(modelo string, cc int, marca string){
	fmt.Println("Modelo:", modelo)
	fmt.Println("Marca:", marca)
	fmt.Println("Cilinrada:", cc)
}

//analogia de la pizzeria 
func hacerPizza(sabor string){
	fmt.Println("Preparando pizza de: ", sabor)
}

//creamos otra funcion nueva
//unafuncionsobrememoria disponible 
func memoriaLibre(total int, usado int) int{
    return total - usado
}

//hacemo una funcion para calcular latencia
func totalLatencia(ping1 int, ping2 int) int{
	return (ping1 + ping2) / 2
}

//si queremos retornar un apellido y nombre
func datosUsuario(name string, apellido2 string) string{
	return name + apellido2
}

//pokemon ataque 
func pokemonAtaque(nombrePokemon string, ataque string) string{
	fmt.Println("Nombre del pokemon:", nombrePokemon, "Ataque:", ataque)
	return nombrePokemon + ataque

}
//funcion de velocidad 
func velocidadAuto(distancia int, tiempo_total int) int{
	return distancia / tiempo_total
}
//hacer funcion deservidor de red 
func conexionesTotales(conexiones_activas int, conexiones_permitidas int) int{
	return conexiones_permitidas - conexiones_activas
}
func main(){

	fmt.Println("Conexiones disponibles:", conexionesTotales(10, 20))

	fmt.Println("Velocidad promedio: ", velocidadAuto(120, 2), "Km por hora")

	fmt.Println(pokemonAtaque("Pikachu", "Fuego"))

	fmt.Println("Nombre completo:", datosUsuario("Jostyn", "Iraheta"))

	fmt.Println("Total latencia:", totalLatencia(120, 140), "ms")

	fmt.Println("Memoria disponible:", memoriaLibre(32, 12), "GB")


	mostarPokemon("Pikachu", 25)

	//queremos mandar datos 
	datosMoto("nx190", 190, "Honda")

	hacerPizza("Peperoni")
	

	age()

}