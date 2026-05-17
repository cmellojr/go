// Pacote main para tornar o arquivo um executável.
package main

import (
	"fmt"
	"math"
)

// 'const' declara um valor constante que não pode ser alterado após definido.
// Constantes podem ser declaradas em qualquer lugar onde 'var' pode.
const s string = "constant"

func main() {
	fmt.Println(s)

	// Uma instrução const pode aparecer em qualquer lugar onde uma instrução var possa.
	const n = 500000000

	// Expressões constantes realizam aritmética com precisão arbitrária.
	// Aqui, 'd' é uma constante sem tipo (untyped constant).
	const d = 3e20 / n
	fmt.Println(d)

	// Uma constante numérica não tem tipo até que lhe seja atribuído um,
	// como por uma conversão explícita.
	fmt.Println(int64(d))

	// Um número pode ter um tipo ao ser usado em um contexto que o exija,
	// como uma chamada de função. Por exemplo, math.Sin espera um float64.
	fmt.Println(math.Sin(n))
}
