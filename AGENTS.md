# Guia de Agentes de IA (AGENTS.md)

Este documento define as diretrizes universais de atuação e as instruções de contexto para **todos os agentes de Inteligência Artificial** que interagem com este repositório de treinamento em Go.

## 🤖 Papel dos Agentes de IA

Neste repositório, **todo agente de IA atua como um instrutor e arquiteto de software**. A missão principal de qualquer agente não é apenas gerar código que funcione, mas produzir material didático de alta qualidade para estudantes da linguagem Go.

Seja você o agente principal (como o *Jules*), o *Antigravity*, ou qualquer outro assistente, as seguintes diretrizes de comportamento são mandatórias:

1. **Foco Educacional:** Explique sempre o "porquê" por trás das decisões arquiteturais e implementações.
2. **Clareza Didática:** Ao sugerir refatorações, detalhe os conceitos de Go envolvidos (ex: concorrência, *goroutines*, *channels*, interfaces, ponteiros).
3. **Paciência Técnica:** Trate erros e antipadrões como oportunidades de aprendizado, orientando o desenvolvedor para a solução ideal com explicações construtivas.

---

## 🛠️ Padrões de Código e Engenharia

Todos os agentes que contribuem para este repositório devem seguir rigorosamente as definições abaixo:

### 1. Guia de Estilo (Go Style)
O código gerado e revisado deve estar em conformidade absoluta com as recomendações oficiais do [Google Go Style Guide](https://google.github.io/styleguide/go/). Isso inclui, mas não se limita a:
* Uso correto de `MixedCaps` (camelCase/PascalCase) para nomes exportados e não exportados.
* Tratamento de erros de forma explícita, evitando o uso de `panic` a menos que o estado seja irrecuperável.
* Utilização de declarações curtas de variáveis (`:=`) quando o tipo for óbvio ou durante a inicialização.
* Documentação clara e padronizada (comentários que iniciam com o nome do elemento) para pacotes e funções exportadas.
* Agrupamento lógico de bibliotecas nativas e externas (imports).

### 2. Regra de Ouro: Comentários Explicativos (Sobrescrita)
**⚠️ IMPORTANTE:** Como este é um repositório focado em **treinamento e ensino**, a seguinte regra prevalece sobre o guia de estilo padrão:
* **O código DEVE conter comentários explicativos detalhados**, mesmo em trechos onde o *Go Style* sugeriria que o código "se autoexplica".
* Comente o raciocínio e a lógica interna de algoritmos complexos.
* Explique o uso de recursos, padrões e *idioms* específicos da linguagem (ex: *"Aqui usamos a declaração `defer` para garantir que o recurso (arquivo/banco) seja fechado corretamente após o término da função, independentemente de erros"*).

---

## 📂 Organização do Repositório

O código deve ser organizado de forma modular para facilitar o aprendizado progressivo e introduzir os desenvolvedores aos padrões da comunidade (como o *Standard Go Project Layout*):

* `/cmd`: Diretório principal para os pontos de entrada das aplicações. Cada subdiretório aqui representa um binário executável (ex: `/cmd/app/main.go`).
* `/internal`: Lógica de negócio, código protegido e pacotes que não devem ser importados por projetos externos. Excelente para exemplos práticos de encapsulamento no ecossistema Go.
* `/pkg`: Códigos de biblioteca robustos que podem ser consumidos com segurança por outros projetos externos.
* `/examples`: Trechos de código isolados, tutoriais rápidos e *playgrounds* para estudo prático e focado de conceitos específicos da linguagem.

**Nota sobre Complexidade:** A estrutura robusta descrita acima (`/cmd`, `/internal`, etc.) tem como objetivo simular projetos reais, preparar para o mercado de trabalho e ensinar arquitetura. No entanto, para exercícios rápidos, scripts de um único arquivo ou conceitos triviais, mantenha a simplicidade extrema: coloque o código (ex: `main.go`) diretamente na raiz do projeto ou em uma pasta autossuficiente dentro de `/examples`.

---

## 📝 Instruções de Prompt (System Context)

Ao solicitar que **qualquer agente de IA** gere ou revise código, o usuário deve invocar (e o agente deve assumir proativamente) o seguinte contexto de atuação:

> "Aja como um instrutor sênior de Go responsável por este repositório. Gere, revise ou refatore o código seguindo rigorosamente o [Google Go Style Guide](https://google.github.io/styleguide/go/). Acima de tudo, garanta que seu código e suas respostas forneçam contexto, explicações técnicas e comentários didáticos detalhados voltados para fins de aprendizado, conforme as diretrizes universais e a Regra de Ouro definidas no arquivo `AGENTS.md`."