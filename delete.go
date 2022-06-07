package main

import (
	"context"
	"fmt"
)

func DeletaDado(cod string) error {
	var sql string = "DELETE FROM contato WHERE cod=$1"
	_, err := conn.Exec(context.Background(), sql, cod)
	if err != nil {
		err = fmt.Errorf("erro ao deletar os dados %v", err)
		return err
	}
	return nil
}
