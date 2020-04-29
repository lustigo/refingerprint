package data

import (
	"github.com/avct/uasurfer"
)

// All the possible browsernames, osnames or device types
// Source: https://github.com/avct/uasurfer
var browsername = []string{
	"BrowserChrome",
	"BrowserSafari",
	"BrowserIE",
	"BrowserFirefox",
	"BrowserAndroid",
	"BrowserOpera",
	"BrowserUCBrowser",
	"BrowserSilk",
	"BrowserQQ",
	"BrowserSpotify",
	"BrowserBlackberry",
	"BrowserYandex",
	"BrowserNintendo",
	"BrowserSamsung",
	"BrowserCocCoc",
	"BrowserUnknown",
}
var osname = []string{
	"OSWindows",
	"OSMacOSX",
	"OSiOS",
	"OSAndroid",
	"OSChromeOS",
	"OSWebOS",
	"OSLinux",
	"OSPlaystation",
	"OSXbox",
	"OSNintendo",
	"OSUnknown",
}
var devicetype = []string{
	"DeviceComputer",
	"DevicePhone",
	"DeviceTablet",
	"DeviceTV",
	"DeviceConsole",
	"DeviceWearable",
	"DeviceUnknown",
}

// ExtractUserAgentInformation extracts the User Agent Informations and stores it
func (features *ProcessedFeatures) ExtractUserAgentInformation(binfo BrowserInfo) {
	ua := uasurfer.Parse(binfo.UserAgent)

	features.DeviceType = ua.DeviceType.String()
	features.OSName = ua.OS.Name.String()
	features.BrowserName = ua.Browser.Name.String()

	v := ua.OS.Version
	version := uint64(v.Major*100000000 + v.Minor*10000 + v.Patch)
	if features.OSName == "OSWindows" {
		features.WindowsVersion = version
	} else if features.OSName == "OSMacOSX" || features.OSName == "OSiOS" {
		features.AppleVersion = version
	} else {
		features.LinuxVersion = version
	}

	v = ua.Browser.Version
	version = uint64(v.Major*100000000 + v.Minor*10000 + v.Patch)
	if features.BrowserName == "BrowserChrome" {
		features.ChromeVersion = version
	} else if features.BrowserName == "BrowserSafari" {
		features.SafariVersion = version
	} else if features.BrowserName == "BrowserIE" {
		features.IEVersion = version
	} else if features.BrowserName == "BrowserFirefox" {
		features.FirefoxVersion = version
	} else if features.BrowserName == "BrowserOpera" {
		features.OperaVersion = version
	}
}
