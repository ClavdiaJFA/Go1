//Primer ejercicio: Definición y decaración de variables
//Segundo ejercicio: Constantes
package main //definimos paquete

import "fmt" //delcaramos dependencias necesarias

/*
--Ejecución de un programa--
go build main.go -> un archivo ejecutable
go run main.go
*/

//segundo ejercicio: Constantes de tipo string (constante: no se mofdifica valor)
/*const Curso string = "Curso ṕrofesional de Go!" */

func main() {
	//fmt.Println("Hola mundo")

	//GO no permite almacenar variables a menos que se utilizada
	//variable de tipo string
	var nombre string // ""
	var edad int      // 0 es el valor por default de este tipo de variables
	var altura = 1.60 //definir un float

	nombre = "Claudia Flores"
	edad = 26

	//Declarar res variables de tipo string:

	/* var nombre, apellido, pais string */

	//Para asingar valor

	/* var nombre, apellido, pais = "Claudia", "Flores", "México"

	fmt.Println(nombre, apellido, pais) */

	//Otra manera más sencilla de hacerlo, es:
	/*

		nombre := "Claudia"
		edad := 26
		/con el := el compilador detecta en auomáico el tipo de variable, ya sea string o entero

	*/

	fmt.Println("nombre:", nombre)
	fmt.Println("edad:", edad)
	fmt.Println("estatura:", altura)
	//fmt.Println(Curso)

}
