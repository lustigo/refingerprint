package data

import (
	"sort"
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
