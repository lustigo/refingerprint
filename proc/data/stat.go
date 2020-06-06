package data

// MaxInt returns the Maximum number contained in the given slice
func MaxInt(slice []int16) int16 {
	l := len(slice)
	if l == 0 {
		return 0
	} else if l == 1 {
		return slice[1]
	}

	maxVal := int16(-32768)
	for _, v := range slice {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// MaxUint8 returns the Maximum number contained in the given slice
func MaxUint8(slice []uint8) uint8 {
	l := len(slice)
	if l == 0 {
		return 0
	} else if l == 1 {
		return slice[1]
	}

	maxVal := uint8(0)
	for _, v := range slice {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// MaxUint16 returns the Maximum number contained in the given slice
func MaxUint16(slice []uint16) uint16 {
	l := len(slice)
	if l == 0 {
		return 0
	} else if l == 1 {
		return slice[1]
	}

	maxVal := uint16(0)
	for _, v := range slice {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// MaxUint64 returns the Maximum number contained in the given slice
func MaxUint64(slice []uint64) uint64 {
	l := len(slice)
	if l == 0 {
		return 0
	} else if l == 1 {
		return slice[1]
	}

	maxVal := uint64(0)
	for _, v := range slice {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// MinInt returns the Minimum number contained in the given slice
func MinInt(slice []int16) int16 {
	l := len(slice)
	if l == 0 {
		return 0
	} else if l == 1 {
		return slice[1]
	}

	minVal := int16(32767)
	for _, v := range slice {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

// MinUint8 returns the Minimum number contained in the given slice
func MinUint8(slice []uint8) uint8 {
	l := len(slice)
	if l == 0 {
		return 0
	} else if l == 1 {
		return slice[1]
	}

	minVal := uint8(255)
	for _, v := range slice {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

// MinUint16 returns the Minimum number contained in the given slice
func MinUint16(slice []uint16) uint16 {
	l := len(slice)
	if l == 0 {
		return 0
	} else if l == 1 {
		return slice[1]
	}

	minVal := uint16(65535)
	for _, v := range slice {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

// MinUint64 returns the Minimum number contained in the given slice
func MinUint64(slice []uint64) uint64 {
	l := len(slice)
	if l == 0 {
		return 0
	} else if l == 1 {
		return slice[1]
	}

	minVal := uint64(18446744073709551615)
	for _, v := range slice {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

// SumFloat returns the Sum of the given slice
func SumFloat(slice []float64) float64 {
	sum := 0.0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumInt returns the Sum of the given slice
func SumInt(slice []int16) int16 {
	sum := int16(0)
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumUint8 returns the Sum of the given slice
func SumUint8(slice []uint8) uint8 {
	sum := uint8(0)
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumUint16 returns the Sum of the given slice
func SumUint16(slice []uint16) uint16 {
	sum := uint16(0)
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumUint64 returns the Sum of the given slice
func SumUint64(slice []uint64) uint64 {
	sum := uint64(0)
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Mean calculates the Mean Value of the given slice
func Mean(slice []float64, sum float64) float64 {
	if len(slice) == 0 {
		return 0.0
	}
	return sum / float64(len(slice))
}
