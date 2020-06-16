package data

import "reflect"

// plugins declares all the possible plugins and maps them to the corresponding structure attribute names
var plugins = map[string]string{
	"Chrome PDF Plugin":   "PluginChromePDF",
	"Chrome PDF Viewer":   "PluginChromePDFViewer",
	"Chromium PDF Plugin": "PluginChromiumPDF",
	"Chromium PDF Viewer": "PluginChromiumPDFViewer",
	"Native Client":       "PluginNativeClient",
}

// ExtractPlugins extracts the Plugins from the Browserinfo and turns them into features
func (features *ProcessedFeatures) ExtractPlugins(data BrowserInfo) {
	f := reflect.ValueOf(features).Elem()
	for _, plugin := range data.Plugins {
		v, o := plugins[plugin]
		if !o {
			features.PluginMore++
			continue
		}
		f.FieldByName(v).SetUint(1)
	}
}
