package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	// "io/ioutil"
	// "reflect"
)

const monitoramentos = 3
const delay = 5

func main() {

	// exibeNomes()

	exibeIntroducao()
	quemExecuta()

	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
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
	versao := 1.1
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("Quem está executando o programa?")
}

func quemExecuta() string {
	var nomeExecutor string
	fmt.Scan(&nomeExecutor)
	fmt.Println("O executor do programa será:", nomeExecutor)

	return nomeExecutor
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
		registraLogs(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
		registraLogs(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	// apenas leitura e print de arquivo
	// arquivo, err := ioutil.ReadFile("sites.txt")

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	// bufio é utilizado para realizar leitura de site linha a linha, extrair o site e salvar informação em um array
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLogs(site string, status bool) {

	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)

	}

	arquivo.WriteString (time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLogs(){

	arquivo, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}