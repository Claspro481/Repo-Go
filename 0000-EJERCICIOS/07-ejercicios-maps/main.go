package main

import(
	"fmt"
)

func main(){

	//crea un mapa de usuarios y roles y valor sea su rol (
	// Admin, editor , viewer)
	//agrega un nuevo usuario 
	//elimina uno existente 

	user := map[string]string{
		"User1" : "Admin",
		"User2" : "Editor",
		"User3" : "Viewer",
	}
	//agregamos 
	//mapa , clave valor 
	user["User4"] = "Tech"
	//eliminamos , delete , 
	delete(user, "User3")
	fmt.Println("Usuarios:", user)

	//servidores e ips 
	//define un mapa , agregar servidor, mostrar servidores activos 
	servidor := map[string]string{
		"Server-1" : "192.168.1.10",
		"Server-2" : "192.168.1.11",

	}
	//aghregamos uno 
	servidor["Server-3"] = "192.168.1.12"
	fmt.Println("Seridores:", servidor)
	marcas := map[string]int{
		"HP" : 21,
		"DELL" : 30,
		"APPLE" : 50,
	}
	marcas["AZUS"] = 30
	fmt.Println(marcas)



	

	




	
}