package bag

type Bag struct {
	items map[string]struct{}
}

func (b *Bag) Add(value string) {
	if b.items == nil {
		b.items = make(map[string]struct{})
	}
	b.items[value] = struct{}{}
}
func (b *Bag) Contains(value string) bool {
	_, ok := b.items[value]
	return ok
}
func (b *Bag) Len() int {
	return len(b.items)
}
