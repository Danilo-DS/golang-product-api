package models

import (
	"encoding/base64"
	"fmt"
	"time"
)

type Product struct {
	ID         uint64   `mapper:"id" gorm:"primaryKey; column:ID"`
	Name       string   `mapper:"name" gorm:"column:NAME"`
	Price      float64  `mapper:"price" gorm:"column:PRICE"`
	Barcode    string   `mapper:"barcode" gorm:"column:BARCODE"`
	CategoryId uint64   `mapper:"categoryId" gorm:"index; column:CATEGORY_ID"`
	Category   Category `mapper:"category" gorm:"foreignKey:ID; references:CategoryId"`
	// Second option
	//CategoryId uint64  `gorm:"uniqueIndex; column:ID_CATEGORY"`
	//Category Category `gorm:"foreignKey:CategoryId"` //This argument in foreignKey reference the attribute CategoryId
}

func (p *Product) GenerateBarcode() {

	barcode := fmt.Sprintf("%d-%s-%d-%s", p.ID, p.Name, p.Category.ID, time.Now().String())

	p.Barcode = base64.StdEncoding.EncodeToString([]byte(barcode))
}

func (p *Product) TableName() string {
	return "PRODUCT"
}
