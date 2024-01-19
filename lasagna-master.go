package main

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time <= 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(ingredient []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0
	for i, _ := range ingredient {
		if ingredient[i] == "noodles" {
			noodles++
		} else if ingredient[i] == "sauce" {
			sauce++
		}
	}
	return noodles * 50, sauce * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) []string {
	for i, _ := range myList {
		if myList[i] == "?" {
			myList[i] = friendList[len(friendList)-1]
		}
	}
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	result := make([]float64, len(quantities))
	for i, _ := range quantities {
		result[i] = quantities[i] * float64(portion) / 2
	}
	return result
}
