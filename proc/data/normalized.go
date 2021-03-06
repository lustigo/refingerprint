package data

// NormalizedMouseData represents the Mouseposition relative to the Screenresolution at a specific moment
type NormalizedMouseData struct {
	X    float64
	Y    float64
	Time uint64
}

// NormalizedScrollEvent represents the ScrollEvent relative to the Screenresolution at a specific moment
type NormalizedScrollEvent struct {
	NormalizedMouseData
	DeltaX    float64
	DeltaY    float64
	DeltaZ    float64
	DeltaMode uint8
}

// NormalizedClickEvent represents the ClickData relative to the Screenresolution at a specific moment
type NormalizedClickEvent struct {
	NormalizedMouseData
	Key     MouseKey
	Release bool // If the Button was released or pushed
}
