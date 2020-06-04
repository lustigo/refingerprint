package data

// Normalize normalizes the Rectangle relative to the ScreenDimension
func (rect Rectangle) Normalize(screen ScreenInfo) Rectangle {
	// Convert DocumentDimension to Screendim
	rect.X += float64(screen.DeltaX)
	rect.Y += float64(screen.DeltaY)
	// Normalize
	len := screen.Length()
	rect.X /= len
	rect.Width /= len
	rect.Y /= len
	rect.Height /= len
	return rect
}
