package models

type Category struct {
	ID          uint64 `mapper:"id" gorm:"primaryKey; column:ID"`
	Name        string `mapper:"name" gorm:"column:NAME"`
	Active      bool   `mapper:"active" gorm:"column:ACTIVE"`
	Description string `mapper:"description" gorm:"column:DESCRIPTION"`
}

func (c *Category) TableName() string {
	return "CATEGORY"
}
