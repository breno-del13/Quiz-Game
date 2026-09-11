package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
	"math/rand/v2"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Bold    = "\033[1m"
)


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
}

func (g *GameState) Init() time.Time {
	//Atribuindo...
	Agora := time.Now()
	//Inicializando o jogo
	fmt.Println("Iniciando o jogo...")
	time.Sleep(5 * time.Second)

	hora := Agora.Hour()

	//Condições
	// Testamos a noite primeiro
	if hora >= 23 || hora < 5 {
		fmt.Println("Olá! Boa madrugada shhhh...., seja bem vindo ao quiz game")
		time.Sleep(2 * time.Second)
	} else if hora >= 18 {
		fmt.Println("Olá! Boa noite, seja bem vindo ao quiz game")
		time.Sleep(2 * time.Second)
	} else if hora >= 12 {
		fmt.Println("Olá! Boa tarde, seja bem vindo ao quiz game")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("Olá! Bom dia, seja bem vindo ao quiz game")
		time.Sleep(2 * time.Second)
	}
	fmt.Println("Me diga, qual o seu primeiro nome?: ")
	var name string
	fmt.Scanln(&name)
	g.name = name
	time.Sleep(5 * time.Second)
	fmt.Printf("\n ┃  Ok então, seu nome é: %s%s%s.\n"+
	" ┃  Vamos começar o jogo!\n"+
	" ┃\n"+
	" ┃  • Você terá 3 VIDAS.\n"+
	" ┃  • As perguntas são baseadas na data: 05/07/2026.\n"+
	" ┃\n"+
	" ┃  OBS: Se você acumular pontos suficientes,\n"+
	" ┃  poderá usá-los para comprar mais corações!\n\n",CorAleatoria(),g.name, Reset)

	//Retornando
	return Agora
}



func (g *GameState) RunGame() { 
	// Exibir a pergunta para o player
	for index, question := range g.Questions{
		time.Sleep(2 * time.Second)
		fmt.Printf("┃ %s%d: %s%s\n", CorAleatoria(), index+1, question.text, Reset)

		time.Sleep(1 * time.Second)
		for j, option := range question.Options{
			fmt.Printf("┃ [%d] • %s\n", j + 1, option)
		}
		var answerus int
		var errodresp string 
		fmt.Println("┃ Digite o numero da opção correta:")
		fmt.Scanln(&answerus)
		for {
			if answerus  == 1 || answerus == 2 || answerus == 3{
				if answerus == question.Answer{
					g.Points += 10
					g.Victorys += 1
					time.Sleep(1 * time.Second)
					fmt.Printf("┃ Resposta correta %s! Você ganhou 10 pontos e sua vida continua intacta!┃\nPontos: %d ┃ Vida: %d\n", g.name, g.Points, g.Lifes)
					fmt.Printf("┃--------------------%d--------------------\n", g.Victorys)
					if g.Victorys == 5{
						var parar string
						time.Sleep(1 * time.Second)
						fmt.Printf("┃ Parabéns %s! Você teve 5 vitorias até agora. Deseja parar de jogar?\n", g.name)
						fmt.Scanln(&parar)
						if parar == "Sim" || parar == "s" {
							time.Sleep(1 * time.Second)
							fmt.Printf("┃ Ok %s, você decidiu parar de jogar. Até a próxima!\n", g.name)
							return
						} else if parar == "não" || parar == "n"{
							fmt.Printf("┃ Ok %s, você decidiu continuar a jogar, BORA!!!\n", g.name)
						}
						
					}
				} else{
					g.Points -= 5
					g.Lifes -= 1
					fmt.Printf("┃ Ops... Você errou %s\n", g.name)
					fmt.Printf("┃--------------------------------------------- ┃\n")
					if g.Lifes <= 0{
				    if g.Points >= 20{
						var buyhearts string
						fmt.Printf("┃ Você está sem vida, você deseja comprar mais corações? você tem %d pontos,┃\ncustam apenas 20 pontos cada┃\n(digite 'sim' ou 'nao'):\n", g.Points)
						time.Sleep(1 * time.Second)
						fmt.Scanln(&buyhearts)
						if buyhearts == "s" || buyhearts == "sim"||buyhearts == "Ok" {
						g.Points -= 20
						g.Lifes += 1 
						fmt.Printf("┃ Você comprou um coração! Agora você tem %d corações e %d pontos!\n", g.Lifes, g.Points)
						time.Sleep(1 * time.Second)
					}else if (buyhearts == "n" || buyhearts == "não" || buyhearts == "nao") && g.Lifes == 0{
						fmt.Printf("┃ %s MORREU...\n", g.name)
						return
					} else{
						fmt.Printf("┃ %s Tente: n, não, nao, sim, s ou Ok.┃\n", g.name)
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
}

func main() {
	
	g := GameState{
		Lifes : 3,
	}
	Agora := g.Init()
	g.ProcessCSV()
	_ = Agora
	g.RunGame()
	
}

func CorAleatoria() string {
	cores := []string{
		// As 5 que você já tinha no topo como constantes
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
	
	// Escolhe um índice aleatório dentro do tamanho total da lista
	indice := rand.N(len(cores))
	
	return cores[indice]
}


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
func toInt(s string) int{
	i,err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}