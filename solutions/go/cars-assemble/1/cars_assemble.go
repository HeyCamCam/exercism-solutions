package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    //var x float64 = successRate / 100
    //answer := float64(productionRate) * x
    WorkingCarsPerHour := float64(productionRate) * (successRate / 100)
	return WorkingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    WorkingCarsPerHour := float64(productionRate) * (successRate / 100)
	return int(WorkingCarsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var groupsOfTenCars int = carsCount / 10
    var remainderCars int = carsCount % 10 // Gets the divisor of carsCount
    cost := (groupsOfTenCars * 95000) + (remainderCars * 10000)
	return uint(cost)
}
