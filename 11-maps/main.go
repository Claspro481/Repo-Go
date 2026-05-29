package main

import (
	"fmt"
)

func main(){

	//mapas 
	fmt.Println("====MAPS 1")

	usuarios := map[string]string{
		"Alice" : "Admin",
		"Bob" : "Editor",
		"Chats" : "Viewer",
	}
	//Agregar un nuevo usuario 
	usuarios["Diana"] = "Admin"

	fmt.Println("Usuarios:", usuarios)
	//accedemos a un valor 
	fmt.Println("Rol de bob: ", usuarios["Bob"])
	
	//tambien podemos eliminar 
	delete(usuarios, "Chats")

	

	fmt.Println("====MAPS 2")

	//aqui hacemos otro map 
	//la sintaxis es de variable , asignacion , map,[tipo]tipo
	//nombre , valor 
	motos := map[string]int{
		"Vstrom" : 250,
		"Cb" : 190,
		"Shine sp": 125,
	}
	fmt.Println("Que cilindrada es la vtrom?:", motos["Vstrom"])
	//varibale del mapa y valor 



	


	fmt.Println("====MAPS 3")
	servidores := map[string]string{
		"Server-01": "192.168.1.10",
		"Server-02": "192.168.1.11",
	}
	fmt.Println("Ip de server-01:", servidores["Server-01"])
	//podemos agregar un nuevo servidor
	servidores["Server-03"] = "192.168.1.12"
	fmt.Println(servidores)





	fmt.Println("====MAPS 4")
	pokemones := map[string]string{
		"Pikachu" : "Electrico",
		"Charmander" : "Fuego",
		"Bulbasaur" : "Platanta",
	}
	fmt.Println("Que tipo es pikachu?", pokemones["Pikachu"])
	//podemos eliminar un pokemon  
	delete(pokemones, "Charmander")
	fmt.Println(pokemones)




	

	









}