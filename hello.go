package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	exibeIntroducao()
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
	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")

}

func usandoIf(comando int) {
	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do Programa")
	} else {
		fmt.Println("Comando inválido")
	}
}
