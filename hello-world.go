// Todo programa executável em Go deve começar com a declaração do pacote main.
// O pacote 'main' informa ao compilador que este arquivo deve ser compilado como um executável,
// e não como uma biblioteca compartilhada.
package main

// O comando import é usado para incluir funcionalidades de outros pacotes.
// Aqui importamos o pacote 'fmt' (format), que contém funções para formatação de texto,
// incluindo a impressão no console.
import "fmt"

// A função main é o ponto de entrada de qualquer programa executável em Go.
// Quando o programa é executado, a execução começa por esta função.
func main() {
	// fmt.Println é uma função do pacote 'fmt' que imprime uma linha de texto no console.
	// O sufixo 'ln' significa 'line', indicando que uma nova linha será adicionada ao final.
	fmt.Println("Hello World!")
}
