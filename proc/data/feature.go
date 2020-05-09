package data

import (
	"sort"

	"github.com/sbinet/go-arff"
)

// ProcessedFeatures contains all the extracted and calculated features from the raw data set
type ProcessedFeatures struct {
	// General Data
	SessionID      string `arff:"sessionID" csv:"sessionID"`
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
	// WebGL Extensions
	WebGLExtMore                 uint16 `arff:"webglExtMore" csv:"Number of additional Extensions"`
	WebGLExtOESTextFloat         uint8  `arff:"webglExtOESTextFloat" csv:"OES_texture_float"`
	WebGLExtOESTextHalfFloat     uint8  `arff:"webglExtOESTextHalfFloat" csv:"OES_texture_half_float"`
	WebGLExtLose                 uint8  `arff:"webglExtLose" csv:"WEBGL_lose_context"`
	WebGLExtOESStd               uint8  `arff:"webglExtOESStd" csv:"OES_standard_derivatives"`
	WebGLExtOESVert              uint8  `arff:"webglExtOESVert" csv:"OES_vertex_array_object"`
	WebGLExtDebugRender          uint8  `arff:"webglExtDebugRender" csv:"WEBGL_debug_renderer_info"`
	WebGLExtDebugShader          uint8  `arff:"webglExtDebugShader" csv:"WEBGL_debug_shaders"`
	WebGLExtCompressTextS3TC     uint8  `arff:"webglExtCompressTextS3TC" csv:"WEBGL_compressed_texture_s3tc"`
	WebGLExtDepth                uint8  `arff:"webglExtDepth" csv:"WEBGL_depth_texture"`
	WebGLExtCompressTextPVRTC    uint8  `arff:"webglExtCompressTextPVRTC" csv:"WEBGL_compressed_texture_pvrtc"`
	WebGLExtCompressTextATC      uint8  `arff:"webglExtCompressTextATC" csv:"WEBGL_compressed_texture_atc"`
	WebGLExtOESEle               uint8  `arff:"webglExtOESEle" csv:"OES_element_index_uint"`
	WebGLExtTextFilter           uint8  `arff:"webglExtTextFilter" csv:"EXT_texture_filter_anisotropic"`
	WebGLExtFrag                 uint8  `arff:"webglExtFrag" csv:"EXT_frag_depth"`
	WebGLExtDraw                 uint8  `arff:"webglExtDraw" csv:"WEBGL_draw_buffers"`
	WebGLExtAngle                uint8  `arff:"webglExtAngle" csv:"ANGLE_instanced_arrays"`
	WebGLExtOESTextFloatLin      uint8  `arff:"webglExtOESTextFloatLin" csv:"OES_texture_float_linear"`
	WebGLExtOESTextHalfFloatLin  uint8  `arff:"webglExtOESTextHalfFloatLin" csv:"OES_texture_half_float_linear"`
	WebGLExtBlend                uint8  `arff:"webglExtBlend" csv:"EXT_blend_minmax"`
	WebGLExtShader               uint8  `arff:"webglExtShader" csv:"EXT_shader_texture_lod"`
	WebGLExtColorBuffFloat       uint8  `arff:"webglExtColorBuffFloat" csv:"EXT_color_buffer_float"`
	WebGLExtGLColorBuffFloat     uint8  `arff:"webglExtGLColorBuffFloat" csv:"WEBGL_color_buffer_float"`
	WebGLExtColorBuffHalfFloat   uint8  `arff:"webglExtColorBuffHalfFloat" csv:"EXT_color_buffer_half_float"`
	WebGLExtSrgb                 uint8  `arff:"webglExtSrgb" csv:"EXT_sRGB"`
	WebGLExtCompressTextETC1     uint8  `arff:"webglExtCompressTextETC1" csv:"WEBGL_compressed_texture_etc1"`
	WebGLExtDisjoint             uint8  `arff:"webglExtDisjoint" csv:"EXT_disjoint_timer_query"`
	WebGLExtOESFbo               uint8  `arff:"webglExtOESFbo" csv:"OES_fbo_render_mipmap"`
	WebGLExtCompressTextASTC     uint8  `arff:"webglExtCompressTextASTC" csv:"WEBGL_compressed_texture_astc"`
	WebGLExtCompressTextETC      uint8  `arff:"webglExtCompressTextETC" csv:"WEBGL_compressed_texture_etc"`
	WebGLExtCompressTextS3TCSrgb uint8  `arff:"webglExtCompressTextS3TCSrgb" csv:"WEBGL_compressed_texture_s3tc_srgb"`
	WebGLExtDisjointGL2          uint8  `arff:"webglExtDisjointGL2" csv:"EXT_disjoint_timer_query_webgl2"`
	WebGLExtFloat                uint8  `arff:"webglExtFloat" csv:"EXT_float_blend"`
	WebGLExtOVR                  uint8  `arff:"webglExtOVR" csv:"OVR_multiview2"`
	WebGLExtKHR                  uint8  `arff:"webglExtKHR" csv:"KHR_parallel_shader_compile"`
	WebGLExtTextCompressBPTC     uint8  `arff:"webglExtTextCompressBPTC" csv:"EXT_texture_compression_bptc"`
	WebGLExtTextCompressRGTC     uint8  `arff:"webglExtTextCompressRGTC" csv:"EXT_texture_compression_rgtc"`
	// TODO: Vendor and Model
	// Screen Resolution
	ScreenHeight uint16 `arff:"screenY" csv:"Screen Height"`
	PixelDensity uint16 `arff:"screenDensity" csv:"Pixel Density"`
	ScreenWidth  uint16 `arff:"screenX" csv:"Screen Width"`
	InnerWidth   uint16 `arff:"screenInnerX" csv:"Inner width of the screen"`
	InnerHeight  uint16 `arff:"screenInnerY" csv:"Inner height of the screen"`
	ScreenDeltaX uint16 `arff:"screenDX" csv:"Screen Delta X"`
	ScreenDeltaY uint16 `arff:"screenDY" csv:"Screen Delta Y"`
	// Frame Position
	CaptchaX      float64 `arff:"captchaX" csv:"Captchaframepos X"`
	CaptchaY      float64 `arff:"captchaY" csv:"Captchaframepos Y"`
	CaptchaWidth  float64 `arff:"captchaWidth" csv:"Captchaframepos Width"`
	CaptchaHeight float64 `arff:"captchaHeight" csv:"Captchaframepos Height"`
	TaskX         float64 `arff:"taskX" csv:"Taskframepos X"`
	TaskY         float64 `arff:"taskY" csv:"Taskframepos Y"`
	TaskWidth     float64 `arff:"taskWidth" csv:"Taskframepos Width"`
	TaskHeight    float64 `arff:"taskHeight" csv:"Taskframepos Height"`
}

// GetARFFHeader returns the Header for an ARFF file which contains ProcessedFeatures instances
func GetARFFHeader() arff.Header {
	header := arff.Header{}
	// General Data
	header.AddAttr("sessionID", arff.String, nil)
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
	header.AddAttr("platform", arff.Nominal, getPlatformList())
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
	// WebGL Extensions
	header.AddAttr("webglExtMore", arff.Numeric, nil)
	header.AddAttr("webglExtOESTextFloat", arff.Numeric, nil)
	header.AddAttr("webglExtOESTextHalfFloat", arff.Numeric, nil)
	header.AddAttr("webglExtLose", arff.Numeric, nil)
	header.AddAttr("webglExtOESStd", arff.Numeric, nil)
	header.AddAttr("webglExtOESVert", arff.Numeric, nil)
	header.AddAttr("webglExtDebugRender", arff.Numeric, nil)
	header.AddAttr("webglExtDebugShader", arff.Numeric, nil)
	header.AddAttr("webglExtCompressTextS3TC", arff.Numeric, nil)
	header.AddAttr("webglExtCompressTextPVRTC", arff.Numeric, nil)
	header.AddAttr("webglExtCompressTextATC", arff.Numeric, nil)
	header.AddAttr("webglExtDepth", arff.Numeric, nil)
	header.AddAttr("webglExtOESEle", arff.Numeric, nil)
	header.AddAttr("webglExtTextFilter", arff.Numeric, nil)
	header.AddAttr("webglExtFrag", arff.Numeric, nil)
	header.AddAttr("webglExtDraw", arff.Numeric, nil)
	header.AddAttr("webglExtAngle", arff.Numeric, nil)
	header.AddAttr("webglExtOESTextFloatLin", arff.Numeric, nil)
	header.AddAttr("webglExtOESTextHalfFloatLin", arff.Numeric, nil)
	header.AddAttr("webglExtBlend", arff.Numeric, nil)
	header.AddAttr("webglExtShader", arff.Numeric, nil)
	header.AddAttr("webglExtColorBuffFloat", arff.Numeric, nil)
	header.AddAttr("webglExtGLColorBuffFloat", arff.Numeric, nil)
	header.AddAttr("webglExtColorBuffHalfFloat", arff.Numeric, nil)
	header.AddAttr("webglExtCompressTextETC1", arff.Numeric, nil)
	header.AddAttr("webglExtCompressTextETC", arff.Numeric, nil)
	header.AddAttr("webglExtSrgb", arff.Numeric, nil)
	header.AddAttr("webglExtDisjoint", arff.Numeric, nil)
	header.AddAttr("webglExtOESFbo", arff.Numeric, nil)
	header.AddAttr("webglExtCompressTextASTC", arff.Numeric, nil)
	header.AddAttr("webglExtCompressTextS3TCSrgb", arff.Numeric, nil)
	header.AddAttr("webglExtDisjointGL2", arff.Numeric, nil)
	header.AddAttr("webglExtFloat", arff.Numeric, nil)
	header.AddAttr("webglExtOVR", arff.Numeric, nil)
	header.AddAttr("webglExtKHR", arff.Numeric, nil)
	header.AddAttr("webglExtTextCompressBPTC", arff.Numeric, nil)
	header.AddAttr("webglExtTextCompressRGTC", arff.Numeric, nil)
	// TODO: Vendor&Model
	// Screen Resolution
	header.AddAttr("screenX", arff.Numeric, nil)
	header.AddAttr("screenY", arff.Numeric, nil)
	header.AddAttr("screenDX", arff.Numeric, nil)
	header.AddAttr("screenDY", arff.Numeric, nil)
	header.AddAttr("screenDensity", arff.Numeric, nil)
	header.AddAttr("screenInnerX", arff.Numeric, nil)
	header.AddAttr("screenInnerY", arff.Numeric, nil)
	// Frame Positions
	header.AddAttr("captchaX", arff.Numeric, nil)
	header.AddAttr("captchaY", arff.Numeric, nil)
	header.AddAttr("captchaWidth", arff.Numeric, nil)
	header.AddAttr("captchaHeight", arff.Numeric, nil)
	header.AddAttr("taskX", arff.Numeric, nil)
	header.AddAttr("taskY", arff.Numeric, nil)
	header.AddAttr("taskWidth", arff.Numeric, nil)
	header.AddAttr("taskHeight", arff.Numeric, nil)

	header.Relation = "refingerprint"
	return header
}

// ExtractFeatures extracts the Features from the raw data
func ExtractFeatures(data *Data) *ProcessedFeatures {
	features := &ProcessedFeatures{}
	features.SessionID = data.SessionID
	features.ProcessingTime = data.Time.End - data.Time.Start
	features.Canvas = data.Canvas
	features.Audio = data.Audio
	features.ExtractBrowserFeatures(data)
	features.ExtractWebGLExtensions(data.WebGL)
	features.ExtractScreenInfo(data.Screen)
	features.ExtractFramePosition(data.FramePosition)
	features.ExtractMouseData(data)
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

// ExtractScreenInfo saves the information about the screen
func (features *ProcessedFeatures) ExtractScreenInfo(sinfo ScreenInfo) {
	features.ScreenDeltaX = sinfo.DeltaX
	features.ScreenDeltaY = sinfo.DeltaY
	features.PixelDensity = sinfo.PixelDensity
	features.ScreenHeight = sinfo.Height
	features.ScreenWidth = sinfo.Width
	features.InnerHeight = sinfo.InnerHeight
	features.InnerWidth = sinfo.InnerWidth
}

// ExtractFramePosition saves the information about the frame rectangles
func (features *ProcessedFeatures) ExtractFramePosition(pinfo FramePosition) {
	features.CaptchaHeight = pinfo.Captcha.Height
	features.CaptchaWidth = pinfo.Captcha.Width
	features.CaptchaX = pinfo.Captcha.X
	features.CaptchaY = pinfo.Captcha.Y
	features.TaskHeight = pinfo.Task.Height
	features.TaskWidth = pinfo.Task.Width
	features.TaskX = pinfo.Task.X
	features.TaskY = pinfo.Task.Y
}

// ExtractMouseData extracts relevant MouseMovement Features
func (features *ProcessedFeatures) ExtractMouseData(data *Data) {
	cleanedCheck := RemoveDuplicateMousePoints(NormalizeMouseData(data.MouseCheckbox, data.Screen, data.Time))
	cleanedRest := RemoveDuplicateMousePoints(NormalizeMouseData(data.MouseRest, data.Screen, data.Time))
	cleanedScroll := RemoveDuplicateScrollEvents(NormalizeScrollEvents(data.ScrollEvents, data.Screen, data.Time))
	cleanedClick := RemoveDuplicateClickEvents(NormalizeClickEvents(data.MouseClicks, data.Screen, data.Time))

	cleanedCheck, cleanedScroll = RemoveConcurrentEvents(cleanedClick, cleanedCheck, cleanedScroll)
	cleanedRest, cleanedScroll = RemoveConcurrentEvents(cleanedClick, cleanedRest, cleanedScroll)

	// Sort the Data according to the time
	// Using sort.SliceStable (InsertionSort) instead of sort.Slice (QuickSort) because data should already be sorted
	sort.SliceStable(cleanedCheck, func(i, j int) bool { return cleanedCheck[i].Time < cleanedCheck[j].Time })
	sort.SliceStable(cleanedRest, func(i, j int) bool { return cleanedRest[i].Time < cleanedRest[j].Time })
	sort.SliceStable(cleanedScroll, func(i, j int) bool { return cleanedScroll[i].Time < cleanedScroll[j].Time })
	sort.SliceStable(cleanedClick, func(i, j int) bool { return cleanedClick[i].Time < cleanedClick[j].Time })

	checkBoxPath, otherPaths := ExtractPaths(cleanedCheck, cleanedRest, cleanedClick, cleanedScroll)
	otherPaths = RemoveSmallPaths(otherPaths)
}
