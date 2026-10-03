package exercise

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, estimateTime int) int {
	// count := make(map[string]int)
 //    for _,layer := range layers {
 //        count[layer] += 1
 //    }
    if estimateTime == 0 {
        return len(layers) * 2
    }
    return len(layers) * estimateTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodle int
	var sauce float64
	for _, layer := range layers {
		if layer == "noodles" {
			noodle += 50
		}else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodle, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) []string {
	secretIngredient := friendList[len(friendList)-1]
	myList[len(myList)-1] = secretIngredient
	return myList
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scale := float64(portions) / 2
    scales :=  make([]float64, len(quantities))
	for i, quantity := range quantities {
		scales[i] = quantity*scale
	}
	return scales
}

fmt.Println(
)