package models

type Delivery_package struct {
	id       string
	weight   float64
	quantity int
}

func NewPackage(id string, weight float64, quantity int) *Delivery_package {

	return &Delivery_package{
		id:       id,
		weight:   weight,
		quantity: quantity,
	}

}

func (p *Delivery_package) GetId() string {
	return p.id
}
