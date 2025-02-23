package entities

type Product struct {
	ID          uint64
	Name        string
	Price       string
	Description string
	CategoryID  uint64
	Inventory   int
}
