package models

var offers = map[string]*Offer{
	"OFR001": {
		id: "OFR001",
		distance: &limit{
			min: 0,
			max: 199,
		},
		weight: &limit{
			min: 70,
			max: 200,
		},
		discount: 10,
	},
	"OFR002": {
		id: "OFR001",
		distance: &limit{
			min: 50,
			max: 150,
		},
		weight: &limit{
			min: 100,
			max: 250,
		},
		discount: 7,
	},
	"OFR003": {
		id: "OFR001",
		distance: &limit{
			min: 50,
			max: 250,
		},
		weight: &limit{
			min: 10,
			max: 150,
		},
		discount: 5,
	},
}

type Offer struct {
	id       string
	distance *limit
	weight   *limit
	discount int
}

type limit struct {
	min int
	max int
}

func NewOffer(id string, minDistance, maxDisatnce, minWeight, maxWeight int) *Offer {
	newOffer := &Offer{
		id: id,
		distance: &limit{
			min: minDistance,
			max: maxDisatnce,
		},
		weight: &limit{
			min: minWeight,
			max: maxWeight,
		},
	}
	offers[newOffer.id] = newOffer
	return newOffer
}

func validateOffer(id string, distance, weight int) bool {

	offer := offers[id]

	if offer.distance.min <= distance && offer.distance.max >= distance && offer.weight.min <= weight && offer.weight.max >= weight {
		return true
	} else {
		return false
	}

}

func getOffer(id string) *Offer {
	return offers[id]
}
