package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
    if (prepTime == 0) {
        return len(layers) * 2
    }
    return len(layers) * prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0
    for i := 0; i < len(layers); i++ {
        if (layers[i] == "noodles") {
            noodles += 50
        }
        if (layers[i] == "sauce") {
            sauce += 0.2
        }
    }
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
    if (len(myList) > 0 && len(friendList) > 0) {
    	myList[len(myList) - 1] = friendList[len(friendList) - 1]
    }
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, noOfPortions int) []float64 {
    newQuans := make([]float64, len(quantities))
    
    scale := float64(noOfPortions) / 2.0
    
    for i := 0; i < len(newQuans); i++ {
       newQuans[i] = quantities[i] * scale
    }
    return newQuans
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
