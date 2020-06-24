package data

import "strings"

// gcVendorsRaw declares the possible graphiccard vendors
var gcVendorsRaw = []string{
	"AMD",
	"Google Inc.",
	"Intel Inc.",
	"NOTNOM", // not nominal
	"NVIDIA Corporation",
	"X.Org",
}

// gcVendors declares the possible graphiccard vendors for the feature
var gcVendors = []string{
	"AMD",
	"GoogleInc.",
	"IntelInc.",
	"NOTNOM", // not nominal
	"NVIDIACorporation",
	"X.Org",
}

// gcModelsRaw declares the possible graphiccard models
var gcModelsRaw = []string{
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
}

// gcModels declares the possible graphiccard model feature values
var gcModels = []string{
	"AMD-Radeon(TM)Vega8Graphics",
	"ATI-RadeonHD3600",
	"GeForceGTX550",
	"GeForceGTX980",
	"Intel(R)HDGraphics",
	"Intel(R)UHDGraphics",
	"IntelIrisPro",
	"NVIDIA-GeForceGTX1060",
	"NVIDIA-GeForceGTX1070",
	"NVIDIA-GeForceGTX970",
	"RadeonRX570",
	"NOTNOM",      // not nominal
	"OTHERAMD",    // not nominal but AMD
	"OTHERINTEL",  // not nominal but Intel
	"OTHERNVIDIA", // not nominal but Nvidia
}

// ExtractGraphicCardInformation extracts the graphic card vendor and model from the given WebGL Info
func (features *ProcessedFeatures) ExtractGraphicCardInformation(data WebGLInfo) {
	found := false
	for i, v := range gcModelsRaw {
		if strings.Contains(data.GcModel, v) {
			features.WebGLGCModel = gcModels[i]
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

	found = false
	for i, v := range gcVendorsRaw {
		if v == data.GcVendor {
			features.WebGLGCVendor = gcVendors[i]
			found = true
			break
		}
	}
	if !found {
		features.WebGLGCVendor = gcVendors[3]
	}
}
