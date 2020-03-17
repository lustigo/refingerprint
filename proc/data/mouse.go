package data

// Normalize normalizes the MouseData with the given Screenresolution and Time
func (point MouseData) Normalize(screen ScreenInfo, time Time) NormalizedMouseData {
	return NormalizedMouseData{
		X:    float64(point.X) / float64(screen.Width),
		Y:    float64(point.Y) / float64(screen.Height),
		Time: point.Time - time.Start,
	}
}

// NormalizeMouseData normalizes all the MouseData with the given Screenresolution and Time
func NormalizeMouseData(points []MouseData, screen ScreenInfo, time Time) []NormalizedMouseData {
	normalized := make([]NormalizedMouseData, len(points))
	for i, point := range points {
		normalized[i] = point.Normalize(screen, time)
	}
	return normalized
}
