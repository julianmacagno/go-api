package database

var (
	kvs = make(map[string]*Item) // should be replaced by a connection to redis or cassandra or something...
)

// Item is kvs item spec
type Item struct {
	Key     string
	Value   []uint64
	Version uint
}

// Get a value from KVS
func Get(key string) *Item {
	return kvs[key]
}

// Set a value on KVS
func Set(item *Item) bool {
	if item == nil {
		return false
	}
	kvs[item.Key] = item
	return true
}
