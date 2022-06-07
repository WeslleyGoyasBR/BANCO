package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type contact struct {
	cod       uuid.UUID
	nome      string
	sobrenome string
	telefone  string
	cidade    string
}

func insertAcontact(c contact) error {
	var sql string = "INSERT INTO contato(cod, nome, sobrenome, telefone, cidade)VALUES ($1, $2, $3,$4, $5)"

	_, err := conn.Exec(context.Background(), sql, c.cod, c.nome, c.sobrenome, c.telefone, c.cidade)
	if err != nil {
		err = fmt.Errorf("erro ao inserir dados %v", err)
		return err
	}
	return nil
}
