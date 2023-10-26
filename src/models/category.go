package models

type Category struct {
	ID          uint64 `mapper:"id"`
	Name        string `mapper:"name"`
	Active      bool   `mapper:"active"`
	Description string `mapper:"description"`
}
