package myfunction

import (
	"fmt"
	"math"
)

func DataTrafficSizeCalculator(a *float64, b *float64) (*string, error) {
	var NullString string = ""
	if a == nil || b == nil {
		return &NullString, fmt.Errorf("input values cannot be nil")
	} else if *a < 0 || *b < 0 {
		return &NullString, fmt.Errorf("input values must be non-negative")
	} else {
		var Traficsize float64 = math.Log(*a) * math.Sin(*b)
		var TrafficsizeString string = "the Trafic size is " + fmt.Sprintf("%f", Traficsize)
		return &TrafficsizeString, nil
	}
}
