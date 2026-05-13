// O pacote main define o ponto de entrada para o executável.
package main

// Importamos fmt para realizar operações de entrada e saída formatada.
import "fmt"

func main() {

	// Strings podem ser concatenadas usando o operador '+'.
	// Go suporta strings UTF-8 por padrão.
	fmt.Println("go" + "lang")

	// Inteiros e aritmética básica.
	// O Go suporta os operadores matemáticos tradicionais: +, -, *, /.
	fmt.Println("1+1 =", 1+1)

	// Números de ponto flutuante (floats).
	// Em Go, se você incluir um ponto decimal, o número é tratado como float.
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleanos e operadores lógicos.
	// '&&' representa o operador lógico E (AND).
	fmt.Println(true && false)

	// '||' representa o operador lógico OU (OR).
	fmt.Println(true || false)

	// '!' representa o operador lógico NÃO (NOT), que inverte o valor booleano.
	fmt.Println(!true)
}
