package indentifications

import (
	"errors"
	db "github.com/javiroberts/key_value_server/webserver/model/database"
	model "github.com/javiroberts/key_value_server/webserver/model/identifications"
)

func HandleIdentification(id *model.IdentificationRequest) error {
	if id == nil {
		return errors.New("nil identifications")
	}

	item := id.ToKVSItem()
	success := db.Set(&item)
	if !success {
		return errors.New("error when saving")
	}

	return nil
}

func GetIdentification(idType string, idNumber uint64) (*model.IdentificationResponse, error) {
	key := model.GetKVSKey(idType, idNumber)
	item := db.Get(key)

	if item == nil {
		return nil, errors.New("item not found")
	}

	return &model.IdentificationResponse{
		Users: item.Value,
	}, nil
}
