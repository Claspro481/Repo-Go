package inventario

import(
	"fmt"
)

func main(){

   nombre_producto := "Computadoras"
   precio := 353.90
   stock := 20

   const IVA = 13
   const descuento = 10

   precio_final := precio + IVA
   pprecio_descuento := precio -  descuento

   //crea un menu de acciones 
   switch 3{
   case 1:
      fmt.Println("Mostrar inventario")
   case 2:
      fmt.Println("Calcular precio")
   case 3:
      fmt.Println("Calcular precio con descuento")
   }

   //slices 
   lista_productos := [3]string{"Computadora", ""}
   fmt.Println(lista_productos)
   
   




}