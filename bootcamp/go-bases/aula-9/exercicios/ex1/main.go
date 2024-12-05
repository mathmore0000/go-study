package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Estudante struct {
	Matricula uuid.UUID `json:"matricula"`
	Nome      string    `json:"nome"`
	Telefone  string    `json:"telefone"`
	Email     string    `json:"email"`
}

var caminhoRelativoArquivoEstudantes string = "estudantes.json"
var estudantes map[uuid.UUID]Estudante

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	carregarEstudantes()

	var opcao int
	for {
		printarTabelaPrincipal()
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			criarNovoEstudante()
		case 2:
			listarTodosEstudantes()
		case 3:
			listarUmEstudante()
		case 4:
			alterarEstudante()
		case 5:
			excluirEstudante()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Opção não econtrada\n")
		}
	}
}

func printarTabelaPrincipal() {
	fmt.Println("=== Menu de Opções ===")
	fmt.Println("1 - Incluir Aluno")
	fmt.Println("2 - Listar Todos os Alunos")
	fmt.Println("3 - Pesquisar Aluno por Matrícula")
	fmt.Println("4 - Alterar Aluno")
	fmt.Println("5 - Excluir Aluno")
	fmt.Println("6 - Sair")
}

func alterarEstudante() {
	if len(estudantes) == 0 {
		fmt.Println("Nenhum estudante cadastrado!\n")
		return
	}
	fmt.Print("Matricula: ")
	uuidStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler o UUID:", err)
		return
	}
	uuidStr = strings.TrimSpace(uuidStr)

	uuidMatricula, err := uuid.Parse(uuidStr)
	if err != nil {
		fmt.Println("UUID inválido:", err)
		return
	}

	_, ok := estudantes[uuidMatricula]
	if !ok {
		fmt.Printf("Estudante com matricula %v não foi encontrado\n", uuidMatricula)
		return
	}

	fmt.Print("Nome: ")
	nome, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler o nome:", err)
		return
	}
	nome = strings.TrimSpace(nome)

	fmt.Print("Telefone: ")
	telefone, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler o telefone:", err)
		return
	}
	telefone = strings.TrimSpace(telefone)

	fmt.Print("E-mail: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler o e-mail:", err)
		return
	}
	email = strings.TrimSpace(email)

	// inserir novo estudante
	var novoEstudante Estudante = Estudante{
		Matricula: uuidMatricula,
		Nome:      nome,
		Telefone:  telefone,
		Email:     email,
	}

	estudantes[novoEstudante.Matricula] = novoEstudante
	escreverEstudantesParaJson()
	fmt.Println("Estudante atualizado!\n")

}

func criarNovoEstudante() {
	// criar struct novo estudante
	fmt.Print("Nome: ")
	nome, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler entrada:", err)
		return
	}
	nome = strings.TrimSpace(nome) // Remove o '\n' do final

	fmt.Print("Telefone: ")
	telefone, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler entrada:", err)
		return
	}
	telefone = strings.TrimSpace(telefone)

	fmt.Print("E-mail: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler entrada:", err)
		return
	}
	email = strings.TrimSpace(email)

	// inserir novo estudante
	var novoEstudante Estudante = Estudante{
		Matricula: uuid.New(),
		Nome:      nome,
		Telefone:  telefone,
		Email:     email,
	}

	estudantes[novoEstudante.Matricula] = novoEstudante
	escreverEstudantesParaJson()

	fmt.Println("Estudante criado!\n")
}

func excluirEstudante() {
	if len(estudantes) == 0 {
		fmt.Println("Nenhum estudante cadastrado!\n")
		return
	}
	fmt.Print("UUID Matrícula: ")
	matricula, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler o UUID:", err)
		return
	}
	matricula = strings.TrimSpace(matricula)

	uuidMatricula, err := uuid.Parse(matricula)
	if err != nil {
		fmt.Println("UUID inválido:", err)
		return
	}

	_, ok := estudantes[uuidMatricula]
	if !ok {
		fmt.Printf("Estudante com matricula %v não foi encontrado\n", uuidMatricula)
		return
	}

	delete(estudantes, uuidMatricula)
	escreverEstudantesParaJson()

	fmt.Println("Estudante excluido!\n")
}

func escreverEstudantesParaJson() {
	err := os.WriteFile(caminhoRelativoArquivoEstudantes, pegarEstudantesMapToEstudantesByteArray(estudantes), 0644)
	if err != nil {
		fmt.Println("Erro ao escrever estudantes para json: ", err)
		return
	}
}
func listarUmEstudante() {
	if len(estudantes) == 0 {
		fmt.Println("Nenhum estudante cadastrado!\n")
		return
	}
	fmt.Print("UUID Matrícula: ")
	uuidStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler o UUID:", err)
		return
	}
	uuidStr = strings.TrimSpace(uuidStr)

	uuidMatricula, err := uuid.Parse(uuidStr)
	if err != nil {
		fmt.Println("UUID inválido:", err)
		return
	}
	estudante, ok := estudantes[uuidMatricula]
	if !ok {
		fmt.Printf("Estudante com matricula %v não foi encontrado\n", uuidMatricula)
		return
	}
	printarUmEstudante(&estudante)
}
func listarTodosEstudantes() {
	if len(estudantes) == 0 {
		fmt.Println("Nenhum estudante cadastrado!\n")
		return
	}
	for _, estudante := range estudantes {
		printarUmEstudante(&estudante)
	}
}
func printarUmEstudante(e *Estudante) {
	fmt.Printf("Matricula: %v | Nome: %v | Telefone: %v | E-mail: %v\n", e.Matricula, e.Nome, e.Telefone, e.Email)
}

func pegarEstudantesMapToEstudantesByteArray(currentEstudantes map[uuid.UUID]Estudante) (estudantesByteArrary []byte) {
	var estudantesArray []Estudante
	for _, estudante := range currentEstudantes {
		estudantesArray = append(estudantesArray, estudante)

	}
	var err error
	estudantesByteArrary, err = json.Marshal(estudantesArray)
	if err != nil {
		panic("json não pode ser convertido para tipo map")
	}
	return

}

func pegarValorArquivo() (valorArquivo []byte) {
	var err error
	valorArquivo, err = os.ReadFile(caminhoRelativoArquivoEstudantes)

	if err != nil {
		panic("arquivo estudantes.json não encontrado")
	}
	return
}

func pegarValorArquivoEmJSON() (estudantes []Estudante) {
	arquivoEstudantes := pegarValorArquivo()
	var err error = json.Unmarshal(arquivoEstudantes, &estudantes)

	// Check if is there any error while filling the instance
	if err != nil {
		panic(err)
	}
	return
}

func carregarEstudantes() {
	estudantesValorJson := pegarValorArquivoEmJSON()

	estudantes = pegarEstudantesJsonToEstudantesMap(estudantesValorJson)
}

func pegarEstudantesJsonToEstudantesMap(estudantes []Estudante) (estudantesMap map[uuid.UUID]Estudante) {
	estudantesMap = make(map[uuid.UUID]Estudante, len(estudantes))
	for _, estudante := range estudantes {
		estudantesMap[estudante.Matricula] = estudante
	}
	return
}
