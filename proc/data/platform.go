package data

// platforms holds possible values for navigator.platform
// Source: https://stackoverflow.com/questions/19877924/what-is-the-list-of-possible-values-for-navigator-platform-as-of-today
var platforms = map[string]string{
	"Android":                              "Android",
	"Linux":                                "Linux",
	"null":                                 "null",
	"":                                     "",
	"iPhone":                               "iPhone",
	"iPod":                                 "iPod",
	"iPad":                                 "iPad",
	"iPhone Simulator":                     "iPhoneSim",
	"iPod Simulator":                       "iPodSim",
	"iPad Simulator":                       "iPadSim",
	"Macintosh":                            "Macintosh",
	"MacIntel":                             "MacIntel",
	"MacPPC":                               "MacPPC",
	"Mac68K":                               "Mac68K",
	"Pike v7.6 release 92":                 "Pike7692",
	"Pike v7.8 release 517":                "Pike78517",
	"BlackBerry":                           "BlackBerry",
	"FreeBSD":                              "FreeBSD",
	"FreeBSD i386":                         "FreeBSD386",
	"FreeBSD amd64":                        "FreeBSD64",
	"Linux aarch64":                        "Linuxaarch64",
	"Linux armv5tejl":                      "Linuxarmv5tejl",
	"Linux armv6l":                         "Linuxarmv6l",
	"Linux armv7l":                         "Linuxarmv7l",
	"Linux i686":                           "Linux686",
	"Linux i686 on x86_64":                 "Linux6868664",
	"Linux i686 X11":                       "Linux68611",
	"Linux MSM8960_v3.2.1.1_N_R069_Rev:18": "LinuxMSM",
	"Linux ppc64":                          "LinuxPPC",
	"Linux x86_64":                         "Linux8664",
	"Linux x86_64 X11":                     "Linux866411",
	"OS/2":                                 "OS2",
	"Pocket PC":                            "PocketPC",
	"Win32":                                "Win32",
	"Win16":                                "Win16",
	"WinCE":                                "WinCE",
	"Windows":                              "Windows",
	"New Nintendo 3DS":                     "NewNintendo3DS",
	"Nintendo DSi":                         "NintendoDSi",
	"Nintendo 3DS":                         "Nintendo3DS",
	"Nintendo Wii":                         "NintendoWii",
	"Nintendo WiiU":                        "NintendoWiiU",
	"OpenBSD amd64":                        "OpenBSD64",
	"Nokia_Series_40":                      "Nokia40",
	"S60":                                  "S60",
	"Symbian":                              "Symbian",
	"Symbian OS":                           "SymbianOS",
	"PalmOS":                               "PalmOS",
	"webOS":                                "webOS",
	"SunOS":                                "SunOS",
	"SunOS i86pc":                          "SunOS86",
	"SunOS sun4u":                          "SunOS4",
	"PLAYSTATION 3":                        "Playstation3",
	"PlayStation 4":                        "Playstation4",
	"PSP":                                  "PSP",
	"HP-UX":                                "HPUX",
	"masking-agent":                        "masking",
	"WebTV OS":                             "WebTVOS",
	"X11":                                  "X11",
}

// ExtractPlatform checks if the given platform is listed, if not sets it to ""
func (features *ProcessedFeatures) ExtractPlatform(binfo BrowserInfo) {
	v, o := platforms[binfo.Platform]
	if o {
		features.Platform = v
	} else {
		features.Platform = ""
	}
}

// getPlatformList returns the list of possible values for the platform attribute
func getPlatformList() []string {
	list := make([]string, len(platforms))
	i := 0
	for _, v := range platforms {
		list[i] = v
		i++
	}
	return list
}
