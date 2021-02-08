package internal

type Item struct {
	ID            string
	Title         string
	OriginalPrice float64
	Price         float64
}

type SalePrice struct {
	RegularPrice float64
	Price        float64
}
