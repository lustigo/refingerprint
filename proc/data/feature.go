package data

import "github.com/sbinet/go-arff"

// ProcessedFeatures contains all the extracted and calculated features from the raw data set
type ProcessedFeatures struct {
	// General Data
	UserID         string `arff:"userID" csv:"userID"`
	ProcessingTime uint64 `arff:"processingTime" csv:"processing time"`
	Canvas         string `arff:"canvasFingerprint" csv:"Canvas Fingerprint"`
	Audio          string `arff:"audioFingerprint" csv:"Audio Fingerprint"`
	// Browser Features
	CookiesEnabled uint8 `arff:"cookiesEnabled" csv:"Cookies Enabled"`
	CPUAmount      uint8 `arff:"cpuAmount" csv:"Amount of CPUs"`
	DNT            uint8 `arff:"dnt" csv:"dnt"` // 0 = false, 1 = true, 2 = not set
	TimeZone       int16 `arff:"timezone" csv:"timezone"`
	// Languages
	Lang1         string `arff:"lang1" csv:"preferred language"`
	Lang2         string `arff:"lang2" csv:"2nd preferred language"`
	Lang3         string `arff:"lang3" csv:"3rd preferred language"`
	Lang4         string `arff:"lang4" csv:"4th preferred language"`
	Lang5         string `arff:"lang5" csv:"5th preferred language"`
	Lang6         string `arff:"lang6" csv:"6th preferred language"`
	Lang7         string `arff:"lang7" csv:"7th preferred language"`
	Lang8         string `arff:"lang8" csv:"8th preferred language"`
	MoreLanguages uint8  `arff:"moreLang" csv:"More than 8 languages"` // 0 = false 1 = true
	// Mime-Types
	MIMEAppOctet     uint8 `arff:"mime-application_octet-stream" csv:"mime-application_octet-stream"`
	MIMEPlain        uint8 `arff:"mime-text_plain" csv:"mime-text_plain"`
	MIMECSS          uint8 `arff:"mime-text_css" csv:"mime-text_css"`
	MIMEHTML         uint8 `arff:"mime-text_html" csv:"mime-text_html"`
	MIMEJS           uint8 `arff:"mime-text_javascript" csv:"mime-text_javascript"`
	MIMEAPNG         uint8 `arff:"mime-image_apng" csv:"mime-image_apng"`
	MIMEBMP          uint8 `arff:"mime-image_bmp" csv:"mime-image_bmp"`
	MIMEGIF          uint8 `arff:"mime-image_gif" csv:"mime-image_gif"`
	MIMEICO          uint8 `arff:"mime-image_x-icon" csv:"mime-image_x-icon"`
	MIMEJPG          uint8 `arff:"mime-image_jpeg" csv:"mime-image_jpeg"`
	MIMEPNG          uint8 `arff:"mime-image_png" csv:"mime-image_png"`
	MIMESVG          uint8 `arff:"mime-image_svg-xml" csv:"mime-image_svg-xml"`
	MIMETIFF         uint8 `arff:"mime-image_tiff" csv:"mime-image_tiff"`
	MIMEWEBP         uint8 `arff:"mime-image_webp" csv:"mime-image_webp"`
	MIMEWAV          uint8 `arff:"mime-audio_wave" csv:"mime-audio_wave"`
	MIMEAudioWEBM    uint8 `arff:"mime-audio_webm" csv:"mime-audio_webm"`
	MIMEVideoWEBM    uint8 `arff:"mime-video_webm" csv:"mime-video_webm"`
	MIMEAudioOGG     uint8 `arff:"mime-audio_ogg" csv:"mime-audio_ogg"`
	MIMEVideoOGG     uint8 `arff:"mime-video_ogg" csv:"mime-video_ogg"`
	MIMEAppOGG       uint8 `arff:"mime-application_ogg" csv:"mime-application_ogg"`
	MIMEAppPDF       uint8 `arff:"mime-application_pdf" csv:"mime-application_pdf"`
	MIMEAppChromePDF uint8 `arff:"mime-application_chrome-pdf" csv:"mime-application_chrome-pdf"`
	MIMEForm         uint8 `arff:"mime-multipart_form-data" csv:"mime-multipart_form-data"`
	MIMEByte         uint8 `arff:"mime-multipart_byteranges" csv:"mime-multipart_byteranges"`
	// TODO: Plugins
	// Platforms
	Platform string `arff:"platform" csv:"Platform"`
	// User Agent
	BrowserName    string `arff:"browsername" csv:"browsername"`
	ChromeVersion  uint64 `arff:"chromeVersion" csv:"Chrome Version"`
	SafariVersion  uint64 `arff:"safariVersion" csv:"Safari Version"`
	IEVersion      uint64 `arff:"ieVersion" csv:"Internet Explorer Version"`
	FirefoxVersion uint64 `arff:"ffVersion" csv:"Firefox Version"`
	OperaVersion   uint64 `arff:"operaVersion" csv:"Opera Version"`
	OSName         string `arff:"osName" csv:"Operating System"`
	WindowsVersion uint64 `arff:"winVersion" csv:"Windows Version"`
	AppleVersion   uint64 `arff:"appleVersion" csv:"Apple Version"`
	LinuxVersion   uint64 `arff:"linVersion" csv:"Linux/Android Version"`
	DeviceType     string `arff:"deviceType" csv:"Device Type"`
}

// GetARFFHeader returns the Header for an ARFF file which contains ProcessedFeatures instances
func GetARFFHeader() arff.Header {
	header := arff.Header{}
	// General Data
	header.AddAttr("userID", arff.String, nil)
	header.AddAttr("processingTime", arff.Numeric, nil)
	header.AddAttr("canvasFingerprint", arff.String, nil)
	header.AddAttr("audioFingerprint", arff.String, nil)
	// Browser Info
	header.AddAttr("cookiesEnabled", arff.Numeric, nil)
	header.AddAttr("cpuAmount", arff.Numeric, nil)
	header.AddAttr("dnt", arff.Numeric, nil)
	header.AddAttr("timezone", arff.Numeric, nil)
	// Languages
	header.AddAttr("lang1", arff.Nominal, languages)
	header.AddAttr("lang2", arff.Nominal, languages)
	header.AddAttr("lang3", arff.Nominal, languages)
	header.AddAttr("lang4", arff.Nominal, languages)
	header.AddAttr("lang5", arff.Nominal, languages)
	header.AddAttr("lang6", arff.Nominal, languages)
	header.AddAttr("lang7", arff.Nominal, languages)
	header.AddAttr("lang8", arff.Nominal, languages)
	header.AddAttr("moreLang", arff.Numeric, nil)
	// MIME-Types
	header.AddAttr("mime-application_octet-stream", arff.Numeric, nil)
	header.AddAttr("mime-text_plain", arff.Numeric, nil)
	header.AddAttr("mime-text_css", arff.Numeric, nil)
	header.AddAttr("mime-text_html", arff.Numeric, nil)
	header.AddAttr("mime-text_javascript", arff.Numeric, nil)
	header.AddAttr("mime-image_apng", arff.Numeric, nil)
	header.AddAttr("mime-image_bmp", arff.Numeric, nil)
	header.AddAttr("mime-image_gif", arff.Numeric, nil)
	header.AddAttr("mime-image_x-icon", arff.Numeric, nil)
	header.AddAttr("mime-image_jpeg", arff.Numeric, nil)
	header.AddAttr("mime-image_png", arff.Numeric, nil)
	header.AddAttr("mime-image_svg-xml", arff.Numeric, nil)
	header.AddAttr("mime-image_tiff", arff.Numeric, nil)
	header.AddAttr("mime-image_webp", arff.Numeric, nil)
	header.AddAttr("mime-audio_wave", arff.Numeric, nil)
	header.AddAttr("mime-audio_webm", arff.Numeric, nil)
	header.AddAttr("mime-video_webm", arff.Numeric, nil)
	header.AddAttr("mime-audio_ogg", arff.Numeric, nil)
	header.AddAttr("mime-video_ogg", arff.Numeric, nil)
	header.AddAttr("mime-application_ogg", arff.Numeric, nil)
	header.AddAttr("mime-application_pdf", arff.Numeric, nil)
	header.AddAttr("mime-application_chrome-pdf", arff.Numeric, nil)
	header.AddAttr("mime-multipart_form-data", arff.Numeric, nil)
	header.AddAttr("mime-multipart_byteranges", arff.Numeric, nil)
	// TODO: Plugins
	// Platforms
	header.AddAttr("platform", arff.Nominal, platforms)
	// User Agent
	header.AddAttr("browsername", arff.Nominal, browsername)
	header.AddAttr("deviceType", arff.Nominal, devicetype)
	header.AddAttr("osName", arff.Nominal, osname)
	header.AddAttr("ieVersion", arff.Numeric, nil)
	header.AddAttr("ffVersion", arff.Numeric, nil)
	header.AddAttr("operaVersion", arff.Numeric, nil)
	header.AddAttr("chromeVersion", arff.Numeric, nil)
	header.AddAttr("safariVersion", arff.Numeric, nil)
	header.AddAttr("winVersion", arff.Numeric, nil)
	header.AddAttr("appleVersion", arff.Numeric, nil)
	header.AddAttr("linVersion", arff.Numeric, nil)

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
	features.ExtractBrowserFeatures(data)
	return features
}

// ExtractBrowserFeatures extracts the BrowserInfo Features from the raw data and stores it
func (features *ProcessedFeatures) ExtractBrowserFeatures(data *Data) {
	binfo := data.Browser
	features.CPUAmount = binfo.CPUAmount

	if binfo.CookieEnabled {
		features.CookiesEnabled = 1
	} else {
		features.CookiesEnabled = 0
	}

	if binfo.DNT == "" {
		features.DNT = 2
	} else if binfo.DNT == "true" {
		features.DNT = 1
	} else {
		features.DNT = 0
	}

	features.TimeZone = binfo.TimeZone

	features.ExtractLanguages(binfo)
	features.ExtractMIMETypes(binfo)
	features.ExtractPlatform(binfo)
	features.ExtractUserAgentInformation(binfo)
}
