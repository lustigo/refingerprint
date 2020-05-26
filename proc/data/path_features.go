package data

import (
	"math"

	"github.com/golang/geo/r2"
)

// PathFeatures contains the raw features extracted from a path
type PathFeatures struct {
	PairwiseDistance                 []float64
	DistanceSum                      float64
	DistanceStartEndPoint            float64
	NumberOfMovementPoints           uint16
	PairwiseVelocity                 []float64
	MeanVelocity                     float64
	PairwiseAcceleration             []float64
	MeanAcceleration                 float64
	PairwiseAngle                    []float64
	AngleBetweenMovementAndStartEnd  []float64 // Angle between Movement Point, Start Point and Click Up Point
	PairwiseAngularVelocity          []float64
	AngleStartEndPoint               float64
	PairwiseDuration                 []uint64
	DurationOfPath                   uint64 // Duration from StartClick to EndClick
	TimeBetweenClickAndMovement      uint64
	TimeBetweeenMovementAndDownClick uint64
	TimeBetweenClickAndRelease       []uint64
	NumberOfBreaks                   uint16
	BreakTimes                       []uint64
	BreakTimeTotalTimeRatio          float64
	MovementDuringClickDistance      []float64
	MovementDuringClickVelocity      []float64
	MovementDuringCickAcceleration   []float64
	MovementDuringClickAngle         []float64
	NumberOfRightClicks              uint8
	NumberOfMiddleClicks             uint8
	NumberOfScrolls                  uint8
	ScrollDX                         []int16
	ScrollDY                         []int16
	ScrollDZ                         []int16
	ScrollDM                         []uint8
	XPoints                          []float64
	YPoints                          []float64
	PairwiseXVelocity                []float64
	PairwiseXDistance                []float64
	MeanXVelocity                    float64
	SumXDistance                     float64
	PairwiseYVelocity                []float64
	PairwiseYDistance                []float64
	SumYDistance                     float64
	MeanYVelocity                    float64
	Straightness                     float64
}

// GetRawFeatures calculates the raw PathFeatures from the path
func (path Path) GetRawFeatures() *PathFeatures {
	features := &PathFeatures{}
	features.calculateStartEndFeatures(path.StartClickUp, path.EndClickUp)
	features.calculateMovementFeatures(path.StartClickUp, path.EndClickDown, path.EndClickUp, path.Movements)
	features.calculateScrollFeatures(path.ScrollEvents)
	features.calculateMovementAcceleration()
	features.calculateMovementDuringClicks(append(path.OtherClickEvents, path.EndClickUp, path.EndClickDown))
	features.calculateMovementDuringClickAcceleration()
	features.calculateXYMeanVelocities()
	features.calculateDistanceSum()
	features.calculateStraightness()
	features.calculateNumberOfClicks(path.OtherClickEvents)
	features.storeTimeBetween()
	features.calculateMovementBreaks()
	return features
}

// calculateMovementFeatures calculates and sets the features extracted from the movement data and the start and end click
func (features *PathFeatures) calculateMovementFeatures(startClick, endClickDown, endClickUp NormalizedClickEvent, points []NormalizedMouseData) {
	features.NumberOfMovementPoints = uint16(len(points))
	startPoint := startClick.NormalizedMouseData
	startVector := r2.Point{X: startPoint.X, Y: startPoint.Y}
	endPoint := endClickUp.NormalizedMouseData
	endVector := r2.Point{X: endPoint.X, Y: endPoint.Y}

	i := -1
	for i < len(points) {
		var point NormalizedMouseData
		if i == -1 {
			point = startPoint
		} else {
			point = points[i]
		}
		vector := r2.Point{X: point.X, Y: point.Y}
		features.XPoints = append(features.XPoints, point.X)
		features.YPoints = append(features.YPoints, point.Y)

		var nextPoint NormalizedMouseData
		if i == len(points)-1 {
			nextPoint = endClickDown.NormalizedMouseData
		} else {
			nextPoint = points[i+1]
		}
		nextVector := r2.Point{X: nextPoint.X, Y: nextPoint.Y}
		way := nextVector.Sub(vector)

		distance := way.Norm()
		duration := nextPoint.Time - point.Time
		velocity := distance / float64(duration)
		angle := math.Acos(nextVector.Dot(vector) / (vector.Norm() * nextVector.Norm()))
		angleVelocity := angle / float64(duration)

		features.PairwiseDistance = append(features.PairwiseDistance, distance)
		features.PairwiseDuration = append(features.PairwiseDuration, duration)
		features.PairwiseVelocity = append(features.PairwiseVelocity, velocity)
		features.PairwiseAngle = append(features.PairwiseAngle, angle)
		features.PairwiseAngularVelocity = append(features.PairwiseAngularVelocity, angleVelocity)

		xDistance := math.Abs(nextPoint.X - point.X)
		yDistance := math.Abs(nextPoint.Y - point.Y)
		xVelocity := xDistance / float64(duration)
		yVelocity := yDistance / float64(duration)

		wayToStart := vector.Sub(startVector)
		wayToEnd := endVector.Sub(vector)
		angle = math.Acos((math.Pow(wayToStart.Norm(), 2.0) + math.Pow(wayToEnd.Norm(), 2.0) - math.Pow(features.DistanceStartEndPoint, 2.0)) / (2 * wayToEnd.Norm() * wayToStart.Norm()))
		features.AngleBetweenMovementAndStartEnd = append(features.AngleBetweenMovementAndStartEnd, angle)

		features.PairwiseXDistance = append(features.PairwiseXDistance, xDistance)
		features.PairwiseYDistance = append(features.PairwiseYDistance, yDistance)
		features.PairwiseXVelocity = append(features.PairwiseXVelocity, xVelocity)
		features.PairwiseYVelocity = append(features.PairwiseYVelocity, yVelocity)

		i++
	}
}

// CalculateScrollFeatures calculates and sets the raw features extracted from the ScrollEvents
func (features *PathFeatures) calculateScrollFeatures(scrolls []NormalizedScrollEvent) {
	features.NumberOfScrolls = uint8(len(scrolls))
	for _, scroll := range scrolls {
		features.ScrollDM = append(features.ScrollDM, scroll.DeltaMode)
		features.ScrollDX = append(features.ScrollDX, scroll.DeltaX)
		features.ScrollDY = append(features.ScrollDY, scroll.DeltaY)
		features.ScrollDZ = append(features.ScrollDZ, scroll.DeltaZ)
	}
}

// CalculateStartEndFeatures calculates and sets the raw features extracted from the Start and End Clicks
func (features *PathFeatures) calculateStartEndFeatures(start, end NormalizedClickEvent) {
	startPoint := r2.Point{X: start.X, Y: start.Y}
	endPoint := r2.Point{X: end.X, Y: end.Y}
	way := endPoint.Sub(startPoint)
	angle := math.Acos(endPoint.Dot(startPoint) / (startPoint.Norm() * endPoint.Norm()))

	features.DistanceStartEndPoint = way.Norm()
	features.AngleStartEndPoint = angle
	features.DurationOfPath = end.Time - start.Time
}

// calculateMovementAcceleration calculates and stores the Acceleration of the movement and the mean Velocity and Acceleration
func (features *PathFeatures) calculateMovementAcceleration() {
	if len(features.PairwiseVelocity) > 0 {
		accSum := 0.0
		for i, vel := range features.PairwiseVelocity {
			features.MeanVelocity += vel

			var beforeVel float64
			if i == 0 {
				beforeVel = 0.0
			} else {
				beforeVel = features.PairwiseVelocity[i-1]
			}
			acc := (vel - beforeVel) / float64(features.PairwiseDuration[i])
			accSum += acc
			features.PairwiseAcceleration = append(features.PairwiseAcceleration, acc)
		}

		features.MeanVelocity /= float64(len(features.PairwiseVelocity))
		if len(features.PairwiseAcceleration) > 0 {
			features.MeanAcceleration = accSum / float64(len(features.PairwiseAcceleration))
		}
	}
}

// CalculateXYMeanVelocities calculates and stores the Mean Velocity for the X and Y values
func (features *PathFeatures) calculateXYMeanVelocities() {
	for _, vel := range features.PairwiseXVelocity {
		features.MeanXVelocity += vel
	}
	for _, vel := range features.PairwiseYVelocity {
		features.MeanYVelocity += vel
	}
	if len(features.PairwiseXVelocity) > 1 {
		features.MeanXVelocity /= float64(len(features.PairwiseXVelocity))
	}
	if len(features.PairwiseYVelocity) > 1 {
		features.MeanYVelocity /= float64(len(features.PairwiseYVelocity))
	}
}

// CalculateDistanceSum calculates and stores the Sum of the distances
func (features *PathFeatures) calculateDistanceSum() {
	for _, dist := range features.PairwiseXDistance {
		features.SumXDistance += dist
	}
	for _, dist := range features.PairwiseYDistance {
		features.SumYDistance += dist
	}
	for _, dist := range features.PairwiseDistance {
		features.DistanceSum += dist
	}
}

// CalculateStraightness calculates and stores the straightness of the path
func (features *PathFeatures) calculateStraightness() {
	features.Straightness = features.DistanceSum / features.DistanceStartEndPoint
}

// CalculateNumberOfClicks calculates and stores the number of right and left clicks
func (features *PathFeatures) calculateNumberOfClicks(clicks []NormalizedClickEvent) {
	for _, click := range clicks {
		if click.Release || click.Key == MouseKeyLEFT {
			continue
		}
		if click.Key == MouseKeyRIGHT {
			features.NumberOfRightClicks++
		}
		features.NumberOfMiddleClicks++
	}
}

// CalculateMovementDuringClicks calculates and stores angle, distance and velocity during clicks and Time between click down and up (duration of the click)
func (features *PathFeatures) calculateMovementDuringClicks(clicks []NormalizedClickEvent) {
	for i, click := range clicks {
		if click.Release {
			continue
		}
		key := click.Key

		// Find corresponding click
		j := i + 1
		var correspondingClick NormalizedClickEvent
		for j < len(clicks) {
			if clicks[j].Release && clicks[j].Key == key {
				correspondingClick = clicks[j]
				break
			}
			j++
		}

		if j == len(clicks) {
			continue
		}

		// Calculate values
		downPoint := r2.Point{X: click.X, Y: click.Y}
		upPoint := r2.Point{X: correspondingClick.X, Y: correspondingClick.Y}
		way := upPoint.Sub(downPoint)

		duration := correspondingClick.Time - click.Time
		distance := way.Norm()
		velocity := distance / float64(duration)
		angle := math.Acos(upPoint.Dot(downPoint) / (downPoint.Norm() * upPoint.Norm()))

		// Store Values
		features.MovementDuringClickAngle = append(features.MovementDuringClickAngle, angle)
		features.MovementDuringClickDistance = append(features.MovementDuringClickDistance, distance)
		features.MovementDuringClickVelocity = append(features.MovementDuringClickVelocity, velocity)
		features.TimeBetweenClickAndRelease = append(features.TimeBetweenClickAndRelease, duration)
	}
}

// calculateMovementDuringClickAcceleration calculates and stores the MovementDuringClickAcceleration value
func (features *PathFeatures) calculateMovementDuringClickAcceleration() {
	i := 0

	for i < len(features.MovementDuringClickVelocity) {
		var beforeVel float64
		if i == 0 {
			beforeVel = 0.0
		} else {
			beforeVel = features.MovementDuringClickVelocity[i-1]
		}

		acc := (features.MovementDuringClickVelocity[i] - beforeVel) / float64(features.TimeBetweenClickAndRelease[i])
		features.MovementDuringCickAcceleration = append(features.MovementDuringCickAcceleration, acc)
		i++
	}
}

// storeTimeBetween stores the Time between the startClick and the first Movement Event and the time between the last movement event and the end click down
func (features *PathFeatures) storeTimeBetween() {
	if len(features.PairwiseDuration) > 0 {
		features.TimeBetweenClickAndMovement = features.PairwiseDuration[0]
		features.TimeBetweeenMovementAndDownClick = features.PairwiseDuration[len(features.PairwiseDuration)-1]
	}
}

func (features *PathFeatures) calculateMovementBreaks() {
	var totalBreakTime uint64
	for i, dist := range features.PairwiseDistance {
		if dist == 0 && features.PairwiseDuration[i] > 100 {
			features.NumberOfBreaks++
			features.BreakTimes = append(features.BreakTimes, features.PairwiseDuration[i])
			totalBreakTime += features.PairwiseDuration[i]
		}
	}
	if features.NumberOfBreaks > 0 && features.DurationOfPath > 0 {
		features.BreakTimeTotalTimeRatio = float64(totalBreakTime) / float64(features.DurationOfPath)
	}
}
