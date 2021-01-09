package services

import (
	"github.com/martinyonathann/bookstore_items-api/domain/items"

	"github.com/martinyonathann/bookstore_items-api/utils/errors"
)

func GetAll(flagActive string) (items.Items, *errors.RestErr) {
	dao := &items.Item{}
	return dao.GetAllBooks(flagActive)
}
