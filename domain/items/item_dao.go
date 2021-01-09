package items

import (
	"net/http"

	"github.com/martinyonathann/bookstore_items-api/datasource/mysql/users_db"
	"github.com/martinyonathann/bookstore_items-api/utils/errors"
)

const (
	queryGetAllBook = "SELECT * FROM books where flag_active = ? ;"
)

func (item *Item) GetAllBooks(flagActive string) ([]Item, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryGetAllBook)
	if err != nil {
		return nil, &errors.RestErr{
			Message: "Error when trying to prepare GetAllBooks Statement",
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
	}
	defer stmt.Close()
	rows, err := stmt.Query(flagActive)
	if err != nil {
		return nil, &errors.RestErr{
			Message: "Error when trying to Query GetAllBooks Statement",
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
	}
	defer rows.Close()

	result := make([]Item, 0)
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.BookName, &item.Detail, &item.Price, &item.Writer, &item.YearCreated, &item.FlagActive); err != nil {
			return nil, &errors.RestErr{
				Message: "error when scan user row into user struct",
				Status:  http.StatusInternalServerError,
				Error:   err.Error(),
			}
		}
		result = append(result, item)
	}
	if len(result) == 0 {
		return nil, &errors.RestErr{
			Message: "no users matching status",
			Status:  http.StatusNotFound,
			Error:   err.Error(),
		}
	}
	return result, nil
}
