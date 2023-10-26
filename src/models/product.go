package models

import (
	"encoding/base64"
	"fmt"
	"time"
)

type Product struct {
	ID       uint64  `mapper:"id"`
	Name     string  `mapper:"name"`
	Price    float64 `mapper:"price"`
	Barcode  string  `mapper:"barcode"`
	Category `mapper:"category"`
}

func (p *Product) GenerateBarcode() {

	barcode := fmt.Sprintf("%d-%s-%d-%s", p.ID, p.Name, p.Category.ID, time.Now().String())

	p.Barcode = base64.StdEncoding.EncodeToString([]byte(barcode))
}
