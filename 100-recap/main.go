package main
import (
	"fmt"
	"slices"
	
)

//cremoas un tipo proio basado en sting
type IDContenedor string

func main(){
	fmt.Println("=================================RECAP 1 variables")
	//datos que cambian 
	var velocidad_actual float64 = 100.40
	fmt.Println("Velocidad: ", velocidad_actual)
	//variables datos que cambian 
	var nivel_gasolina float64 =  30.30
	fmt.Println("Porcentaje de gasolina: ", nivel_gasolina, "%")
	var marcha_actual int = 2
	fmt.Println("Marcha actual: ", marcha_actual)

	var variable1 float64 = 12.2
	fmt.Println("variable: ", variable1)


	fmt.Println("=================================RECAP 2 constantes")
	const(
		numero_chasis = "MDSEW13"
		numero_placas = "M123P"
	)
	fmt.Println("CHASIS: ", numero_chasis, "PLACA:", numero_placas )
	//difeennte forma de hacer constantes  
	const limite_RPM float64 = 6500
	fmt.Println("Limite de revoluciones: ", limite_RPM)

	//hacemos otra constante
	const(
		id = 1
		name = "jos"
		age = 23
	)
	fmt.Println("Iduser:", id)

	const temp float64 = 35.5
	fmt.Println("Temperatura:", temp)

	const puerto = 8080
	fmt.Println(puerto)
	
	fmt.Println("=================================RECAP 3 conversion de tipos")
	//sensor mide con decimales 
	var rpmSensor float64 = 5400.64
	//envolvemos la variable en el tipo deseado // covertirlo a int
	rpm_pantalla := int(rpmSensor)
	fmt.Println("Rpm en entero: ", rpm_pantalla)
	//tenemos otra varable de bytes a string 
	//datos crudos que llegaron de la base de datos 
	pokemonEnBytes := []byte{80, 105, 107, 97, 99, 104, 117}
	//comvertirl en string 
	pokemon_nombre := string(pokemonEnBytes)
	fmt.Println("Nombre pokemon: ", pokemon_nombre)
	//tenemos un tipo de string 
	var miServidor IDContenedor = "srv-prod-01"
	//CONVERSION : De tipo personalizado a string standar de go 
	var textoComun string = string(miServidor)
	fmt.Println(textoComun)

	fmt.Println("=================================RECAP 3 swicht")
	//se usan para remplzar multiples if els enredados 
	rpm := 1500
	var mapaMotor string
	switch{
	case rpm < 2000:
		mapaMotor = "Modo eco"

	case rpm >= 2000 && rpm <= 6000:
		mapaMotor = "Modo confort"

	case rpm > 6000:
		mapaMotor = "Modo sport"
	}
	fmt.Println("Modo manejo: ", mapaMotor)

	fmt.Println("====SWITCH 2")
	i := 1 // depende el valor llamamos dependiendo el caso 1 2 o 3
	switch i{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	fmt.Println("====SWITCH 2")

	marca_moto := "Kawasaki"
	switch marca_moto{
	case "Yamaha":
		fmt.Println("Yamaha R3 , 300cc, ABS, 6 velocidades")
	case "Honda":
		fmt.Println("Honda, Twister, 300cc, 6 velocidades, ")
	case "Kawasaki":
		fmt.Println("Kawa , 400cc, 6 velocidades, Abs, deportiva")
	default:
		fmt.Println("Marca no encontrada")
	}



	fmt.Println("=================================RECAP 4 arreglos")
	
	var cilindradas [5]int
	cilindradas[0] = 125
	cilindradas[1] = 160
	cilindradas[2] = 200
	cilindradas[3] = 250
	cilindradas[4] = 300
	fmt.Println("Clindrada mas baja:", cilindradas[0])


	fmt.Println("====Arreglo 2")
	cantidad_facturas := [7]int{30, 40, 50, 100, 200, 300, 56 }
	fmt.Println("Cantidades iniciales facturas:",cantidad_facturas[0])

	usuarioA := []string{"Ana", "Carlos", "Maria"}
	usuarioB := []string{"Ana", "Carlos", "Maria"}
	usuarioC := []string{"Ana", "Maria", "Carlos"}

	fmt.Println(slices.Equal(usuarioA, usuarioB))
	fmt.Println(slices.Equal(usuarioA, usuarioC))

	
	fmt.Println("====Arreglo 3")
	//aqui tendremos un arreglo normal . variable cantidad tipo
	//variable , indice , asignamos , valor
	var edades[3]int
	edades[0] = 21
	edades[1] = 22
	edades[2] = 23
	fmt.Println("Edad 1:", edades[0])

	var marcas_computadoras[3]string
	marcas_computadoras[0] = "DELL"
	marcas_computadoras[1] = "HP"
	marcas_computadoras[2] = "Lenovo"
	fmt.Println(marcas_computadoras)

	fmt.Println("====Arreglo 4")

	//var -- variable -- asignamos -- cantidad -- tipo --

	var citys = [3]string{"San salvador", "Santa Tecla", "La libertad"}
	fmt.Println(citys)
	var numbers = [2]int{1,2}
	fmt.Println(numbers[0])
	var numbers2 = [2]float64{1.23,2.30}
	fmt.Println(numbers2[1])

	primos := [...]int{2 ,3, 5, 7, 11}
	fmt.Println(primos[0])

	



	





	










	

	
	








	












}