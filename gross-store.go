package main

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	grossStore := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return grossStore
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if val, exists := units[unit]; !exists {
		return false
	} else {
		bill[item] += val
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	valUnits, existsUnits := units[unit]
	valBill, existsBill := bill[item]

	if existsUnits && existsBill && valUnits <= valBill {
		if valUnits == valBill {
			delete(bill, item)
		} else {
			bill[item] -= valUnits
		}
		return true
	}

	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, exists := bill[item]
	return val, exists
}
