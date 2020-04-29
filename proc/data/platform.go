package data


// platforms holds possible values for navigator.platform
// Source: https://stackoverflow.com/questions/19877924/what-is-the-list-of-possible-values-for-navigator-platform-as-of-today
var platforms = []string{
	"Android",
	"Linux",
	"null",
	"",
	"iPhone",
	"iPod",
	"iPad",
	"iPhone Simulator",
	"iPod Simulator",
	"iPad Simulator",
	"Macintosh",
	"MacIntel",
	"MacPPC",
	"Mac68K",
	"Pike v7.6 release 92",
	"Pike v7.8 release 517",
	"BlackBerry",
	"FreeBSD",
	"FreeBSD i386",
	"FreeBSD amd64",
	"Linux aarch64",
	"Linux armv5tejl",
	"Linux armv6l",
	"Linux armv7l",
	"Linux i686",
	"Linux i686 on x86_64",
	"Linux i686 X11",
	"Linux MSM8960_v3.2.1.1_N_R069_Rev:18",
	"Linux ppc64",
	"Linux x86_64",
	"Linux x86_64 X11",
	"OS/2",
	"Pocket PC",
	"Win32",
	"Win16",
	"WinCE",
	"Windows",
	"New Nintendo 3DS",
	"Nintendo DSi",
	"Nintendo 3DS",
	"Nintendo Wii",
	"Nintendo WiiU",
	"OpenBSD amd64",
	"Nokia_Series_40",
	"S60",
	"Symbian",
	"Symbian OS",
	"PalmOS",
	"webOS",
	"SunOS",
	"SunOS i86pc",
	"SunOS sun4u",
	"PLAYSTATION 3",
	"PlayStation 4",
	"PSP",
	"HP-UX",
	"masking-agent",
	"WebTV OS",
	"X11",
}

// Contains checks if the given needle is in the haystack
func Contains(haystack []string, needle string) bool {
	for _, x := range haystack {
		if x == needle {
			return true
		}
	}
	return false
}

// ExtractPlatform checks if the given platform is listed, if not sets it to ""
func (features *ProcessedFeatures) ExtractPlatform(binfo BrowserInfo){
	if Contains(platforms,binfo.Platform){
		features.Platform = binfo.Platform
	} else {
		features.Platform = ""
	}
}
