package data

// NormalizedMouseData represents the Mouseposition relative to the Screenresolution at a specific moment
type NormalizedMouseData struct {
	X    float64
	Y    float64
	Time uint64
}
