package main
/*
import "math"
import "fmt"
*/
// OR
import (
	"fmt"
	"math"
)
func main() {
/*
	var investmentAmount float64 = 1000
    its a better way to declare them beforehand 
	var expectedReturnRate = 5.5
	expectedreturnRate := 5.5
	var years float64 = 10 */

	const inflationRate float64 = 2
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	// After adding scan it waits for input .
	// Its interesting because we overwritten it in this way.
	fmt.Print("Invesment Amount:")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years:")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate:")
	fmt.Scan(&expectedReturnRate)
	
	//var investmentAmount, years float64= 100, 10
    // or transform during the calculation with the float64(investmentamoun)
	futureValue := float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100,float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100,years)
	fmt.Println("Future Value",futureValue)
	fmt.Println("Future Real Value",futureRealValue)
}
