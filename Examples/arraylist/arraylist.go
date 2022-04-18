package arraylist

type ArrayList struct {
	data []interface{}
}

// NewArrayList creates a new ArrayList

func NewArrayList() *ArrayList {
	return &ArrayList{
		data: make([]interface{}, 0),
	}
}

// Add adds a new item to the ArrayList
func (a *ArrayList) Add(item interface{}) {
	a.data = append(a.data, item)
}

// Get returns the item at the given index
func (a *ArrayList) Get(index int) interface{} {
	return a.data[index]
}

// Size returns the number of items in the ArrayList
func (a *ArrayList) Size() int {
	return len(a.data)
}

// Remove removes the item at the given index
func (a *ArrayList) Remove(index int) {
	a.data = append(a.data[:index], a.data[index+1:]...)
}
