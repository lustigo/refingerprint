package data

import "github.com/lustigo/go-arff"

// ProcessedFeatures contains all the extracted and calculated features from the raw data set
type ProcessedFeatures struct {
	UserID string `arff:"userID" csv:"userID"`
}

// GetARFFHeader returns the Header for an ARFF file which contains ProcessedFeatures instances
func GetARFFHeader() arff.Header {
	header := arff.Header{}
	header.AddAttr("userID", arff.String, nil)
	header.Relation = "refingerprint"
	return header
}

// ExtractFeatures extracts the Features from the raw data
func ExtractFeatures(data *Data) *ProcessedFeatures {
	features := &ProcessedFeatures{}
	features.UserID = data.UserID
	return features
}
