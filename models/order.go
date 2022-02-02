package models

type Order struct {
	Delivery_package *Delivery_package
	distance         float64
	offer_code       string
	baseDeliveryCost int
}

func NewOrder(newPackage *Delivery_package, distance float64, offer_code string, baseDeliveryCost int) *Order {
	return &Order{
		Delivery_package: newPackage,
		distance:         distance,
		offer_code:       offer_code,
		baseDeliveryCost: baseDeliveryCost,
	}
}

func (order *Order) Delivery_cost() int {

	cost_after_discount := order.GetStandardDeliveryCost() - order.GetDiscount()

	return cost_after_discount
	//return 0
}

func (order *Order) GetDiscount() int {

	//if offer is valid apply discount else dont
	if validateOffer(order.offer_code, int(order.distance), int(order.Delivery_package.weight)) {

		offer := getOffer(order.offer_code)

		discount := order.GetStandardDeliveryCost() * offer.discount / 100
		return discount
		//return 0
	} else {
		return 0
	}

}

func (order *Order) GetStandardDeliveryCost() int {

	delivery_cost := order.baseDeliveryCost + int(order.Delivery_package.weight)*10 + int(order.distance)*5

	return delivery_cost
	//return 0
}
