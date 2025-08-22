package purchase


const Car = "car"
const Truck = "truck"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if (kind == Car || kind == Truck) {
        return true
    }
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    betterChoiceText := " is clearly the better choice."
	if (option1 < option2) {
        return option1 + betterChoiceText
    }
    return option2 + betterChoiceText
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age >= 10 {
        return originalPrice / 100 * 50
    }
    
    if( age >= 3 && age < 10) {
        return originalPrice / 100 * 70
    }
    
    return originalPrice / 100 * 80
}
