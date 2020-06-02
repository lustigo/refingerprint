// Package gen generates code which will be used in the data package
package gen

import (
	"fmt"
	"strings"
)

//GenerateCheckbox generates code for the checkboxPath Features
func GenerateCheckbox() {
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
	}
	rawFeatureTypes := []string{
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"uint64",
		"uint64",
		"uint64",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"uint8",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
		"float64",
	}
	generateCheckboxStructFields(rawFeatureNames, rawFeatureTypes, featureValueNames)
	generateCheckboxHeader(rawFeatureNames, featureValueNames)
	generateCheckboxCalculation(rawFeatureNames, rawFeatureTypes)
}

// generateCheckboxHeader generates the Code for the generation of the ARFF Header for the CheckboxPath Features and prints it
func generateCheckboxHeader(rawFeatureNames, featureValueNames []string) {
	fmt.Print("//\n// Headers\n// Auto-Generated\n//\n")
	for _, v := range rawFeatureNames {
		for _, feature := range featureValueNames {
			fmt.Print("header.AddAttr(\"checkboxpath" + v + feature + "\", arff.Numeric, nil)\n")
		}
	}

	fmt.Print("\n\n\n")
}

// generateCheckboxStructFields generates the Code for ProcessedFeatures struct fields for the CheckboxPath Features and prints it
func generateCheckboxStructFields(rawFeatureNames, rawFeatureTypes, featureValueNames []string) {
	fmt.Print("//\n// Fields\n// Auto-Generated\n//\n")

	for i, v := range rawFeatureNames {
		fields := generateSingleFields("CheckBoxPath", v, rawFeatureTypes[i], featureValueNames)
		for _, field := range fields {
			fmt.Print(field)
		}
	}

	fmt.Print("\n\n\n")
}

// generateSingleFields generates one Struct Field for the given Feature
func generateSingleFields(prefix, name, typ string, featureValueNames []string) []string {
	fields := make([]string, len(featureValueNames))
	for i, v := range featureValueNames {
		lowerPrefix := strings.ToLower(prefix)
		newName := prefix + name + v
		tag := "`arff:\"" + lowerPrefix + name + v + "\" csv:\"" + lowerPrefix + name + v + "\"`"
		if i > 3 {
			fields[i] = newName + "\t" + "float64 " + tag + "\n"
		} else {
			fields[i] = newName + "\t" + typ + " " + tag + "\n"
		}
	}
	return fields
}

// generateCheckboxCalculation generates the Code for the generation of code for the calculation of the CheckboxPath Features and prints it
func generateCheckboxCalculation(rawFeatureNames, rawFeatureTypes []string) {
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

	// Logic
	fmt.Print("//\n// Calculations\n//\n//\n\n")
	funcNames := make([]string, 0)

	for i, name := range rawFeatureNames {
		typ := rawFeatureTypes[i]
		newName := "CheckBoxPath" + name
		funcName := "calc" + newName + "Features"
		fmt.Print("// " + funcName + " calculates the features of the " + name + " vector from the given path; Auto-generated\nfunc (features *ProcessedFeatures) " + funcName + " (path *PathFeatures) {\n")
		funcNames = append(funcNames, funcName)

		if typ == "float64" {
			fmt.Print("\tfeatures." + newName + "Min, features." + newName + "Mean, features." + newName + "Max, features." + newName + "StdDev = rnd.StatBasic(path." + name + ", true) \n")
			fmt.Print("\tfeatures." + newName + "Sum = " + sumFunctions[typ] + "(path." + name + ")\n")
			fmt.Print("\tif len(path." + name + ") == 1 {\n\t\tfeatures." + newName + "Min, features." + newName + "Mean, features." + newName + "Max = path." + name + "[0], path." + name + "[0], path." + name + "[0]\n\t}\n")
			fmt.Print("\tfeatures." + newName + "Diff = features." + newName + "Max - features." + newName + "Min \n")
			fmt.Print("\tdefer func(){\n\t\trecover()\n\t}()\n")
			fmt.Print("\t_, _, _, _, _, features." + newName + "Skew ,_ = rnd.StatMoments(path." + name + ") \n")
		} else {
			fmt.Print("\tfeatures." + newName + "Max = " + maxFunctions[typ] + "(path." + name + ")\n")
			fmt.Print("\tfeatures." + newName + "Min = " + minFunctions[typ] + "(path." + name + ")\n")
			fmt.Print("\tfeatures." + newName + "Diff = features." + newName + "Max - features." + newName + "Min\n")
			fmt.Print("\tfeatures." + newName + "Sum = " + sumFunctions[typ] + "(path." + name + ")\n")
			fmt.Print("\tfeatures." + newName + "Mean = Mean(" + convertFunctions[typ] + "(path." + name + "),float64(features." + newName + "Sum))\n")
			fmt.Print("\tdefer func(){\n\t\trecover()\n\t}()\n")
			fmt.Print("\t_,_, _, _, features." + newName + "StdDev, features." + newName + "Skew ,_ = rnd.StatMoments(" + convertFunctions[typ] + "(path." + name + ")) \n")

		}
		fmt.Print("}\n")
	}

	fmt.Print("// GenerateCheckboxCalculation calculates the features from the checkbox movement path; Auto-generated\nfunc (features *ProcessedFeatures) GenerateCheckboxCalculation (path *PathFeatures) {\n")
	for _, v := range funcNames {
		fmt.Print("\tfeatures." + v + "(path)\n")
	}
	fmt.Print("}\n\n")
}
