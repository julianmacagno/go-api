package identifications

import (
	"fmt"
	kvs "github.com/javiroberts/key_value_server/webserver/model/database"
)

type IdentificationResponse struct {
	Users []uint64 `json:"users"`
}

type IdentificationRequest struct {
	UserId         uint64             `json:"user_id"`
	Identification UserIdentification `json:"identification"`
}

type UserIdentification struct {
	Type   string `json:"type"`
	Number uint64 `json:"number"`
}

func (req IdentificationRequest) ToKVSItem() kvs.Item {
	var result kvs.Item
	result.Key = GetKVSKey(req.Identification.Type, req.Identification.Number)
	var val []uint64
	val = append(val, req.UserId)
	result.Value = val
	result.Version = 1
	return result
}

func GetKVSKey(idtype string, number uint64) string {
	return fmt.Sprintf("%s_%d", idtype, number)
}
