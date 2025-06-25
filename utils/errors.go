package utils

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func IsUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		if mysqlErr.Number == 1062 {
			return true
		}
	}
	return false
}
