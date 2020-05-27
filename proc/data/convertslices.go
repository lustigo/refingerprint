package data

// ConvertInt16ToFloat64 converts an Int16 Slice to an Float64 Slice
func ConvertInt16ToFloat64(slice []int16) []float64 {
	converted := make([]float64, len(slice))
	for i, val := range slice {
		converted[i] = float64(val)
	}
	return converted
}

// ConvertUint8ToFloat64 converts an Uint8 Slice to an Float64 Slice
func ConvertUint8ToFloat64(slice []uint8) []float64 {
	converted := make([]float64, len(slice))
	for i, val := range slice {
		converted[i] = float64(val)
	}
	return converted
}

// ConvertUint16ToFloat64 converts an Uint16 Slice to an Float64 Slice
func ConvertUint16ToFloat64(slice []uint16) []float64 {
	converted := make([]float64, len(slice))
	for i, val := range slice {
		converted[i] = float64(val)
	}
	return converted
}

// ConvertUint64ToFloat64 converts an Uint64 Slice to an Float64 Slice
func ConvertUint64ToFloat64(slice []uint64) []float64 {
	converted := make([]float64, len(slice))
	for i, val := range slice {
		converted[i] = float64(val)
	}
	return converted
}
