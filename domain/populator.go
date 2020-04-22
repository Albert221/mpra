package domain

type Populator interface {
	Populate([]*Product)
}
