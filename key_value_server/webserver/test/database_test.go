package test

import (
	"github.com/javiroberts/key_value_server/webserver/model/database"
	"testing"
)

func TestGetEmptyKey(t *testing.T) {
	result := database.Get("")
	if result != nil {
		t.Error("error")
	}
}

func TestSetItem(t *testing.T) {
	item := &database.Item{
		Key:     "key",
		Value:   ["value"],
		Version: 1,
	}
	result := database.Set(item)
	if !result {
		t.Error("error")
	}
}

func TestSetNilItem(t *testing.T) {
	result := database.Set(nil)
	if result {
		t.Error("error")
	}
}
