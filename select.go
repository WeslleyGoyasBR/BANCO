package main

import (
	"context"
	"fmt"
)

type user struct {
	cod       string
	nome      string
	sobrenome string
	telefone  string
	cidade    string
}

func SearchForname(searchName string) (*user, error) {
	var sql string = "SELECT cod, nome, sobrenome, telefone, cidade FROM contato WHERE nome=$1"
	var user user
	err := conn.QueryRow(context.Background(), sql, searchName).Scan(&user.cod, &user.nome, &user.sobrenome, &user.telefone, &user.cidade)
	if err != nil {
		fmt.Errorf("erro na consulta do dado solicitado %v", err)
		return nil, err
	}
	return &user, nil

}
