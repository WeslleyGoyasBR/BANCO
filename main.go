package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v4"
)

func ConnectDb() (*pgx.Conn, error) {
	connstr := os.Getenv("DATATESTE_URL")
	if len(connstr) == 0 {
		err := fmt.Errorf("URL não estabelecida ou não encontrada")
		return nil, err

	}

	conn, err := pgx.Connect(context.Background(), connstr)
	if err != nil {
		err = fmt.Errorf("sem acesso ao banco de dados [%v]: %v", connstr, err)
		return nil, err
	}
	return conn, nil
}

var conn *pgx.Conn

const EXIT_ERR_NO_CONN = 1
const EXIT_ERR_NO_ARGS = 2
const EXIT_ERR_INS = 3
const EXIT_ERR_DEL = 4
const EXIT_ERR_ALT = 5
const EXIT_ERR_QUERY = 6

func init() {
	var err error
	conn, err = ConnectDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, "impossivel estabelecer conexão %v\n", err)
		os.Exit(EXIT_ERR_NO_CONN)
	}
}
func main() {
	defer conn.Close(context.Background())
	var err error
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Faltou argumentos:\n\tUse:\n%v <INCLUIR|APAGAR|ALTERAR|CONSULTAR> ...\n", os.Args[0])
		os.Exit(EXIT_ERR_NO_ARGS)
	}
	operacao := os.Args[1]
	if operacao == "INCLUIR" {
		ct := contact{
			cod:       uuid.New(),
			nome:      os.Args[2],
			sobrenome: os.Args[3],
			telefone:  os.Args[4],
			cidade:    os.Args[5],
		}
		err = insertAcontact(ct)
		if err != nil {
			fmt.Fprintf(os.Stderr, "não foi possivel inserir contato %v\n", err)
			os.Exit(EXIT_ERR_INS)
		}
	} else if operacao == "APAGAR" {

		err = DeletaDado(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "impossivel deletar os dados %v\n", err)
			os.Exit(EXIT_ERR_DEL)

		}
	} else if operacao == "ALTERAR" {
		count, err := AlterDado(os.Args[2], os.Args[3], os.Args[4])
		if err != nil {
			fmt.Fprintf(os.Stderr, "impossilvel alterar os dados %v", err)
			os.Exit(EXIT_ERR_ALT)
		}
		fmt.Printf("SUCESSO! Alterou %v registros\n", count)

	} else if operacao == "CONSULTAR" {
		user, err := SearchForname(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "ipossivel fazer a consulta no banco de dados %v\n", err)
			os.Exit(EXIT_ERR_QUERY)
		}

		fmt.Printf("dados gerados pelo banco de dados:\n")
		fmt.Printf("cod\t%v\n", user.cod)
		fmt.Printf("nome\t%v\n", user.nome)
		fmt.Printf("sobrenome\t%v\n", user.sobrenome)
		fmt.Printf("telefone\t%v\n", user.telefone)
		fmt.Printf("cidade\t%v\n", user.cidade)
	} else {
		fmt.Fprint(os.Stderr, "OPERACAO INVALIDA %v, as operacoes validas são INCLUIR, ALTERAR, CONSULTAR, APAGAR\n", operacao)
	}

}
