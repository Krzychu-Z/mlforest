package mlforest

import (
	"errors"
)

type XSeries interface {
	[]float64 | [][]float64
}

// Hello returns a greeting for the named person.
func LinearGradientDescent[xSeries XSeries](x_series xSeries, y_series []float64) (bool, error) {
    if len(x_series) != len(y_series) {
		return false, errors.New("Input data size (x series) does not match the expected values data size (y series)")
	}

	// var coefCount int
	var ok bool

	_, ok = any(x_series).([]float64)
	
	return ok, nil
}
