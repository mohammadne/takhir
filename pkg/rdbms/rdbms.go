package rdbms

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type RDBMS interface {
	Execute(query string, in []any) error

	QueryRow(query string, in []any, out []any) error

	Query(query string, in []any, out [][]any) error
}

type rdbms struct {
	db *sql.DB
}

var (
	ErrPrepareStatement = "Error when tying to prepare statement"
	ErrNotFound         = "Error no entry found with given arguments"
	ErrDuplicate        = "Error operation canceled due to the duplication entry"
)

func (db *rdbms) Execute(query string, in []any) error {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s\n%v", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(in...); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New(ErrDuplicate)
		}
		return fmt.Errorf("%s\n%v", "error when tying to excute statement", err)
	}

	return nil
}

func (db *rdbms) QueryRow(query string, in []any, out []any) error {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s\n%v", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	if err = stmt.QueryRow(in...).Scan(out...); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New(ErrDuplicate)
		} else if err == sql.ErrNoRows {
			return errors.New(ErrNotFound)
		}
		return fmt.Errorf("%s\n%v", "Error while executing the query or scanning the row", err)
	}

	return nil
}

func (db *rdbms) Query(query string, in []any, out [][]any) error {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s\n%v", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(in...)
	if err != nil {
		return fmt.Errorf("%s\n%v", "Error executing the query", err)
	}
	defer rows.Close()

	var index = 0
	for ; rows.Next(); index++ {
		if err = rows.Scan(out[index]...); err != nil {
			return fmt.Errorf("%s\n%v", "Error while scanning the row", err)
		}
	}
	out = out[:index+1]

	if err := rows.Err(); err != nil {
		return fmt.Errorf("%s\n%v", "There's an error in result of the query", err)
	}

	return nil
}
