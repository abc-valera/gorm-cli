package samepackage

import "time"

// Example query interfaces with raw SQL annotations

// IProductQuery demonstrates raw SQL query generation
type IProductQuery interface {
	// SELECT * FROM @@table WHERE price > @minPrice
	FindExpensiveProducts(minPrice float64) ([]Product, error)

	// UPDATE @@table
	//  {{set}}
	//    {{if discountFactor > 0}} price = price * @discountFactor, {{end}}
	//    {{if !beforeDate.IsZero()}} updated_at = NOW() {{end}}
	//  {{end}}
	// WHERE created_at < @beforeDate
	DiscountOldProducts(discountFactor float64, beforeDate time.Time) error
}
