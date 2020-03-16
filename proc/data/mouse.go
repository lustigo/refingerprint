package data

// Normalize normalizes the MouseData with the given Screenresolution
func (point MouseData) Normalize(screen ScreenInfo) NormalizedMouseData {
	return NormalizedMouseData{
		X:    float64(point.X) / float64(screen.Width),
		Y:    float64(point.Y) / float64(screen.Height),
		Time: point.Time,
	}
}

// NormalizeMouseData normalizes all the MouseData with the given Screenresolution
func NormalizeMouseData(points []MouseData, screen ScreenInfo) []NormalizedMouseData {
	normalized := make([]NormalizedMouseData, len(points))
	for i, point := range points {
		normalized[i] = point.Normalize(screen)
	}
	return normalized
}
