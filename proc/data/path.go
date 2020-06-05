package data

// Path describes a MouseMovement between two Left-Button Clicks
type Path struct {
	StartClickUp     NormalizedClickEvent
	EndClickDown     NormalizedClickEvent
	EndClickUp       NormalizedClickEvent
	ScrollEvents     []NormalizedScrollEvent
	OtherClickEvents []NormalizedClickEvent
	Movements        []NormalizedMouseData
}

// ExtractPaths extracts the Path until the Checkboxclick and the rest of the paths from the given sorted and cleaned data
func ExtractPaths(checkBoxMovement, restMovement []NormalizedMouseData, clicks []NormalizedClickEvent, scrolls []NormalizedScrollEvent) (Path, []Path) {
	// Loop Counters
	clickMovementCounter := 0
	restMovementCounter := 0
	restCounter := 0
	scrollCounter := 0

	//
	// Get the Checkbox Path, it ends with the first click
	//
	checkboxPath := Path{
		OtherClickEvents: make([]NormalizedClickEvent, 0),
		Movements:        checkBoxMovement,
		ScrollEvents:     make([]NormalizedScrollEvent, 0),
	}
	// Get the Click Events
	for clickMovementCounter < len(clicks) {
		if click := clicks[clickMovementCounter]; click.Key == MouseKeyLEFT {
			if click.Release {
				checkboxPath.EndClickUp = click
				clickMovementCounter++
				break
			} else {
				checkboxPath.EndClickDown = click
			}
		} else {
			checkboxPath.OtherClickEvents = append(checkboxPath.OtherClickEvents, click)
		}
		clickMovementCounter++
	}
	// Get the ScrollEvents
	for scrollCounter < len(scrolls) {
		if scrolls[scrollCounter].Time > checkboxPath.EndClickUp.Time {
			break
		}
		checkboxPath.ScrollEvents = append(checkboxPath.ScrollEvents, scrolls[scrollCounter])
		scrollCounter++
	}

	//
	// Rest Paths
	//
	restPathAmount := getAmountOfPaths(clicks) - 1
	if restPathAmount < 0 {
		restPathAmount = 0
	}
	restPaths := make([]Path, restPathAmount)
	for restCounter < restPathAmount {
		restPaths[restCounter] = Path{
			OtherClickEvents: make([]NormalizedClickEvent, 0),
			Movements:        make([]NormalizedMouseData, 0),
			ScrollEvents:     make([]NormalizedScrollEvent, 0),
		}
		// Start Click
		if restCounter == 0 {
			restPaths[restCounter].StartClickUp = checkboxPath.EndClickUp
		} else {
			restPaths[restCounter].StartClickUp = restPaths[restCounter-1].EndClickUp
		}
		// End Click and other Clicks
		for clickMovementCounter < len(clicks) {
			if click := clicks[clickMovementCounter]; click.Key == MouseKeyLEFT {
				if click.Release {
					restPaths[restCounter].EndClickUp = click
					clickMovementCounter++
					break
				} else {
					restPaths[restCounter].EndClickDown = click
				}
			} else {
				restPaths[restCounter].OtherClickEvents = append(restPaths[restCounter].OtherClickEvents, click)
			}
			clickMovementCounter++
		}
		// Movement Events
		for restMovementCounter < len(restMovement) {
			if restMovement[restMovementCounter].Time > restPaths[restCounter].EndClickUp.Time {
				break
			}
			restPaths[restCounter].Movements = append(restPaths[restCounter].Movements, restMovement[restMovementCounter])
			restMovementCounter++
		}
		// Scroll Events
		for scrollCounter < len(scrolls) {
			if scrolls[scrollCounter].Time > restPaths[restCounter].EndClickUp.Time {
				break
			}
			restPaths[restCounter].ScrollEvents = append(restPaths[restCounter].ScrollEvents, scrolls[scrollCounter])
			scrollCounter++
		}

		restCounter++
	}

	return checkboxPath, restPaths
}

// GetAmountOfPaths calculates the amount of paths
func getAmountOfPaths(clicks []NormalizedClickEvent) int {
	paths := 0
	for _, click := range clicks {
		if click.Key == MouseKeyLEFT && click.Release {
			paths++
		}
	}
	return paths
}

// RemoveSmallPaths removes every path that has less than 4 movement events
func RemoveSmallPaths(paths []Path) []Path {
	indexToRemove := make([]int, 0)
	for i, path := range paths {
		if len(path.Movements) < 4 {
			indexToRemove = append(indexToRemove, i)
		}
	}

	if len(indexToRemove) == 0 {
		return paths
	}

	cleanedPaths := make([]Path, len(paths)-len(indexToRemove))
	newCounter := 0
	removeCounter := 0
	for i, path := range paths {
		if removeCounter < len(indexToRemove) && i == indexToRemove[removeCounter] {
			removeCounter++
		} else {
			cleanedPaths[newCounter] = path
			newCounter++
		}
	}
	return cleanedPaths
}
