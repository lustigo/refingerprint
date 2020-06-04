package gen

import "fmt"

//GenerateRest generates code for the Path Features of the rest
func GenerateRest() {
	featureValueNames := []string{
		"Min",
		"Max",
		"Diff", //Max-Min
		"Sum",
		"Mean",
		"StdDev",
		"Skew", // Third moment
	}
	rawFeatureNames := []string{
		"PairwiseDistance",
		"PairwiseVelocity",
		"PairwiseAcceleration",
		"PairwiseAngle",
		"AngleBetweenMovementAndStartEnd",
		"PairwiseAngularVelocity",
		"PairwiseDuration",
		"TimeBetweenClickAndRelease",
		"BreakTimes",
		"MovementDuringClickDistance",
		"MovementDuringClickVelocity",
		"MovementDuringClickAcceleration",
		"MovementDuringClickAngle",
		"ScrollDX",
		"ScrollDY",
		"ScrollDZ",
		"ScrollDM",
		"XPoints",
		"YPoints",
		"PairwiseXVelocity",
		"PairwiseYVelocity",
		"PairwiseXDistance",
		"PairwiseYDistance",
		"SumYDistance",
		"MeanYVelocity",
		"SumXDistance",
		"MeanXVelocity",
		"Straightness",
		"NumberOfRightClicks",
		"NumberOfMiddleClicks",
		"NumberOfScrolls",
		"BreakTimeTotalTimeRatio",
		"NumberOfBreaks",
		"DurationOfPath",
		"TimeBetweenClickAndMovement",
		"TimeBetweeenMovementAndDownClick",
		"AngleStartEndPoint",
		"MeanAcceleration",
		"MeanVelocity",
		"DistanceStartEndPoint",
		"DistanceSum",
		"NumberOfMovementPoints",
	}
	rawFeatureTypes := []string{
		"[]float64",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]uint64",
		"[]uint64",
		"[]uint64",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]uint8",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]float64",
		"[]float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"uint8",
		"uint8",
		"uint8",
		"float64",
		"uint16",
		"uint64",
		"uint64",
		"uint64",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"uint16",
	}

	generateRestHeaders(rawFeatureNames, featureValueNames)
	generateRestStructFields(rawFeatureNames, featureValueNames, rawFeatureTypes)
	generateRestCalculation(rawFeatureNames, featureValueNames, rawFeatureTypes)
}

// isArray checks if the given type name s is an array
func isArray(s string) bool {
	if len(s) < 2 {
		return false
	}
	if s[:2] == "[]" {
		return true
	}
	return false
}

// getPlainType returns s if it does not start with "[]" and otherwise crops the brackets
func getPlainType(s string) string {
	if isArray(s) {
		return s[2:]
	}
	return s
}

// generateRestHeader generates the Code for the generation of the ARFF Header for the Rest Path Features and prints it
func generateRestHeaders(rawFeatureNames, featureValueNames []string) {
	generateHeaders(rawFeatureNames, featureValueNames, "restpath")
}

// generateRestStructFields generates the Code for ProcessedFeatures struct fields for the Rest Path Features and prints it
func generateRestStructFields(rawFeatureNames, featureValueNames, rawFeatureTypes []string) {
	fmt.Print("//\n// Fields\n// Auto-Generated\n//\n")

	for i, v := range rawFeatureNames {
		fields := generateSingleFields("RestPath", v, getPlainType(rawFeatureTypes[i]), featureValueNames)
		for _, field := range fields {
			fmt.Print(field)
		}
	}

	fmt.Print("\n\n\n")
}

// generateRestCalculation generates the Code for the generation of code for the calculation of the Rest Path Features and prints it
func generateRestCalculation(rawFeatureNames, featureValueNames, rawFeatureTypes []string) {
	fmt.Print("//\n// Calculations\n//\n//\n\n")

	functions := genRestFunctions(rawFeatureNames, rawFeatureTypes)
	genRestStruct(rawFeatureNames, rawFeatureTypes)
	fmt.Print("// GenerateRestCalculation caluclates and stores the Features of the RestPath; Auto-generated\nfunc (features *ProcessedFeatures) GenerateRestCalculation (paths []*PathFeatures){\n")
	fmt.Print("\tcollected := &collection{}\n")
	genCollect(rawFeatureNames, rawFeatureTypes, "collected", "paths")
	for _, fun := range functions {
		fmt.Printf("\tfeatures.%v(collected)\n", fun)
	}
	fmt.Print("}\n")

	fmt.Print("\n\n\n")
}

// genRestStruct generates a temporary struct for the calculation of the Rest Path Features and prints the code for it
func genRestStruct(rawFeatureNames, rawFeatureTypes []string) {
	fmt.Print("//collection is a struct to collect all the paths; Auto-generated\ntype collection struct {\n")
	for i, feature := range rawFeatureNames {
		typ := rawFeatureTypes[i]
		if isArray(typ) {
			fmt.Printf("\t%v\t%v\n", feature, typ)
		} else {
			fmt.Printf("\t%v\t[]%v\n", feature, typ)
		}
	}
	fmt.Print("}\n")
}

// genCollect generates the code for the collection of all path data into one and prints it
func genCollect(rawFeatureNames, rawFeatureTypes []string, collectionName, pathsName string) {
	fmt.Printf("\tfor _, v := range %v {\n", pathsName)
	for i, v := range rawFeatureNames {
		if isArray(rawFeatureTypes[i]) {
			fmt.Printf("\t\t%v.%v = append(%v.%v, v.%v...)\n", collectionName, v, collectionName, v, v)
		} else {
			fmt.Printf("\t\t%v.%v = append(%v.%v, v.%v)\n", collectionName, v, collectionName, v, v)
		}
	}
	fmt.Print("\t}\n")
}

// genRestFunctions generates the code for the functions for the calculation of the restpath features and returns the generated feature names
func genRestFunctions(rawFeatureNames, rawFeatureTypes []string) []string {
	// Declarations
	minFunctions := map[string]string{
		"int16":  "MinInt",
		"uint8":  "MinUint8",
		"uint16": "MinUint16",
		"uint64": "MinUint64",
	}
	maxFunctions := map[string]string{
		"int16":  "MaxInt",
		"uint8":  "MaxUint8",
		"uint16": "MaxUint16",
		"uint64": "MaxUint64",
	}
	sumFunctions := map[string]string{
		"int16":   "SumInt",
		"uint8":   "SumUint8",
		"uint16":  "SumUint16",
		"uint64":  "SumUint64",
		"float64": "SumFloat",
	}
	convertFunctions := map[string]string{
		"int16":  "ConvertInt16ToFloat64",
		"uint8":  "ConvertUint8ToFloat64",
		"uint16": "ConvertUint16ToFloat64",
		"uint64": "ConvertUint64ToFloat64",
	}

	funcNames := make([]string, 0)

	for i, name := range rawFeatureNames {
		typ := getPlainType(rawFeatureTypes[i])
		newName := "RestPath" + name
		funcName := "calc" + newName + "Features"
		fmt.Print("// " + funcName + " calculates the features of the " + name + " vector from the given path; Auto-generated\nfunc (features *ProcessedFeatures) " + funcName + " (col *collection) {\n")
		funcNames = append(funcNames, funcName)

		if typ == "float64" {
			fmt.Print("\tfeatures." + newName + "Min, features." + newName + "Mean, features." + newName + "Max, features." + newName + "StdDev = rnd.StatBasic(col." + name + ", true) \n")
			fmt.Print("\tfeatures." + newName + "Sum = " + sumFunctions[typ] + "(col." + name + ")\n")
			fmt.Print("\tif len(col." + name + ") == 1 {\n\t\tfeatures." + newName + "Min, features." + newName + "Mean, features." + newName + "Max = col." + name + "[0], col." + name + "[0], col." + name + "[0]\n\t}\n")
			fmt.Print("\tfeatures." + newName + "Diff = features." + newName + "Max - features." + newName + "Min \n")
			fmt.Print("\tdefer func(){\n\t\trecover()\n\t}()\n")
			fmt.Print("\t_, _, _, _, _, features." + newName + "Skew ,_ = rnd.StatMoments(col." + name + ") \n")
		} else {
			fmt.Print("\tfeatures." + newName + "Max = " + maxFunctions[typ] + "(col." + name + ")\n")
			fmt.Print("\tfeatures." + newName + "Min = " + minFunctions[typ] + "(col." + name + ")\n")
			fmt.Print("\tfeatures." + newName + "Diff = features." + newName + "Max - features." + newName + "Min\n")
			fmt.Print("\tfeatures." + newName + "Sum = " + sumFunctions[typ] + "(col." + name + ")\n")
			fmt.Print("\tfeatures." + newName + "Mean = Mean(" + convertFunctions[typ] + "(col." + name + "),float64(features." + newName + "Sum))\n")
			fmt.Print("\tdefer func(){\n\t\trecover()\n\t}()\n")
			fmt.Print("\t_,_, _, _, features." + newName + "StdDev, features." + newName + "Skew ,_ = rnd.StatMoments(" + convertFunctions[typ] + "(col." + name + ")) \n")

		}
		fmt.Print("}\n")
	}
	return funcNames
}
