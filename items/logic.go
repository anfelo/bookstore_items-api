package items

import "github.com/anfelo/bookstore_utils/errors"

type itemsService struct {
	itemsRepo ItemsRepository
}

// NewItemsService method incharge of creating a new items service
func NewItemsService(itemsRepo ItemsRepository) ItemsService {
	return &itemsService{
		itemsRepo,
	}
}

func (s *itemsService) GetByID(itemID string) (*Item, *errors.RestErr) {
	return nil, errors.NewInternatServerError("not implemented")
}

func (s *itemsService) Create(item *Item) (*Item, *errors.RestErr) {
	return nil, errors.NewInternatServerError("not implemented")
}
