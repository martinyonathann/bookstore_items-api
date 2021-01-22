package helpers

import (
	"github.com/martinyonathann/bookstore_items-api/datasource/mysql/users_db"
	"github.com/martinyonathann/bookstore_items-api/utils/errors"
)

const (
	queryValidate = "SELECT COUNT(*) FROM items where book_name = ? AND year_created = ?;"
)

func Validate(bookName, year_created string) (int, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryValidate)
	if err != nil {
		return 0, errors.NewInternalServerError("error when prepare query Validate")
	}
	defer stmt.Close()

	rows := stmt.QueryRow(bookName, year_created)

	var count int
	if err := rows.Scan(&count); err != nil {
		return 0, errors.NewInternalServerError("error when count rows")
	}
	return count, nil
}
