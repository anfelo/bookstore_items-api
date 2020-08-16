package items

import "github.com/anfelo/bookstore_utils/errors"

// ItemsRepository interface
type ItemsRepository interface {
	GetByID(string) (*Item, *errors.RestErr)
	Create(Item) (*Item, *errors.RestErr)
}
