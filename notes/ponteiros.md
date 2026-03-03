# Ponteiros em Go para Programadores C

Se você já programa em C, vai reconhecer a sintaxe, mas Go simplifica bastante o uso de ponteiros.

## Sintaxe básica

Em Go, a sintaxe é praticamente igual ao C:

```go
x := 42
p := &x          // p é um ponteiro para x
fmt.Println(*p)  // dereferencia: imprime 42
*p = 100         // modifica x através do ponteiro
```

## Declaração de tipos

```go
var p *int       // ponteiro para int, valor zero é nil
p = new(int)     // aloca um int na heap, retorna ponteiro
*p = 10
```

## Diferenças importantes em relação ao C

1. **Sem aritmética de ponteiros**
   Go não permite `p++` ou `p + 2`. Isso elimina uma classe inteira de bugs comuns em C. Se precisar iterar sobre dados, use slices.

2. **Garbage collector — sem free()**
   Não há `malloc`/`free`. O GC cuida da memória. Você aloca com `new()` ou com literais compostos e esquece.

3. **Ponteiros não podem ser nulos de forma silenciosa**
   O compilador e o nil check explícito protegem você. Acessar um nil pointer ainda causa panic, mas é mais fácil de rastrear do que em C.

4. **Escape analysis**
   O compilador decide automaticamente se uma variável vai para a stack ou para a heap — você não precisa (nem pode) controlar isso diretamente.

```go
func novaPersona(nome string) *Pessoa {
    p := Pessoa{Nome: nome} // parece stack...
    return &p               // ...mas Go move para heap automaticamente
}
```
Em C isso seria um bug clássico. Em Go, é perfeitamente válido.

## Quando usar ponteiros em Go

| Situação | Usar ponteiro? |
| :--- | :--- |
| Modificar argumento dentro de função | ✅ Sim |
| Struct grande (evitar cópia) | ✅ Sim |
| Métodos que mutam o receiver | ✅ Sim |
| Valor simples só para leitura | ❌ Não precisa |
| Indicar ausência de valor (nil) | ✅ Sim |

```go
// Sem ponteiro — copia a struct
func (p Pessoa) Saudacao() string {
    return "Olá, " + p.Nome
}

// Com ponteiro — modifica o original
func (p *Pessoa) Renomear(novo string) {
    p.Nome = novo
}
```

## Sintaxe mais limpa que C

Em C você escreve `pessoa->nome`. Em Go, o operador `->` não existe — o compilador dereferencia automaticamente:

```go
p := &Pessoa{Nome: "Ana"}
fmt.Println(p.Nome) // Go faz (*p).Nome automaticamente
```

## Resumo mental

* A mecânica (`&`, `*`, `nil`) é idêntica ao C
* Sem aritmética de ponteiros
* Sem gerenciamento manual de memória
* O compilador é seu aliado, não seu adversário

A principal mudança de mentalidade é: em Go você usa ponteiros para semântica de mutação e identidade, não para gerenciar memória.
