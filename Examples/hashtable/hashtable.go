package hashtable

type HashTable struct {
	size int
	data []interface{}
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size: size,
		data: make([]interface{}, size),
	}
}

// add adds a new item to the hashtable.
func (h *HashTable) Add(key string, value interface{}) {
	h.data[h.hash(key)] = value
}

// hash function
func (h *HashTable) hash(key string) int {
	var sum int
	for _, c := range key {
		sum += int(c)
	}
	return sum % h.size
}

// get returns the value for a given key.
func (h *HashTable) Get(key string) interface{} {
	return h.data[h.hash(key)]
}
