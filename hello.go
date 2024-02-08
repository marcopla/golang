package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 3

func main() {

	exibeIntroducao()

	registraLog("site-falso", false)

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	var nome string = "Marco"
	var idade int = 37
	var versao float32 = 1.1
	fmt.Println("Olá, ", nome, "sua idade é", idade)
	fmt.Println("Este software está na versão:", versao)
	fmt.Println("O tipo da variável é: ", reflect.TypeOf(versao))
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	//fmt.Scanf("%d", &comandoLido)
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	//fmt.Println("")
	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	//[]string{"http://random-status-code.herokuapp.com/", "http://alura.com.br", "http://caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, "Site acessado:", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "está on")
		registraLog(site, true)
	} else {
		fmt.Println("O site:", site, "está com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu o erro:", err)
	}
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

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Ocorreu o erro:", err)
	}

	fmt.Println(arquivo)
}
