package textadv

type Item struct {
  Entity
  carryable bool
  name string
}

type KeyItem struct {
  Item
}

func (self *Item) IsCarryable() bool {
  return self.carryable
}

func CreateItem(description, name string, carryable bool) (*Item) {
  i := Item{
              Entity: Entity{ description: description },
              name: name,
              carryable: carryable,
  }
  return &i
}
