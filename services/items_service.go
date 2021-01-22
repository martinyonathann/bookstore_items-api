package services

import (
	"github.com/martinyonathann/bookstore_items-api/domain/items"
	"github.com/martinyonathann/bookstore_items-api/utils/errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct {
}

type itemsServiceInterface interface {
	GetItemByID(int64) (*items.Item, *errors.RestErr)
	GetAll(string) (items.Items, *errors.RestErr)
	CreateBook(items.Item) (*items.Item, *errors.RestErr)
}

func (item *itemsService) GetAll(flagActive string) (items.Items, *errors.RestErr) {
	dao := &items.Item{}
	return dao.GetAllBooks(flagActive)
}

func (item *itemsService) GetItemByID(itemsID int64) (*items.Item, *errors.RestErr) {
	result := &items.Item{ID: itemsID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (item *itemsService) CreateBook(book items.Item) (*items.Item, *errors.RestErr) {
	book.FlagActive = "1"
	if err := book.CreateBook(); err != nil {
		return nil, err
	}
	return &book, nil
}
