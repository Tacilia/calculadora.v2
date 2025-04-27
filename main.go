package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*func main() {
	var operando1, operando2 float64
	var operacao string
	var resultado float64

	fmt.Println("Calculadora Ligada")
	fmt.Println("-------------------")

	// Lê o primeiro número
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&operando1)

	// Lê a operação
	fmt.Print("Digite a operação (soma, subtracao, multiplicacao, divisao): ")
	fmt.Scanln(&operacao)

	// Lê o segundo número
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&operando2)

	// Estrutura de decisão baseado em uma condição.Aqui Estrutura que faz o cálculo de cada operacao.
	if operacao == "soma" {
		resultado = operando1 + operando2
	} else if operacao == "subtracao" {
		resultado = operando1 - operando2
	} else if operacao == "multiplicacao" {
		resultado = operando1 * operando2
	} else if operacao == "divisao" {
		if operando2 == 0 {
			fmt.Println("Erro: Nenhum número pode ser dividido por zero!")
			return
		}
		resultado = operando1 / operando2 // variável resultado que guarda e retorna o resultados da  operacao.
	}

	// Mostra o resultado e finaliza a calculadora
	fmt.Printf("Resultado da %s: %.2f\n", operacao, resultado)
	fmt.Println("-------------------")
	fmt.Println("Calculadora Desligada")
	fmt.Println("-------------------")
}*/

// Estrutura do Tipo OperacaoRequest
type OperacaoRequest struct {
	Operando1 float64 `json:"operando1"`
	Operando2 float64 `json:"operando2"`
	Operacao  string  `json:"operacao"`
}

// Estrutura do Tipo ResultadoResponse
type ResultadoResponse struct {
	Resultado float64 `json:"resultado"`
}

// Funcao princial de entrada que inicia o servidor
func main() {
	http.HandleFunc("/soma", somaHandler) //Rota de cada Api
	http.HandleFunc("/subtracao", subtracaoHandler)
	http.HandleFunc("/multiplicacao", multiplicacaoHandler)
	http.HandleFunc("/divisao", divisaoHandler)

	fmt.Println("Servidor rodando na porta 8080...") // Imprime a menssagem de que o servidor está rodando na porta 8080.
	http.ListenAndServe(":8080", nil)
}

// Função que responde a um pedido (requisição) que chega no seu servidor.
func somaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	} // Checando se a requiscao é metodo POST, se nao for envia a menssagem metodo nao permitido.

	var req OperacaoRequest //req é uma variável do tipo OperacaoRequest.
	json.NewDecoder(r.Body).Decode(&req)

	resultado := req.Operando1 + req.Operando2 //Crie uma variável chamada resultado e guarde dentro dela o valor do primeiro número e segundo número que chegou como resposta na requisição.

	json.NewEncoder(w).Encode(ResultadoResponse{Resultado: resultado}) // Serve para mandar uma resposta no formato JSON para quem chamou a sua API.
}

func subtracaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req OperacaoRequest
	json.NewDecoder(r.Body).Decode(&req)

	resultado := req.Operando1 - req.Operando2

	json.NewEncoder(w).Encode(ResultadoResponse{Resultado: resultado})
}

func multiplicacaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req OperacaoRequest
	json.NewDecoder(r.Body).Decode(&req)

	resultado := req.Operando1 * req.Operando2

	json.NewEncoder(w).Encode(ResultadoResponse{Resultado: resultado})
}

func divisaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req OperacaoRequest
	json.NewDecoder(r.Body).Decode(&req)

	//Estrutura condicional que coloca uma condição: náo é permitido que o operando2 seja igual 0.
	if req.Operando2 == 0 {
		http.Error(w, "Erro: nenhum número pode ser dividido por zero", http.StatusBadRequest)
		return
	}

	resultado := req.Operando1 / req.Operando2

	json.NewEncoder(w).Encode(ResultadoResponse{Resultado: resultado})
}
