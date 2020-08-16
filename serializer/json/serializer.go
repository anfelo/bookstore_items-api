package json

import (
	"encoding/json"

	"github.com/anfelo/bookstore_items-api.git/items"
	"github.com/anfelo/bookstore_utils/errors"
)

// ItemSerializer struct
type ItemSerializer struct{}

// Decode method that decodes a []byte into an item
func (i *ItemSerializer) Decode(input []byte) (*items.Item, *errors.RestErr) {
	item := &items.Item{}
	if err := json.Unmarshal(input, item); err != nil {
		return nil, errors.NewBadRequestError("invalid json body")
	}
	return item, nil
}

// Encode method that encodes an item in json format
func (i *ItemSerializer) Encode(input *items.Item) ([]byte, *errors.RestErr) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.NewInternatServerError("error trying to marshal item")
	}
	return rawMsg, nil
}
