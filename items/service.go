package items

import "github.com/anfelo/bookstore_utils/errors"

// ItemsService interface
type ItemsService interface {
	GetByID(string) (*Item, *errors.RestErr)
	Create(Item) (*Item, *errors.RestErr)
}
