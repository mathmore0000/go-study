package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

/*
   Exercício 3 - Registro de clientes
   O mesmo estudo do exercício anterior solicita uma funcionalidade para poder registrar novos dados de clientes. Os dados necessários são:

   File
   Name
   ID
   Phone number
   Adress

   Atividade 1: Antes de registrar um cliente, é necessário verificar se o cliente já existe. Para fazer isso, você precisa ler os dados de uma matriz.
   Caso ele se repita, você precisa tratar o erro adequadamente, como vimos até agora. Esse erro deve:
      1.- gerar um panic;

      2.- console iniciar a mensagem: “Error: client already exists”, e continuar com a execução do programa normalmente.

   Atividade 2: Depois de tentar verificar se o cliente a ser registrado já existe, desenvolva uma função para validar se todos os dados
      a serem registrados para um cliente contêm um valor diferente de zero.
   Essa função deve retornar pelo menos dois valores. Um deles deverá ser do tipo erro, caso um valor zero seja inserido como parâmetro
      (lembre-se dos valores zero de cada tipo de dado, por exemplo: 0, “”, nil).

   Atividade 3: Antes de encerrar a execução, mesmo que ocorram panics, as seguintes mensagens devem ser
      impressas no console: “End of execution” e “Several errors were detected at runtime”. Use o defer  para atender a esse requisito..

   Requisitos gerais:

   Use a recover para recuperar o valor de qualquer pânico que possa ocorrer.
   Lembre-se de realizar as validações necessárias para cada retorno que possa conter um valor de erro.
   Gere um erro, personalizando-o de acordo com sua preferência usando uma das funções Go (execute também a validação relevante para o caso de erro retornado).
*/

type Cliente struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

var caminhoRelativoArquivoClientes string = "clientes.json"
var clientes map[int]Cliente

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("End of execution")
		fmt.Println("Several errors were detected at runtime")
	}()
	carregarClientes()
	fmt.Println(clientes)

	// criar struct novo cliente
	var cliente Cliente = Cliente{
		ID:          2,
		Name:        "Matheus Moreira",
		PhoneNumber: "11958101160",
		Address:     "Melicidade",
	}

	// inserir novo cliente
	_, err := inserirNovoCliente(clientes, &cliente)
	if err != nil {
		fmt.Println(err)
		return
	}

	// printar clientes
	fmt.Println(clientes)
}

func inserirNovoCliente(clientes map[int]Cliente, cliente *Cliente) (*Cliente, error) {
	if _, ok := clientes[cliente.ID]; ok == true {
		panic("Error: client already exists")
	}

	if _, err := isClienteValido(cliente); err != nil {
		return cliente, err
	}
	clientes[cliente.ID] = *cliente
	err := os.WriteFile(caminhoRelativoArquivoClientes, pegarClientesMapToClientesByteArray(clientes), 0644)

	return cliente, err

}

func pegarClientesMapToClientesByteArray(currentClientes map[int]Cliente) (clientesByteArrary []byte) {
	var clientesArray []Cliente
	for _, cliente := range currentClientes {
		clientesArray = append(clientesArray, cliente)

	}
	var err error
	clientesByteArrary, err = json.Marshal(clientesArray)
	if err != nil {
		panic("json não pode ser convertido para tipo map")
	}
	return

}

func pegarValorArquivo() (valorArquivo []byte) {
	var err error
	valorArquivo, err = os.ReadFile(caminhoRelativoArquivoClientes)

	if err != nil {
		panic("arquivo clientes.json não encontrado")
	}
	return
}

func pegarValorArquivoEmJSON() (clientes []Cliente) {
	arquivoClientes := pegarValorArquivo()
	var err error = json.Unmarshal(arquivoClientes, &clientes)

	// Check if is there any error while filling the instance
	if err != nil {
		panic(err)
	}
	return
}

func isClienteValido(cliente *Cliente) (campoInvalido string, err error) {
	if cliente.ID == 0 {
		return "id", errors.New("Campo ID é inválido")
	}
	if cliente.Name == "" {
		return "Name", errors.New("Campo Name é inválido")
	}
	if cliente.PhoneNumber == "" {
		return "PhoneNumber", errors.New("Campo PhoneNumber é inválido")
	}
	if cliente.Address == "" {
		return "Address", errors.New("Campo Address é inválido")
	}
	return "", nil
}

func carregarClientes() {
	clientesValorJson := pegarValorArquivoEmJSON()

	clientes = pegarClientesJsonToClientesMap(clientesValorJson)
}

func pegarClientesJsonToClientesMap(clientes []Cliente) (clientesMap map[int]Cliente) {
	clientesMap = make(map[int]Cliente, len(clientes))
	for _, cliente := range clientes {
		clientesMap[cliente.ID] = cliente
	}
	return
}
