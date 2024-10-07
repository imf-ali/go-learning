package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var result []float64
	for _, stringVal := range strings{
		floatPrice, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
      return nil, errors.New("failed to convert string to float")
    }
		result = append(result, floatPrice)
	}
	return result, nil
}