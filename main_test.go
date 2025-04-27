package main

import (
	"net/http"          //permite criar um servidor web e lidar com requisições HTTP.É o pacote para criar servidor web, receber requisições e responder.
	"net/http/httptest" // permite simular um servidor HTTP para testar sua API sem precisar rodá-la de verdade.
	"testing"           // pacote usado para escrever testes em Go.
)

// Testando o resultados das operacoes, se os calculos estao corretos
func TestOperacoesSimples(t *testing.T) {
	// Cenário positivo: soma correta
	req := OperacaoRequest{Operando1: 10.0, Operando2: 5.0}
	resultado := req.Operando1 + req.Operando2
	if resultado != 15.0 {
		t.Errorf("Esperado 15.0, e veio %v", resultado)
		// t.Errorf ee a resposta não for o que esperamos, mostra um erro no teste.
	}

	// Cenário negativo: soma errada (esperado incorreto de propósito)
	if resultado == 6.0 {
		t.Errorf("Erro: Algo está errado no cálculo")
	}

	// Cenário positivo: subtração
	req = OperacaoRequest{Operando1: 10.0, Operando2: 5.0}
	resultado = req.Operando1 - req.Operando2
	if resultado != 5.0 {
		t.Errorf("Esperado 5.0, e veio %v", resultado)
	}

	// Cenário negativo: subtração errada
	if resultado == 7.0 {
		t.Errorf("Erro: Algo está errado no cálculo")
	}

	// Cenário positivo: multiplicação
	req = OperacaoRequest{Operando1: 10.0, Operando2: 5.0}
	resultado = req.Operando1 * req.Operando2
	if resultado != 50.0 {
		t.Errorf("Esperado 50.0, e veio %v", resultado)
	}

	// Cenário negativo: multiplicação errada
	if resultado == 80.0 {
		t.Errorf("Erro: Algo está errado no cálculo")
	}

	// Cenário positivo: divisão
	req = OperacaoRequest{Operando1: 10.0, Operando2: 5.0}
	resultado = req.Operando1 / req.Operando2
	if resultado != 2.0 {
		t.Errorf("Esperado 2.0, e veio %v", resultado)
	}

	// Cenário negativo: divisão errada
	if resultado == 4.0 {
		t.Errorf("Erro: Algo está errado no cálculo")
	}
}

// Testa se o handler de SOMA só aceita POST
func TestSomaHandler_AceitaApenasPOST(t *testing.T) {
	req := httptest.NewRequest("GET", "/soma", nil)
	//Cria uma requisição falsa com método GET (errado para simular o navegador ou o Insomnia).

	// Cria um gravador de resposta para simular a resposta do servidor
	res := httptest.NewRecorder()

	//Chama a função somahandler que queremos, passando a requisição e o gravador de resposta, simulando um acesso à API.
	somaHandler(res, req)

	// Verifica se o código de resposta é 405 (método não permitido)
	if res.Code != http.StatusMethodNotAllowed { // Espera erro 405
		// res.Code guarda o código de resposta que a função retornou (tipo 200, 404, 405).
		t.Errorf("Esperava código 405, veio %d", res.Code)
		// t.Errorf ee a resposta não for o que esperamos, mostra um erro no teste.
	}
}

// Testa se o handler de SUBTRAÇÃO só aceita POST
func TestSubtracaoHandler_AceitaApenasPOST(t *testing.T) {
	req := httptest.NewRequest("GET", "/subtracao", nil) // Requisição com método errado
	res := httptest.NewRecorder()

	subtracaoHandler(res, req) // Chama o handler de subtração

	if res.Code != http.StatusMethodNotAllowed { // Espera erro 405
		t.Errorf("Esperava código 405, veio %d", res.Code)
	}
}

// Testa se o handler de MULTIPLICAÇÃO só aceita POST
func TestMultiplicacaoHandler_AceitaApenasPOST(t *testing.T) {
	req := httptest.NewRequest("GET", "/multiplicacao", nil) // Requisição errada
	res := httptest.NewRecorder()

	multiplicacaoHandler(res, req) // Chama o handler de multiplicação

	if res.Code != http.StatusMethodNotAllowed { // Testa se a resposta é 405
		t.Errorf("Esperava código 405, veio %d", res.Code)
	}
}

// Testa se o handler de DIVISÃO só aceita POST
func TestDivisaoHandler_AceitaApenasPOST(t *testing.T) {
	req := httptest.NewRequest("GET", "/divisao", nil) // Criando requisição errada
	res := httptest.NewRecorder()

	divisaoHandler(res, req) // Chamando o handler de divisão

	if res.Code != http.StatusMethodNotAllowed { // Espera que venha erro 405
		t.Errorf("Esperava código 405, veio %d", res.Code)
	}
}
