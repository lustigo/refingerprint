package data

import "strings"

// gcVendors declares the possible graphiccard vendors
var gcVendors = []string{
	"Google Inc.",
	"Intel Inc.",
	"NOTNOM", // not nominal
	"NVIDIA Corporation",
	"X.Org",
}

// gcModels declares the possible graphiccard models
var gcModels = []string{
	"AMD Radeon(TM) Vega 8 Graphics",
	"ATI Radeon HD 3600",
	"GeForce GTX 550",
	"GeForce GTX 980",
	"Intel(R) HD Graphics",
	"Intel(R) UHD Graphics",
	"Intel Iris Pro",
	"NVIDIA GeForce GTX 1060",
	"NVIDIA GeForce GTX 1070",
	"NVIDIA GeForce GTX 970",
	"Radeon RX 570",
	"NOTNOM",       // not nominal
	"OTHER AMD",    // not nominal but AMD
	"OTHER INTEL",  // not nominal but Intel
	"OTHER NVIDIA", // not nominal but Nvidia
}

// ExtractGraphicCardInformation extracts the graphic card vendor and model from the given WebGL Info
func (features *ProcessedFeatures) ExtractGraphicCardInformation(data WebGLInfo) {
	found := false
	for _, v := range gcModels {
		if strings.Contains(data.GcModel, v) {
			features.WebGLGCModel = v
			found = true
			break
		}
	}

	if !found {
		if strings.Contains(data.GcModel, "NVIDIA") || strings.Contains(data.GcModel, "GeForce") || strings.Contains(data.GcModel, "GTX") {
			features.WebGLGCModel = gcModels[len(gcModels)-1]
		} else if strings.Contains(data.GcModel, "AMD") {
			features.WebGLGCModel = gcModels[len(gcModels)-3]
		} else if strings.Contains(data.GcModel, "Intel") {
			features.WebGLGCModel = gcModels[len(gcModels)-2]
		} else {
			features.WebGLGCModel = gcModels[len(gcModels)-4]
		}
	}

	if BinSearch(gcVendors, data.GcVendor) {
		features.WebGLGCVendor = data.GcVendor
	} else {
		features.WebGLGCVendor = gcVendors[2]
	}
}
