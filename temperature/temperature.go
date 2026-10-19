package temperature

import (
	"fmt"
	"strconv"
)

func ParseCelsius(raw string) (int16, error) {
	num, err := strconv.ParseInt(raw, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("parse %q as int16: %w", raw, err)
	}
	return int16(num), nil
}

func ToFahrenheit(c int16) float64 {
	return float64(c)*9/5 + 32
}
