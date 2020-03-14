package data

// Normalize normalizes the Rectangle relative to the ScreenDimension
func (rect Rectangle) Normalize(screen ScreenInfo) Rectangle {
	rect.X /= float64(screen.Width)
	rect.Width /= float64(screen.Width)
	rect.Y /= float64(screen.Height)
	rect.Height /= float64(screen.Height)
	return rect
}
