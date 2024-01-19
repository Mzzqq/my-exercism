package main

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time <= 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

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
