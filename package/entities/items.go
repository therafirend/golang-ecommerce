package entities

type Items struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
	Price  int    `db:"price"`
	Stock  int    `db:"stock"`
	Seller int    `db:"id_seller"`
}

func (itm *Items) ToItems() *Items {
	return &Items{
		ID:     itm.ID,
		Name:   itm.Name,
		Price:  itm.Price,
		Stock:  itm.Stock,
		Seller: itm.Seller,
	}
}
