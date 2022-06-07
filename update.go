package main

import (
	"context"
	"fmt"
)

func AlterDado(cod, campo, valor string) (int64, error) {
	var sql string = fmt.Sprintf("UPDATE contato SET %v=$1 WHERE cod=$2", campo)

	result, err := conn.Exec(context.Background(), sql, valor, cod)
	if err != nil {
		err = fmt.Errorf("erro ao alterar dados %v", err)
		return 0, err
	}
	return result.RowsAffected(), nil
}
