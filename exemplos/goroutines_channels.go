package main

import (
	"fmt"
	"time"
)

// greet imprime uma saudação em um canal. Esta função será executada
// concorrentemente como uma goroutine. O parâmetro 'done' é um canal de
// sinalização: enviar um valor para ele indica que a goroutine terminou
// seu trabalho. Isso permite que a função main espere a conclusão de forma
// segura, sem necessidade de timeouts arbitrários.
func greet(name string, done chan<- bool) {
	// A declaração 'defer' garante que, independentemente de como a função
	// termine (normalmente ou por pânico), o sinal de conclusão será enviado.
	// Isso é um idiom comum em Go para cleanup e notificação.
	defer func() { done <- true }()

	// time.Sleep simula um trabalho que leva tempo para ser concluído,
	// como uma chamada de API ou processamento de dados.
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Olá, %s!\n", name)
}

// producer envia uma sequência de números inteiros no canal 'out'.
// Quando termina, fecha o canal para sinalizar que não há mais dados.
// Fechar o canal é um padrão importante: permite que o consumidor
// itere com 'range' sem deadlock.
func producer(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out) // Fechar o canal informa ao consumidor que o envio terminou.
}

// consumer lê números inteiros do canal 'in' até que ele seja fechado.
// A cláusula 'range' sobre um canal recebe valores até o canal ser fechado.
func consumer(in <-chan int, done chan<- bool) {
	for val := range in {
		fmt.Printf("Consumido: %d\n", val)
	}
	done <- true
}

// fib envia números da sequência de Fibonacci pelo canal 'ch'.
// Ela para de enviar quando atinge o limite 'n'.
// Este exemplo demonstra um canal unbuffered (sem buffer): cada envio
// bloqueia até que o receptor esteja pronto para receber, criando
// uma sincronização natural entre as goroutines.
func fib(ch chan<- int, n int) {
	a, b := 0, 1
	for range n {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	// =========================================================================
	// 1. Goroutine básica com canal de sincronização (unbuffered)
	// =========================================================================
	// Um canal unbuffered (sem buffer) tem capacidade zero. Toda operação de
	// envio (ch <- valor) bloqueia até que outra goroutine esteja pronta para
	// receber (<-ch), e vice-versa. Isso força a sincronização entre as
	// goroutines envolvidas.
	fmt.Println("=== Goroutines com canal de sincronização ===")

	// Cria um canal de sinalização do tipo 'chan bool'.
	// Canais são tipados: só podem transportar valores do tipo declarado.
	done := make(chan bool)

	// Inicia a goroutine. A palavra-chave 'go' precede a chamada de função
	// e faz com que ela execute concorrentemente em uma nova goroutine
	// (uma linha de execução leve gerenciada pelo runtime do Go).
	go greet("Alice", done)
	go greet("Bob", done)

	// Aguarda ambas as goroutines terminarem.
	// Cada receive (<-done) bloqueia até que algo seja enviado no canal.
	// Como temos duas goroutines, fazemos dois receives.
	<-done
	<-done

	fmt.Println("Ambas as saudações foram concluídas.\n")

	// =========================================================================
	// 2. Produtor e Consumidor (unbuffered)
	// =========================================================================
	// Demonstra o padrão clássico producer-consumer com canais unbuffered.
	fmt.Println("=== Produtor e Consumidor ===")

	ch := make(chan int)
	sig := make(chan bool)

	go producer(ch)
	go consumer(ch, sig)

	<-sig // Aguarda o consumidor terminar.
	fmt.Println()

	// =========================================================================
	// 3. Fibonacci com range sobre canal
	// =========================================================================
	// O 'range' sobre um canal é um loop que recebe valores até o canal ser
	// fechado com close(). Se o canal nunca for fechado, o range causaria
	// deadlock (todas as goroutines dormindo).
	fmt.Println("=== Fibonacci com range sobre canal ===")

	fibCh := make(chan int)
	go fib(fibCh, 10)

	for f := range fibCh {
		fmt.Printf("Fib: %d\n", f)
	}
	fmt.Println()

	// =========================================================================
	// 4. Canal bufferizado (buffered channel)
	// =========================================================================
	// Um canal bufferizado tem capacidade > 0. O envio só bloqueia quando
	// o buffer está cheio; o recebimento só bloqueia quando o buffer está
	// vazio. Isso reduz a sincronização estrita entre produtor e consumidor.
	fmt.Println("=== Canal bufferizado ===")

	// Cria um canal bufferizado com capacidade 3.
	bufCh := make(chan string, 3)

	// Como o buffer tem espaço para 3 valores, podemos enviar 3 valores sem
	// uma goroutine receptora correspondente.
	bufCh <- "msg1"
	bufCh <- "msg2"
	bufCh <- "msg3"

	// Agora podemos receber os 3 valores sem bloquear (o buffer está cheio).
	fmt.Println(<-bufCh)
	fmt.Println(<-bufCh)
	fmt.Println(<-bufCh)
	fmt.Println()

	// =========================================================================
	// 5. Select: multiplexação de canais
	// =========================================================================
	// A declaração 'select' permite que uma goroutine aguarde múltiplas
	// operações de canal simultaneamente. Ela bloqueia até que um dos seus
	// casos possa prosseguir. Se vários estiverem prontos, um é escolhido
	// aleatoriamente.
	fmt.Println("=== Select: multiplexação de canais ===")

	chan1 := make(chan string)
	chan2 := make(chan string)

	// Goroutine que envia mensagens com tempos diferentes.
	go func() {
		time.Sleep(50 * time.Millisecond)
		chan1 <- "mensagem do canal 1"
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		chan2 <- "mensagem do canal 2"
	}()

	// O select aguarda o primeiro canal que estiver pronto.
	select {
	case msg1 := <-chan1:
		fmt.Println("Recebido:", msg1)
	case msg2 := <-chan2:
		fmt.Println("Recebido:", msg2)
	}

	// =========================================================================
	// 6. Select com timeout e caso default
	// =========================================================================
	// É possível usar 'select' com time.After para implementar timeouts, e
	// com 'default' para operações não-bloqueantes.
	fmt.Println("\n=== Select com timeout ===")

	select {
	case msg := <-chan1:
		fmt.Println("Extra do chan1:", msg)
	case <-time.After(10 * time.Millisecond):
		// time.After retorna um canal que envia um valor após a duração.
		// Se nenhum canal estiver pronto dentro de 10ms, este caso executa.
		fmt.Println("Timeout: nenhuma mensagem recebida a tempo")
	}

	fmt.Println("\nFim do exemplo!")
}
