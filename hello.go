package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"bufio"
	// "io/ioutil"
	// "reflect"
)

const monitoramentos = 3
const delay = 5

func main() {

	// exibeNomes()

	exibeIntroducao()
	leSitesDoArquivo()
	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

	// _, idade := devolveNomeEIdade()
	// fmt.Println( "e tenho", idade, "anos")
}

// func devolveNomeEIdade() (string, int) {
// 	nome := "Karlla"
// 	idade := 31
// 	return nome, idade
// }

func exibeIntroducao() {
	nome := "Karlla"
	versao := 1.1
	fmt.Println("Hello, sra.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	//leitura de arquivos com uso convencional de slice
	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	apenas leitura e print de arquivo
	arquivo, err := ioutil.ReadFile("sites.txt")

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	linha, err := leitor.ReadString('\n')
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(linha)

	return sites
}