package greetings

import "fmt"

//Funcion Hello que devuuelve el saludo de la persona especifica.
func Hello(name string) string {
	message := fmt.Sprintf("Hola, %v Bienvenido", name)
	return message
}
