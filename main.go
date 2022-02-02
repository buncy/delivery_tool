package main

import (
	"delivery_tool/models"
	"fmt"
)

func main() {

	var (
		base_delivery_cost int
		number_of_packages int
		orders             []*models.Order
	)
	fmt.Println("please enter the base delivery cost and number of packages")
	fmt.Scanf("%d %d", &base_delivery_cost, &number_of_packages)
	for i := 0; i < number_of_packages; i++ {
		fmt.Println("enter the package name, weight, distance and offer code")

		var (
			id       string
			weight   int
			distance int
			offer    string
		)

		fmt.Scanf("%s %d %d %s", &id, &weight, &distance, &offer)

		newPackage := models.NewPackage(id, float64(weight), 1)
		orders = append(orders, models.NewOrder(newPackage, float64(distance), offer, base_delivery_cost))
	}

	fmt.Println("Final cost of delivery is")

	for _, o := range orders {
		fmt.Printf("%s %d %d\n", o.Delivery_package.GetId(), o.GetDiscount(), o.Delivery_cost())
	}

}
