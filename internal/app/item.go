package app

type Item struct {
	key   string
	value string
	items map[string]*Item
}

func (i Item) getValue() string {
	return i.value
}

func (i Item) setValue(value string) {
	i.value = value
}

func newItem(key string, value string) *Item {
	return &Item{
		key:   key,
		value: value,
		items: make(map[string]*Item),
	}
}

func newItemList(key string) *Item {
	return &Item{
		key:   key,
		items: make(map[string]*Item),
	}
}
