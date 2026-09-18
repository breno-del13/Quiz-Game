# Quiz Game em Go 🎮

Este projeto é um **jogo de perguntas e respostas desenvolvido em Go (Golang)** durante o curso de **Introdução ao Go da Rocketseat**.

O objetivo do projeto foi colocar em prática os fundamentos da linguagem através da construção de um jogo interativo executado diretamente pelo terminal.

### 🧠 Sobre o projeto

O jogo permite que o jogador informe seu nome, escolha a quantidade de perguntas que deseja responder e participe de uma sequência de perguntas carregadas a partir de um arquivo `CSV`.

Durante a partida, o jogador possui um sistema de **vidas, pontos e vitórias**. Respostas corretas adicionam pontos, enquanto respostas incorretas fazem o jogador perder pontos e vidas. Também existe a possibilidade de utilizar pontos para adquirir novas vidas e continuar jogando.

As perguntas são armazenadas em estruturas (`structs`) e carregadas dinamicamente através da biblioteca `encoding/csv`, enquanto as respostas são processadas e convertidas para números utilizando `strconv`.

### ⚙️ Conceitos de Go utilizados

Durante o desenvolvimento foram utilizados diversos conceitos fundamentais da linguagem, incluindo:

* `structs`
* Métodos
* Funções
* Ponteiros em métodos
* Variáveis e constantes
* Condicionais (`if`, `else if`, `else`)
* Loops (`for`)
* Entrada de dados pelo terminal
* Manipulação de strings
* Conversão de `string` para `int`
* Leitura e processamento de arquivos CSV
* Tratamento de erros
* Pacotes da biblioteca padrão
* Geração de valores aleatórios
* Manipulação de tempo com `time`
* Controle de estado do jogo

Também foram utilizados códigos ANSI para criar diferentes cores e melhorar a apresentação das informações no terminal.

### 🎯 Sistemas implementados

* Sistema de perguntas e respostas
* Sistema de pontuação
* Sistema de vidas/corações
* Compra de vidas utilizando pontos
* Escolha da quantidade de perguntas
* Possibilidade de continuar jogando
* Carregamento das perguntas através de CSV
* Mensagens personalizadas utilizando o nome do jogador
* Saudação de acordo com o horário
* Interface colorida no terminal
* Seleção aleatória de cores para determinadas mensagens

O projeto foi desenvolvido como uma forma de **praticar e consolidar os fundamentos de Go através de um projeto próprio**, indo além de exemplos isolados e aplicando os conceitos aprendidos na construção de uma aplicação funcional.
