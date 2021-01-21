package items

import (
	"log"
	"net/http"

	"github.com/martinyonathann/bookstore_items-api/datasource/mysql/users_db"
	"github.com/martinyonathann/bookstore_items-api/utils/errors"
)

const (
	queryGetAllBook = "SELECT * FROM items where flag_active = ? ;"
	queryGetBook    = "SELECT * FROM items where id = ? ;"
	queryValidate   = "SELECT COUNT(*) FROM items where book_name = ? ;"
	queryCreateBook = "INSERT INTO items (book_name, detail_book, price, writer, year_created, flag_active) VALUES (?, ?, ?, ?, ?, ?)"
)

func (item *Item) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetBook)
	if err != nil {
		return errors.NewBadRequestError("error when trying to prepare get items")
	}
	defer stmt.Close()
	result := stmt.QueryRow(item.ID)
	if err := result.Scan(&item.ID, &item.BookName, &item.Detail, &item.Price, &item.Writer, &item.YearCreated, &item.FlagActive); err != nil {
		return &errors.RestErr{
			Message: "error when scan items row into books struct",
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
	}
	return nil
}

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
			Error:   "Failed",
		}
	}
	return result, nil
}

func validate(BookName string) (int, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryValidate)
	log.Println("sudah masuk validate " + BookName)
	if err != nil {
		return 0, errors.NewInternalServerError("error when prepare query Validate")
	}
	defer stmt.Close()

	rows := stmt.QueryRow(BookName)

	var count int
	if err := rows.Scan(&count); err != nil {
		return 0, errors.NewInternalServerError("error when count rows")
	}
	return count, nil
}

func (item *Item) CreateBook() *errors.RestErr {
	count, Counterr := validate(item.BookName)
	if Counterr != nil {
		return errors.NewInternalServerError("error when validate")
	}
	if count > 0 {
		return errors.NewInternalServerError("book already exist")
	}

	stmt, err := users_db.Client.Prepare(queryCreateBook)
	if err != nil {
		return errors.NewInternalServerError("error when prepare query createbook")
	}
	defer stmt.Close()
	insertResult, SaveErr := stmt.Exec(item.BookName, item.Detail, item.Price, item.Writer, item.YearCreated, item.FlagActive)
	if SaveErr != nil {
		return errors.NewInternalServerError("error when execute query createbook")
	}
	bookID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error when get bookID for sequence")
	}
	item.ID = bookID
	return nil
}
