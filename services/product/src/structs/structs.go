package structs

import "time"

// Product representation
type Product struct {
	ID       uint64 `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Sku      string `json:"sku" db:"sku"`
	Price    int32  `json:"price_in_cents" db:"price_in_cents"`
	Discount `json:"discount"`
}

// Discount representation
type Discount struct {
	DiscountRate  string `json:"discount_rate,omitempty"`
	DiscountPrice int32  `json:"product_price_in_cents,omitempty"`
}

// User respresentation
type User struct {
	ID       string    `db:"id"`
	Name     string    `db:"name"`
	Birthday time.Time `db:"birthday"`
}
