package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Stock Tax CLI")
	fmt.Println("Digite operações em JSON ou 'exit' para sair")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		// sair explicitamente
		if line == "exit" {
			fmt.Println("Encerrando...")
			break
		}

		// ignora linha vazia
		if line == "" {
			fmt.Println("Encerrando...")
			break
		}

		var ops []Operation

		err := json.Unmarshal([]byte(line), &ops)
		if err != nil {
			fmt.Println("Erro ao ler JSON:", err)
			continue
		}

		taxes := ProcessOperations(ops)
		output := FormatTaxes(taxes)

		jsonOutput, err := json.Marshal(output)
		if err != nil {
			fmt.Println("Erro ao gerar saída:", err)
			continue
		}

		fmt.Println(string(jsonOutput))
	}
}
