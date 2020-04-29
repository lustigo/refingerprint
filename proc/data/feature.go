package data

import "github.com/sbinet/go-arff"

// ProcessedFeatures contains all the extracted and calculated features from the raw data set
type ProcessedFeatures struct {
	UserID         string `arff:"userID" csv:"userID"`
	ProcessingTime uint64 `arff:"processingTime" csv:"processing time"`
	Canvas         string `arff:"canvasFingerprint" csv:"Canvas Fingerprint"`
	Audio          string `arff:"audioFingerprint" csv:"Audio Fingerprint"`
}

// GetARFFHeader returns the Header for an ARFF file which contains ProcessedFeatures instances
func GetARFFHeader() arff.Header {
	header := arff.Header{}
	header.AddAttr("userID", arff.String, nil)
	header.AddAttr("processingTime", arff.Numeric, nil)
	header.AddAttr("canvasFingerprint", arff.String, nil)
	header.AddAttr("audioFingerprint", arff.String, nil)
	header.Relation = "refingerprint"
	return header
}

// ExtractFeatures extracts the Features from the raw data
func ExtractFeatures(data *Data) *ProcessedFeatures {
	features := &ProcessedFeatures{}
	features.UserID = data.UserID
	features.ProcessingTime = data.Time.End - data.Time.Start
	features.Canvas = data.Canvas
	features.Audio = data.Audio
	return features
}
