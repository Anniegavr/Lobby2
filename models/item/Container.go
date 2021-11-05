package item

type Container struct {
	items []Item
}

func NewContainer(items []Item) *Container {
	return &Container{items: items}
}

func (container *Container) GetId(id int) (*Item, bool) {
	for i := range container.items {
		if container.items[i].Id == id {
			return &container.items[i], true
		}
	}
	return nil, false
}

func (container *Container) Get(index int) (*Item, bool) {
	if index < 0 || index >= len(container.items) {
		return nil, false
	}
	return &container.items[index], true
}

func (container *Container) GetLen() int {
	return len(container.items)
}
