package data

import (
	"sort"

	"github.com/golang/geo/r2"
)

// Equals returns true if the two points are the same
func (point NormalizedMouseData) Equals(otherPoint NormalizedMouseData) bool {
	return point.X == otherPoint.X && point.Y == otherPoint.Y && point.Time == otherPoint.Time
}

// Equals returns true if the two events are the same
func (event NormalizedScrollEvent) Equals(otherEvent NormalizedScrollEvent) bool {
	return event.NormalizedMouseData.Equals(otherEvent.NormalizedMouseData) && event.DeltaMode == otherEvent.DeltaMode && event.DeltaX == otherEvent.DeltaX && event.DeltaY == otherEvent.DeltaY && event.DeltaZ == otherEvent.DeltaZ
}

// Equals returns true if the two events are the same
func (event NormalizedClickEvent) Equals(otherEvent NormalizedClickEvent) bool {
	return event.NormalizedMouseData.Equals(otherEvent.NormalizedMouseData) && event.Key == otherEvent.Key && event.Release == otherEvent.Release
}

// Normalize normalizes the MouseData with the given Screenresolution and Time
func (point MouseData) Normalize(screen ScreenInfo, time Time) NormalizedMouseData {
	len := screen.Length()
	return NormalizedMouseData{
		X:    float64(point.X) / len,
		Y:    float64(point.Y) / len,
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

// NormalizeScrollEvents normalizes all the ScrollData with the given Screenresolution and Time
func NormalizeScrollEvents(events []ScrollData, screen ScreenInfo, time Time) []NormalizedScrollEvent {
	normalized := make([]NormalizedScrollEvent, len(events))
	for i, event := range events {
		normalized[i] = NormalizedScrollEvent{
			event.Normalize(screen, time),
			event.DeltaX,
			event.DeltaY,
			event.DeltaZ,
			event.DeltaMode,
		}
	}
	return normalized
}

// NormalizeClickEvents normalizes all the ClickData with the given Screenresolution and Time
func NormalizeClickEvents(events []ClickData, screen ScreenInfo, time Time) []NormalizedClickEvent {
	normalized := make([]NormalizedClickEvent, len(events))
	for i, event := range events {
		normalized[i] = NormalizedClickEvent{
			event.Normalize(screen, time),
			event.Key,
			event.Release,
		}
	}
	return normalized
}

// RemoveDuplicateMousePoints removes identical events that are next to each other or events that have the same timestamp
func RemoveDuplicateMousePoints(points []NormalizedMouseData) []NormalizedMouseData {
	if len(points) == 0 {
		return points
	}
	// Remove Duplicates
	cleaned := make([]NormalizedMouseData, 1)
	cleaned[0] = points[0]
	i := 0
	for i < len(points)-1 {
		if !points[i].Equals(points[i+1]) {
			cleaned = append(cleaned, points[i+1])
		}
		i++
	}

	// Remove events with same time
	toRemove := make([]int, 0)
	for i := range cleaned {
		j := i + 1
		for j < len(cleaned) {
			if cleaned[i].Time == cleaned[j].Time {
				toRemove = append(toRemove, j)
			}
			j++
		}
	}

	if len(toRemove) > 0 {
		sort.Ints(toRemove)
		timeCleaned := make([]NormalizedMouseData, len(cleaned)-len(toRemove))
		cleanedCounter := 0
		toRemoveCounter := 0
		timeCleanedCounter := 0
		for cleanedCounter < len(cleaned) {
			if toRemove[toRemoveCounter] == cleanedCounter {
				toRemoveCounter++
			} else {
				timeCleaned[timeCleanedCounter] = cleaned[cleanedCounter]
				timeCleanedCounter++
			}
			cleanedCounter++
		}
		return timeCleaned
	}

	return cleaned
}

// RemoveDuplicateScrollEvents removes identical events that are next to each other or events that have the same timestamp
func RemoveDuplicateScrollEvents(events []NormalizedScrollEvent) []NormalizedScrollEvent {
	if len(events) == 0 {
		return events
	}
	// Remove Duplicates
	cleaned := make([]NormalizedScrollEvent, 1)
	cleaned[0] = events[0]
	i := 0
	for i < len(events)-1 {
		if !events[i].Equals(events[i+1]) {
			cleaned = append(cleaned, events[i+1])
		}
		i++
	}

	// Remove events with same time
	toRemove := make([]int, 0)
	for i := range cleaned {
		j := i + 1
		for j < len(cleaned) {
			if cleaned[i].Time == cleaned[j].Time {
				toRemove = append(toRemove, j)
			}
			j++
		}
	}

	if len(toRemove) > 0 {
		sort.Ints(toRemove)
		timeCleaned := make([]NormalizedScrollEvent, len(cleaned)-len(toRemove))
		cleanedCounter := 0
		toRemoveCounter := 0
		timeCleanedCounter := 0
		for cleanedCounter < len(cleaned) {
			if toRemove[toRemoveCounter] == cleanedCounter {
				toRemoveCounter++
			} else {
				timeCleaned[timeCleanedCounter] = cleaned[cleanedCounter]
				timeCleanedCounter++
			}
			cleanedCounter++
		}
		return timeCleaned
	}

	return cleaned
}

// RemoveDuplicateClickEvents removes identical events that are next to each other or events that have the same timestamp
func RemoveDuplicateClickEvents(events []NormalizedClickEvent) []NormalizedClickEvent {
	if len(events) == 0 {
		return events
	}
	// Remove Duplicates
	cleaned := make([]NormalizedClickEvent, 1)
	cleaned[0] = events[0]
	i := 0
	for i < len(events)-1 {
		if !events[i].Equals(events[i+1]) {
			cleaned = append(cleaned, events[i+1])
		}
		i++
	}

	// Remove events with same time
	toRemove := make([]int, 0)
	for i := range cleaned {
		j := i + 1
		for j < len(cleaned) {
			if cleaned[i].Time == cleaned[j].Time {
				toRemove = append(toRemove, j)
			}
			j++
		}
	}

	if len(toRemove) > 0 {
		sort.Ints(toRemove)
		timeCleaned := make([]NormalizedClickEvent, len(cleaned)-len(toRemove))
		cleanedCounter := 0
		toRemoveCounter := 0
		timeCleanedCounter := 0
		for cleanedCounter < len(cleaned) {
			if toRemove[toRemoveCounter] == cleanedCounter {
				toRemoveCounter++
			} else {
				timeCleaned[timeCleanedCounter] = cleaned[cleanedCounter]
				timeCleanedCounter++
			}
			cleanedCounter++
		}
		return timeCleaned
	}

	return cleaned
}

// RemoveConcurrentEvents removes events which happened at the same time. ClickEvents have priority before MouseMovements before ScrollEvents
func RemoveConcurrentEvents(clicks []NormalizedClickEvent, points []NormalizedMouseData, scrolls []NormalizedScrollEvent) ([]NormalizedMouseData, []NormalizedScrollEvent) {
	// Check which one to remove
	pointsToRemove := make([]int, 0)
	scrollsToRemove := make([]int, 0)

	for _, click := range clicks {
		for i, point := range points {
			if point.Equals(click.NormalizedMouseData) {
				pointsToRemove = append(pointsToRemove, i)
			}
		}
		for i, scroll := range scrolls {
			if scroll.NormalizedMouseData.Equals(click.NormalizedMouseData) {
				scrollsToRemove = append(scrollsToRemove, i)
			}
		}
	}

	sort.Ints(pointsToRemove)
	sort.Ints(scrollsToRemove)
	pointRemoveCounter := 0
	scrollRemoveCounter := 0

	for i, point := range points {
		if len(pointsToRemove) > pointRemoveCounter && pointsToRemove[pointRemoveCounter] == i {
			pointRemoveCounter++
			continue
		}
		scrollRemoveCounter = 0
		for j, scroll := range scrolls {
			if len(scrollsToRemove) > scrollRemoveCounter && scrollsToRemove[scrollRemoveCounter] == j {
				scrollRemoveCounter++
				continue
			}
			if point.Equals(scroll.NormalizedMouseData) {
				scrollsToRemove = append(scrollsToRemove, j)
			}
		}
	}

	sort.Ints(scrollsToRemove)
	// Remove
	var cleanedPoints []NormalizedMouseData
	var cleanedScrolls []NormalizedScrollEvent

	if len(pointsToRemove) > 0 {
		cleanedPoints = make([]NormalizedMouseData, len(points)-len(pointsToRemove))
		rawCounter := 0
		toRemoveCounter := 0
		cleanedCounter := 0
		for rawCounter < len(points) {
			if toRemoveCounter == len(pointsToRemove) || rawCounter != pointsToRemove[toRemoveCounter] {
				cleanedPoints[cleanedCounter] = points[rawCounter]
				cleanedCounter++
			} else {
				toRemoveCounter++
			}
			rawCounter++
		}
	} else {
		cleanedPoints = points
	}
	if len(scrollsToRemove) > 0 {
		cleanedScrolls = make([]NormalizedScrollEvent, len(scrolls)-len(scrollsToRemove))
		rawCounter := 0
		toRemoveCounter := 0
		cleanedCounter := 0
		for rawCounter < len(scrolls) {
			if toRemoveCounter == len(scrollsToRemove) || rawCounter != scrollsToRemove[toRemoveCounter] {
				cleanedScrolls[cleanedCounter] = scrolls[rawCounter]
				cleanedCounter++
			} else {
				toRemoveCounter++
			}
			rawCounter++
		}
	} else {
		cleanedScrolls = scrolls
	}

	return cleanedPoints, cleanedScrolls
}

// Length returns the vector length of the screen
func (screen ScreenInfo) Length() float64 {
	return r2.Point{X: float64(screen.Width), Y: float64(screen.Height)}.Norm()
}
