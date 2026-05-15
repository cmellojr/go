// Pacote main para tornar o arquivo um executável.
package main

import "fmt"

func main() {

	// 'var' declara uma ou mais variáveis.
	// Aqui o Go infere o tipo da variável 'a' como string com base no valor inicial.
	var a = "initial"
	fmt.Println(a)

	// Você pode declarar múltiplas variáveis de uma vez.
	// O tipo 'int' é especificado explicitamente para 'b' e 'c'.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go também infere o tipo booleano para a variável 'd'.
	var d = true
	fmt.Println(d)

	// Variáveis declaradas sem um valor inicial correspondente recebem o "zero value" (valor zero).
	// Para o tipo int, o valor zero é 0.
	// Outros exemplos: para string é "", para bool é false.
	var e int
	fmt.Println(e)

	// A sintaxe ':=' é uma forma curta para declarar e inicializar uma variável.
	// É equivalente a 'var f string = "apple"'.
	// Nota: Esta sintaxe só pode ser usada dentro de funções.
	f := "apple"
	fmt.Println(f)
}
