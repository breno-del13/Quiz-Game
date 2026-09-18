
//----------------------------------------------------------------------------------------------------------------------------------//

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
	"math/rand/v2"
)

// Cores Simples.

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Bold    = "\033[1m"
)

// Cores Especificas.

const (
	// --- TONS VERMELHOS E ROSAS ---
	DarkRed        = "\033[38;5;88m"
	Crimson        = "\033[38;5;160m"
	BrightRed      = "\033[38;5;196m"
	OrangeRed      = "\033[38;5;202m"
	DeepPink       = "\033[38;5;198m"
	HotPink        = "\033[38;5;205m"
	LightPink      = "\033[38;5;211m"
	Salmon         = "\033[38;5;210m"

	// --- TONS LARANJAS E MARRONS ---
	DarkOrange     = "\033[38;5;208m"
	LightOrange    = "\033[38;5;214m"
	Amber          = "\033[38;5;215m"
	Gold           = "\033[38;5;220m"
	Brown          = "\033[38;5;94m"
	Chocolate      = "\033[38;5;130m"
	Tan            = "\033[38;5;180m"

	// --- TONS AMARELOS E VERDES ---
	BrightYellow   = "\033[38;5;226m"
	Khaki          = "\033[38;5;185m"
	DarkGreen      = "\033[38;5;22m"
	ForestGreen    = "\033[38;5;28m"
	LimeGreen      = "\033[38;5;46m"
	BrightGreen    = "\033[38;5;82m"
	Olive          = "\033[38;5;64m"
	SeaGreen       = "\033[38;5;29m"
	Emerald        = "\033[38;5;48m"

	// --- TONS AZUIS ---
	NavyBlue       = "\033[38;5;18m"
	DarkBlue       = "\033[38;5;21m"
	MediumBlue     = "\033[38;5;27m"
	BrightBlue     = "\033[38;5;33m"
	SkyBlue        = "\033[38;5;39m"
	LightBlue      = "\033[38;5;45m"
	Turquoise      = "\033[38;5;43m"
	Cyan256        = "\033[38;5;51m"
	Teal           = "\033[38;5;30m"

	// --- TONS ROXOS E VIOLETAS ---
	Indigo         = "\033[38;5;54m"
	Purple         = "\033[38;5;93m"
	DeepPurple     = "\033[38;5;128m"
	Magenta256     = "\033[38;5;165m"
	DarkMagenta    = "\033[38;5;90m"
	Violet         = "\033[38;5;135m"
	Plum           = "\033[38;5;96m"
	Lavender       = "\033[38;5;147m"

	// --- TONS CINZAS E ESCUROS ---
	DarkGray       = "\033[38;5;235m"
	MediumGray     = "\033[38;5;240m"
	Gray           = "\033[38;5;244m"
	LightGray      = "\033[38;5;250m"
	Silver         = "\033[38;5;7m"
	Charcoal       = "\033[38;5;238m"
	SteelBlue      = "\033[38;5;67m"
	Beige          = "\033[38;5;230m"
	OffWhite       = "\033[38;5;254m"
)


type Question struct {
	text    string
	Options []string
	Answer  int
}

type GameState struct {
	Points    int
	Lifes    int
	name      string
	Questions []Question
	Victorys int
	QuantParts int
}

// Função Responsavel Pela Leitura Do Arquivo CSV.

func (g *GameState) ProcessCSV() {
	f, err := os.Open("Questions.csv")
	if err != nil {
		//panic("Erro ao abrir o arquivo.")
		panic(err)

	}


	defer f.Close()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("┃Error ao ler o arquivo.")

	}
	for index, record := range records {
		//fmt.Println(record)
		if index > 0 {
		question := Question{
			text:    record[0],
			Options: record[1:4],
			Answer:  toInt(record[4]),
			
		}
		g.Questions = append(g.Questions, question)
	}
}


}

// Função Que Inicia e Coleta Dados Para o Jogo.

func (g *GameState) Init() time.Time {
	//Atribuindo...
	Agora := time.Now()
	//Inicializando o jogo
	fmt.Println("\033[H\033[2J╭ Carregando jogo │ ▰▱▱▱▱▱ │")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\033[H\033[2J╭ Carregando jogo │ ▰▰▱▱▱▱ │")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\033[H\033[2J╭ Carregando jogo │ ▰▰▰▱▱▱ │")
	time.Sleep( 500 * time.Millisecond)
	fmt.Print("\033[H\033[2J╭ Carregando jogo │ ▰▰▰▰▱▱ │")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\033[H\033[2J╭ Carregando jogo │ ▰▰▰▰▰▱ │")
	time.Sleep(2 * time.Second)
	fmt.Print("\033[H\033[2J╭ Carregando jogo │ ▰▰▰▰▰▰ │\n")
	time.Sleep(2 * time.Second)

	hora := Agora.Hour()
	TotaldPerguntas := len(g.Questions)

	//Condições.

	if hora >= 23 || hora < 5 {

		fmt.Println("│Olá! Boa madrugada shhhh...., seja bem vindo ao quiz game")
		time.Sleep(2 * time.Second)
	} else if hora >= 18 {

		fmt.Println("│Olá! Boa noite, seja bem vindo ao quiz game")
		time.Sleep(2 * time.Second)
	} else if hora >= 12 {

		fmt.Println("│Olá! Boa tarde, seja bem vindo ao quiz game")
		time.Sleep(2 * time.Second)
	} else {

		fmt.Println("│Olá! Bom dia, seja bem vindo ao quiz game")
		time.Sleep(2 * time.Second)
	}
	fmt.Println("╰ Me diga, qual o seu primeiro nome?: ")
	var name string
	fmt.Scanln(&name)
	g.name = name
	
	time.Sleep(5 * time.Second)
	fmt.Printf("\n┃  Ok então, seu nome é: %s%s%s.\n"+
	"┃  Vamos começar o jogo!\n"+
	"┃\n"+
	"┃  • Você terá 3 VIDAS.\n"+
	"┃  • As perguntas são baseadas na data: 05/07/2026.\n"+
	"┃\n"+
	"┃  OBS: Se você acumular pontos suficientes,\n"+
	"┃  poderá usá-los para comprar mais corações!\n\n" +
	"┃  Quantas perguntas você deseja fazer?: ",CorAleatoria(),g.name, Reset)
	fmt.Scanln(&g.QuantParts)

	if g.QuantParts > TotaldPerguntas {

		fmt.Printf("┃  AVISO: A Quantidade perguntas pedidas são: %d e a Quantidade de perguntas disponivel é: %d\n", g.QuantParts, TotaldPerguntas)
		if TotaldPerguntas == 0{
			fmt.Println("┃  Seu Arquivo CSV está vazio ou tem informações incopativeis a esse Jogo/Game")

			

		}
		fmt.Println("┃  Quantas perguntas você deseja fazer?: ")
		fmt.Scanln(&g.QuantParts)


		if g.QuantParts > TotaldPerguntas{

			if TotaldPerguntas == 0{
				fmt.Println("┃  Seu Arquivo CSV está vazio!!!\n┃  Saindo do jogo...")
				time.Sleep(5 * time.Second)
				return Agora

			} else if g.QuantParts == 0{
				fmt.Printf("┃  Engraçadão em %s HAHAHA(Risada forçada)\n┃ TCHAU ENGRAÇADÃO...", g.name)
				time.Sleep(3 * time.Second)
				return Agora
			}

			fmt.Println("┃  Definindo o Numero de Perguntas ao MAXIMO.\n")
			
			g.QuantParts = TotaldPerguntas
			time.Sleep(1 * time.Second)

		}else if TotaldPerguntas > g.QuantParts{

			fmt.Printf("┃  Você está pedindo %d Perguntas, existem %d Perguntas no seu Arquivo CSV\n", g.QuantParts, TotaldPerguntas)
			fmt.Printf("┃  Mas ja que pediu %d Perguntas, deixarei %d Mesmo...\n", g.QuantParts, g.QuantParts)
		}

	}

	
	return Agora
}


// A RunGame é a função principal do jogo que roda o loop de perguntas,
// Lê as respostas do usuário e controla o sistema de vidas, pontos e fim de jogo.

func (g *GameState) RunGame() { 
	// Exibir a pergunta para o player.
	for i := 0; i < g.QuantParts; i++{

		question := g.Questions[i]
		
		time.Sleep(2 * time.Second)
		fmt.Printf("┃ %s%d: %s%s\n", CorAleatoria(), i+1, question.text, Reset)
		//fmt.Printf("┃ %s%d: %s%s\n", CorAleatoria(), index+1, question.text, Reset)

		time.Sleep(1 * time.Second)
		for j, option := range question.Options{
			fmt.Printf("┃ [%d] • %s\n", j + 1, option)
		}
		
		var answerus int
		var errodresp string 
		fmt.Println("┃ Digite o numero da opção correta: ")
		fmt.Scanln(&answerus)
		for {
			if answerus  == 1 || answerus == 2 || answerus == 3{
				if answerus == question.Answer{
					g.Points += 10
					g.Victorys += 1
					time.Sleep(1 * time.Second)
					fmt.Printf("┃ Resposta correta %s! Você ganhou +10 pontos!\n┋Pontos: %d ┋ Vida: %d\n", g.name, g.Points, g.Lifes)
					fmt.Printf("┃--------------------%d--------------------\n", g.Victorys)
					if g.Victorys == 5{
						var parar string
						time.Sleep(1 * time.Second)
						fmt.Printf("┃ Parabéns %s! Você teve 5 vitorias até agora. Deseja parar de jogar?: ", g.name)
						fmt.Scanln(&parar)
						if parar == "Sim" || parar == "s" {
							time.Sleep(1 * time.Second)
							fmt.Printf("┃ Ok %s, você decidiu parar de jogar. Até a próxima!\n", g.name)
							return
						} else if parar == "não" || parar == "n"{
							fmt.Printf("┃ Ok %s, você decidiu continuar a jogar, BORA!!!\n", g.name)
						}
						
					}else if g.Victorys == g.QuantParts{
						var pararg string
						fmt.Printf("┃ Parabéns %s! VOCÊ CAGABIRITOU!!! Deseja continuar? E com um PREMIO de 50 Pontos?: ", g.name)
						fmt.Scanln(&pararg)
						if pararg == "Sim" || pararg == "s" {
							time.Sleep(1 * time.Second)
							g.Points += 50
							var gabarito int 
							fmt.Printf("┃ Ok %s, você decidiu continuar a jogar, BORA!!!\n", g.name)
							fmt.Printf("┃ %s, Quantas perguntas quer agora?: ", g.name)
							fmt.Scanln(&gabarito)
							g.QuantParts = gabarito
							
							
						} else if pararg == "não" || pararg == "n"{
							fmt.Printf("┃ Ok %s, você decidiu parar de jogar. Até a próxima!\n", g.name)
							return
						} else {
							var gabarito int 
							fmt.Printf("┃ %s Pense em Tentar: \"n\", \"não\", \"Sim\" ou \"s\"\n", g.name)
							fmt.Printf("┃ %s, Quantas perguntas quer agora?: ", g.name)
							fmt.Scanln(&gabarito)
							g.QuantParts = gabarito
						}
					} 
				} else {
					 	g.Points -= 5
					 	g.Lifes -= 1
					 	fmt.Printf("┃ Ops... Você errou %s, Você -5 Pontos\n┃ E tem %d VIDAS não as Desperdice....\n", g.name, g.Lifes)
					 	fmt.Printf("┃--------------------------------------------- ┃\n")
						if  i == g.QuantParts - 1 {
							var parar string
							time.Sleep(1 * time.Second)
							fmt.Printf("┃ Você teve %d VITÓRIAS até agora, Deseja continuar?: ", g.Victorys )
							fmt.Scanln(&parar)
							if parar == "Sim" || parar == "sim" || parar == "s" || parar == "Ok" || parar == "ok"{
								var respp int
								fmt.Println("Quantas perguntas mais quer responder?: ")
								fmt.Scanln(&respp)
								g.QuantParts = respp
								i = -1
								g.Victorys = 0

							} else {
								return
							}
						} else if g.Lifes == 0{
				    		if g.Points >= 20{

								var buyhearts string


								for{
									fmt.Printf("┃ Você está sem vida(Corações: %d)\n", g.Lifes)
									fmt.Printf("┃ Custa 20 Pontos cada, e você tem %d Pontos\n", g.Points)
									fmt.Printf("┃ Deseja Comprar mais corações %s???: ", g.name)
									fmt.Scanln(&buyhearts)

									if buyhearts == "sim"|| buyhearts =="Sim" || buyhearts == "s"|| buyhearts == "S" || buyhearts == "Ok" || buyhearts == "ok" || buyhearts == "SIM"{
										if g.QuantParts == 0 {
											g.Lifes += 1
											g.Points -= 5
											fmt.Printf("┃ Você Comprou corações você tem agora: Pontos: %d Corações: %d\n┃ Mais também tem %d Perguntas para responder. Agora quantas perguntas mais você quer?:", g.Points, g.Lifes, g.QuantParts)
											fmt.Scanln(&g.QuantParts)
											time.Sleep(1 * time.Second)
											
										} else if g.QuantParts >= 1{
											g.Lifes += 1 
											g.Points -= 5 
											fmt.Printf("┃ Você Comprou corações você tem agora: Pontos: %d Corações: %d\n ┃ Você ainda tem %d Partidas para completar.\n", g.Points, g.Lifes, g.QuantParts)
											break
										}
										g.Lifes += 1
										g.Points -= 5
										fmt.Printf("┃ Você Comprou corações você tem agora: Pontos: %d Corações: %d\n", g.Points, g.Lifes)
										return
									} else if buyhearts == "não"|| buyhearts =="Não" || buyhearts == "n"|| buyhearts == "N" || buyhearts == "nao" || buyhearts == "Nao" || buyhearts == "NAO" || buyhearts =="NÃO"{
										fmt.Println("┃ Você está sem vida, e decidiu não comprar corações\n")
										fmt.Printf("┃ %s MORREU...F", g.name)
										time.Sleep(1 * time.Second)
										return 
										


									} else{
										fmt.Printf("┃ %s Tente Usar Respostas como: Não, não, nao, Nao, NAO, n, s, Sim, sim, SIM, Ok ou ok\n", g.name)
										fmt.Printf("┃ %s Você deseja comprar mais Corações/Vidas?: ", g.name)
										fmt.Scanln(&buyhearts)
										time.Sleep(1 * time.Second)
									}
								}
					} else{
						fmt.Printf("┃ %s Morreu por falta de PONTOS e CORAÇÕES...F", g.name)
						
						return
					}
				  }
					
				}
				
				break

			} else{
				errodresp = "┃ Erro: Tente novamente escrevendo o numero da alternativa correta que fica ao lado dela!"
				fmt.Println(errodresp)
			
				fmt.Scanln(&answerus)
			}
		}
	}
	if g.Victorys == 0 && g.QuantParts > 0{
		fmt.Printf("┃ Nossa... %s! você errou todas...\n", g.name)
		fmt.Printf("┃ %s MORREU...\n", g.name)
		return
	}
}


func main() {
	
	fmt.Print("\033[?1049h\033[H") 
	g := GameState{
		Lifes : 3,
	}
	
	g.ProcessCSV()
	Agora := g.Init()
	
	_ = Agora
	g.RunGame()
	
}

// Função que Escolhe uma Cor Aleatoria a uma Certa Frase ou Texto.

func CorAleatoria() string {
	cores := []string{
		
		Red, Green, Yellow, Blue, Bold,
		
		// Vermelhos e Rosas
		"\033[38;5;88m", "\033[38;5;160m", "\033[38;5;196m", "\033[38;5;202m", "\033[38;5;198m", "\033[38;5;205m", "\033[38;5;211m", "\033[38;5;210m",
		// Laranjas e Marrons
		"\033[38;5;208m", "\033[38;5;214m", "\033[38;5;215m", "\033[38;5;220m", "\033[38;5;94m", "\033[38;5;130m", "\033[38;5;180m",
		// Amarelos e Verdes
		"\033[38;5;226m", "\033[38;5;185m", "\033[38;5;22m", "\033[38;5;28m", "\033[38;5;46m", "\033[38;5;82m", "\033[38;5;64m", "\033[38;5;29m", "\033[38;5;48m",
		// Azuis
		"\033[38;5;18m", "\033[38;5;21m", "\033[38;5;27m", "\033[38;5;33m", "\033[38;5;39m", "\033[38;5;45m", "\033[38;5;43m", "\033[38;5;51m", "\033[38;5;30m",
		// Roxos e Violetas
		"\033[38;5;54m", "\033[38;5;93m", "\033[38;5;128m", "\033[38;5;165m", "\033[38;5;90m", "\033[38;5;135m", "\033[38;5;96m", "\033[38;5;147m",
		// Cinzas e Escuros
		"\033[38;5;235m", "\033[38;5;240m", "\033[38;5;244m", "\033[38;5;250m", "\033[38;5;7m", "\033[38;5;238m", "\033[38;5;67m", "\033[38;5;230m", "\033[38;5;254m",
	}
	
	// Escolhe um Índice Aleatório Dentro do Tamanho Total da Lista.

	indice := rand.IntN(len(cores))
	
	return cores[indice]
}

// Funcão que Converte String em Int.

func toInt(s string) int{
	i,err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

//----------------------------------------------------------------------------------------------------------------------------------//
