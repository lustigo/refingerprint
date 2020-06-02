package data

import (
	"sort"

	"github.com/cpmech/gosl/rnd"
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
	//
	// Movement
	//
	NumberOfPaths uint16 `arff:"numberOfPaths" csv:"Number of Paths including checkboxPath"`
	// Checkbox Path Features
	CheckboxPathDistanceStartEndPoint            float64 `arff:"checkBoxPathDistStartEnd" csv:"checkBoxPathDistStartEnd"`
	CheckboxPathNumberOfMovementPoints           uint16  `arff:"checkBoxPathNumMovementPoints" csv:"checkBoxPathNumMovementPoints"`
	CheckboxPathAngleStartEndPoint               float64 `arff:"checkBoxPathAngleStartEnd" csv:"checkBoxPathAngleStartEnd"`
	CheckboxPathDurationOfPath                   uint64  `arff:"checkBoxPathDurPath" csv:"checkBoxPathDurPath"`
	CheckboxPathTimeBetweenClickAndMovement      uint64  `arff:"checkBoxPathTimeClickMovement" csv:"checkBoxPathTimeClickMovement"`
	CheckboxPathTimeBetweeenMovementAndDownClick uint64  `arff:"checkBoxPathTimeMovementClick" csv:"checkBoxPathTimeMovementClick"`
	CheckboxPathNumberOfBreaks                   uint16  `arff:"checkBoxPathNoBreaks" csv:"checkBoxPathNoBreaks"`
	CheckboxPathBreakTimeTotalTimeRatio          float64 `arff:"checkBoxPathBreakRatio" csv:"checkBoxPathBreakRatio"`
	CheckboxPathNumberOfRightClicks              uint8   `arff:"checkBoxPathNoRightClick" csv:"checkBoxPathNoRightClick"`
	CheckboxPathNumberOfMiddleClicks             uint8   `arff:"checkBoxPathNoMiddleClick" csv:"checkBoxPathNoMiddleClick"`
	CheckboxPathNumberOfScrolls                  uint8   `arff:"checkBoxPathNoScrolls" csv:"checkBoxPathNoScrolls"`
	CheckboxPathStraightness                     float64 `arff:"checkBoxPathStraightness" csv:"checkBoxPathStraightness"`
	// Auto-Generated
	//
	CheckBoxPathPairwiseDistanceMin                   float64 `arff:"checkboxpathPairwiseDistanceMin" csv:"checkboxpathPairwiseDistanceMin"`
	CheckBoxPathPairwiseDistanceMax                   float64 `arff:"checkboxpathPairwiseDistanceMax" csv:"checkboxpathPairwiseDistanceMax"`
	CheckBoxPathPairwiseDistanceDiff                  float64 `arff:"checkboxpathPairwiseDistanceDiff" csv:"checkboxpathPairwiseDistanceDiff"`
	CheckBoxPathPairwiseDistanceSum                   float64 `arff:"checkboxpathPairwiseDistanceSum" csv:"checkboxpathPairwiseDistanceSum"`
	CheckBoxPathPairwiseDistanceMean                  float64 `arff:"checkboxpathPairwiseDistanceMean" csv:"checkboxpathPairwiseDistanceMean"`
	CheckBoxPathPairwiseDistanceStdDev                float64 `arff:"checkboxpathPairwiseDistanceStdDev" csv:"checkboxpathPairwiseDistanceStdDev"`
	CheckBoxPathPairwiseDistanceSkew                  float64 `arff:"checkboxpathPairwiseDistanceSkew" csv:"checkboxpathPairwiseDistanceSkew"`
	CheckBoxPathPairwiseVelocityMin                   float64 `arff:"checkboxpathPairwiseVelocityMin" csv:"checkboxpathPairwiseVelocityMin"`
	CheckBoxPathPairwiseVelocityMax                   float64 `arff:"checkboxpathPairwiseVelocityMax" csv:"checkboxpathPairwiseVelocityMax"`
	CheckBoxPathPairwiseVelocityDiff                  float64 `arff:"checkboxpathPairwiseVelocityDiff" csv:"checkboxpathPairwiseVelocityDiff"`
	CheckBoxPathPairwiseVelocitySum                   float64 `arff:"checkboxpathPairwiseVelocitySum" csv:"checkboxpathPairwiseVelocitySum"`
	CheckBoxPathPairwiseVelocityMean                  float64 `arff:"checkboxpathPairwiseVelocityMean" csv:"checkboxpathPairwiseVelocityMean"`
	CheckBoxPathPairwiseVelocityStdDev                float64 `arff:"checkboxpathPairwiseVelocityStdDev" csv:"checkboxpathPairwiseVelocityStdDev"`
	CheckBoxPathPairwiseVelocitySkew                  float64 `arff:"checkboxpathPairwiseVelocitySkew" csv:"checkboxpathPairwiseVelocitySkew"`
	CheckBoxPathPairwiseAccelerationMin               float64 `arff:"checkboxpathPairwiseAccelerationMin" csv:"checkboxpathPairwiseAccelerationMin"`
	CheckBoxPathPairwiseAccelerationMax               float64 `arff:"checkboxpathPairwiseAccelerationMax" csv:"checkboxpathPairwiseAccelerationMax"`
	CheckBoxPathPairwiseAccelerationDiff              float64 `arff:"checkboxpathPairwiseAccelerationDiff" csv:"checkboxpathPairwiseAccelerationDiff"`
	CheckBoxPathPairwiseAccelerationSum               float64 `arff:"checkboxpathPairwiseAccelerationSum" csv:"checkboxpathPairwiseAccelerationSum"`
	CheckBoxPathPairwiseAccelerationMean              float64 `arff:"checkboxpathPairwiseAccelerationMean" csv:"checkboxpathPairwiseAccelerationMean"`
	CheckBoxPathPairwiseAccelerationStdDev            float64 `arff:"checkboxpathPairwiseAccelerationStdDev" csv:"checkboxpathPairwiseAccelerationStdDev"`
	CheckBoxPathPairwiseAccelerationSkew              float64 `arff:"checkboxpathPairwiseAccelerationSkew" csv:"checkboxpathPairwiseAccelerationSkew"`
	CheckBoxPathPairwiseAngleMin                      float64 `arff:"checkboxpathPairwiseAngleMin" csv:"checkboxpathPairwiseAngleMin"`
	CheckBoxPathPairwiseAngleMax                      float64 `arff:"checkboxpathPairwiseAngleMax" csv:"checkboxpathPairwiseAngleMax"`
	CheckBoxPathPairwiseAngleDiff                     float64 `arff:"checkboxpathPairwiseAngleDiff" csv:"checkboxpathPairwiseAngleDiff"`
	CheckBoxPathPairwiseAngleSum                      float64 `arff:"checkboxpathPairwiseAngleSum" csv:"checkboxpathPairwiseAngleSum"`
	CheckBoxPathPairwiseAngleMean                     float64 `arff:"checkboxpathPairwiseAngleMean" csv:"checkboxpathPairwiseAngleMean"`
	CheckBoxPathPairwiseAngleStdDev                   float64 `arff:"checkboxpathPairwiseAngleStdDev" csv:"checkboxpathPairwiseAngleStdDev"`
	CheckBoxPathPairwiseAngleSkew                     float64 `arff:"checkboxpathPairwiseAngleSkew" csv:"checkboxpathPairwiseAngleSkew"`
	CheckBoxPathAngleBetweenMovementAndStartEndMin    float64 `arff:"checkboxpathAngleBetweenMovementAndStartEndMin" csv:"checkboxpathAngleBetweenMovementAndStartEndMin"`
	CheckBoxPathAngleBetweenMovementAndStartEndMax    float64 `arff:"checkboxpathAngleBetweenMovementAndStartEndMax" csv:"checkboxpathAngleBetweenMovementAndStartEndMax"`
	CheckBoxPathAngleBetweenMovementAndStartEndDiff   float64 `arff:"checkboxpathAngleBetweenMovementAndStartEndDiff" csv:"checkboxpathAngleBetweenMovementAndStartEndDiff"`
	CheckBoxPathAngleBetweenMovementAndStartEndSum    float64 `arff:"checkboxpathAngleBetweenMovementAndStartEndSum" csv:"checkboxpathAngleBetweenMovementAndStartEndSum"`
	CheckBoxPathAngleBetweenMovementAndStartEndMean   float64 `arff:"checkboxpathAngleBetweenMovementAndStartEndMean" csv:"checkboxpathAngleBetweenMovementAndStartEndMean"`
	CheckBoxPathAngleBetweenMovementAndStartEndStdDev float64 `arff:"checkboxpathAngleBetweenMovementAndStartEndStdDev" csv:"checkboxpathAngleBetweenMovementAndStartEndStdDev"`
	CheckBoxPathAngleBetweenMovementAndStartEndSkew   float64 `arff:"checkboxpathAngleBetweenMovementAndStartEndSkew" csv:"checkboxpathAngleBetweenMovementAndStartEndSkew"`
	CheckBoxPathPairwiseAngularVelocityMin            float64 `arff:"checkboxpathPairwiseAngularVelocityMin" csv:"checkboxpathPairwiseAngularVelocityMin"`
	CheckBoxPathPairwiseAngularVelocityMax            float64 `arff:"checkboxpathPairwiseAngularVelocityMax" csv:"checkboxpathPairwiseAngularVelocityMax"`
	CheckBoxPathPairwiseAngularVelocityDiff           float64 `arff:"checkboxpathPairwiseAngularVelocityDiff" csv:"checkboxpathPairwiseAngularVelocityDiff"`
	CheckBoxPathPairwiseAngularVelocitySum            float64 `arff:"checkboxpathPairwiseAngularVelocitySum" csv:"checkboxpathPairwiseAngularVelocitySum"`
	CheckBoxPathPairwiseAngularVelocityMean           float64 `arff:"checkboxpathPairwiseAngularVelocityMean" csv:"checkboxpathPairwiseAngularVelocityMean"`
	CheckBoxPathPairwiseAngularVelocityStdDev         float64 `arff:"checkboxpathPairwiseAngularVelocityStdDev" csv:"checkboxpathPairwiseAngularVelocityStdDev"`
	CheckBoxPathPairwiseAngularVelocitySkew           float64 `arff:"checkboxpathPairwiseAngularVelocitySkew" csv:"checkboxpathPairwiseAngularVelocitySkew"`
	CheckBoxPathPairwiseDurationMin                   uint64  `arff:"checkboxpathPairwiseDurationMin" csv:"checkboxpathPairwiseDurationMin"`
	CheckBoxPathPairwiseDurationMax                   uint64  `arff:"checkboxpathPairwiseDurationMax" csv:"checkboxpathPairwiseDurationMax"`
	CheckBoxPathPairwiseDurationDiff                  uint64  `arff:"checkboxpathPairwiseDurationDiff" csv:"checkboxpathPairwiseDurationDiff"`
	CheckBoxPathPairwiseDurationSum                   uint64  `arff:"checkboxpathPairwiseDurationSum" csv:"checkboxpathPairwiseDurationSum"`
	CheckBoxPathPairwiseDurationMean                  float64 `arff:"checkboxpathPairwiseDurationMean" csv:"checkboxpathPairwiseDurationMean"`
	CheckBoxPathPairwiseDurationStdDev                float64 `arff:"checkboxpathPairwiseDurationStdDev" csv:"checkboxpathPairwiseDurationStdDev"`
	CheckBoxPathPairwiseDurationSkew                  float64 `arff:"checkboxpathPairwiseDurationSkew" csv:"checkboxpathPairwiseDurationSkew"`
	CheckBoxPathTimeBetweenClickAndReleaseMin         uint64  `arff:"checkboxpathTimeBetweenClickAndReleaseMin" csv:"checkboxpathTimeBetweenClickAndReleaseMin"`
	CheckBoxPathTimeBetweenClickAndReleaseMax         uint64  `arff:"checkboxpathTimeBetweenClickAndReleaseMax" csv:"checkboxpathTimeBetweenClickAndReleaseMax"`
	CheckBoxPathTimeBetweenClickAndReleaseDiff        uint64  `arff:"checkboxpathTimeBetweenClickAndReleaseDiff" csv:"checkboxpathTimeBetweenClickAndReleaseDiff"`
	CheckBoxPathTimeBetweenClickAndReleaseSum         uint64  `arff:"checkboxpathTimeBetweenClickAndReleaseSum" csv:"checkboxpathTimeBetweenClickAndReleaseSum"`
	CheckBoxPathTimeBetweenClickAndReleaseMean        float64 `arff:"checkboxpathTimeBetweenClickAndReleaseMean" csv:"checkboxpathTimeBetweenClickAndReleaseMean"`
	CheckBoxPathTimeBetweenClickAndReleaseStdDev      float64 `arff:"checkboxpathTimeBetweenClickAndReleaseStdDev" csv:"checkboxpathTimeBetweenClickAndReleaseStdDev"`
	CheckBoxPathTimeBetweenClickAndReleaseSkew        float64 `arff:"checkboxpathTimeBetweenClickAndReleaseSkew" csv:"checkboxpathTimeBetweenClickAndReleaseSkew"`
	CheckBoxPathBreakTimesMin                         uint64  `arff:"checkboxpathBreakTimesMin" csv:"checkboxpathBreakTimesMin"`
	CheckBoxPathBreakTimesMax                         uint64  `arff:"checkboxpathBreakTimesMax" csv:"checkboxpathBreakTimesMax"`
	CheckBoxPathBreakTimesDiff                        uint64  `arff:"checkboxpathBreakTimesDiff" csv:"checkboxpathBreakTimesDiff"`
	CheckBoxPathBreakTimesSum                         uint64  `arff:"checkboxpathBreakTimesSum" csv:"checkboxpathBreakTimesSum"`
	CheckBoxPathBreakTimesMean                        float64 `arff:"checkboxpathBreakTimesMean" csv:"checkboxpathBreakTimesMean"`
	CheckBoxPathBreakTimesStdDev                      float64 `arff:"checkboxpathBreakTimesStdDev" csv:"checkboxpathBreakTimesStdDev"`
	CheckBoxPathBreakTimesSkew                        float64 `arff:"checkboxpathBreakTimesSkew" csv:"checkboxpathBreakTimesSkew"`
	CheckBoxPathMovementDuringClickDistanceMin        float64 `arff:"checkboxpathMovementDuringClickDistanceMin" csv:"checkboxpathMovementDuringClickDistanceMin"`
	CheckBoxPathMovementDuringClickDistanceMax        float64 `arff:"checkboxpathMovementDuringClickDistanceMax" csv:"checkboxpathMovementDuringClickDistanceMax"`
	CheckBoxPathMovementDuringClickDistanceDiff       float64 `arff:"checkboxpathMovementDuringClickDistanceDiff" csv:"checkboxpathMovementDuringClickDistanceDiff"`
	CheckBoxPathMovementDuringClickDistanceSum        float64 `arff:"checkboxpathMovementDuringClickDistanceSum" csv:"checkboxpathMovementDuringClickDistanceSum"`
	CheckBoxPathMovementDuringClickDistanceMean       float64 `arff:"checkboxpathMovementDuringClickDistanceMean" csv:"checkboxpathMovementDuringClickDistanceMean"`
	CheckBoxPathMovementDuringClickDistanceStdDev     float64 `arff:"checkboxpathMovementDuringClickDistanceStdDev" csv:"checkboxpathMovementDuringClickDistanceStdDev"`
	CheckBoxPathMovementDuringClickDistanceSkew       float64 `arff:"checkboxpathMovementDuringClickDistanceSkew" csv:"checkboxpathMovementDuringClickDistanceSkew"`
	CheckBoxPathMovementDuringClickVelocityMin        float64 `arff:"checkboxpathMovementDuringClickVelocityMin" csv:"checkboxpathMovementDuringClickVelocityMin"`
	CheckBoxPathMovementDuringClickVelocityMax        float64 `arff:"checkboxpathMovementDuringClickVelocityMax" csv:"checkboxpathMovementDuringClickVelocityMax"`
	CheckBoxPathMovementDuringClickVelocityDiff       float64 `arff:"checkboxpathMovementDuringClickVelocityDiff" csv:"checkboxpathMovementDuringClickVelocityDiff"`
	CheckBoxPathMovementDuringClickVelocitySum        float64 `arff:"checkboxpathMovementDuringClickVelocitySum" csv:"checkboxpathMovementDuringClickVelocitySum"`
	CheckBoxPathMovementDuringClickVelocityMean       float64 `arff:"checkboxpathMovementDuringClickVelocityMean" csv:"checkboxpathMovementDuringClickVelocityMean"`
	CheckBoxPathMovementDuringClickVelocityStdDev     float64 `arff:"checkboxpathMovementDuringClickVelocityStdDev" csv:"checkboxpathMovementDuringClickVelocityStdDev"`
	CheckBoxPathMovementDuringClickVelocitySkew       float64 `arff:"checkboxpathMovementDuringClickVelocitySkew" csv:"checkboxpathMovementDuringClickVelocitySkew"`
	CheckBoxPathMovementDuringClickAccelerationMin    float64 `arff:"checkboxpathMovementDuringClickAccelerationMin" csv:"checkboxpathMovementDuringCilckAccelerationMin"`
	CheckBoxPathMovementDuringClickAccelerationMax    float64 `arff:"checkboxpathMovementDuringClickAccelerationMax" csv:"checkboxpathMovementDuringCilckAccelerationMax"`
	CheckBoxPathMovementDuringClickAccelerationDiff   float64 `arff:"checkboxpathMovementDuringClickAccelerationDiff" csv:"checkboxpathMovementDuringClickAccelerationDiff"`
	CheckBoxPathMovementDuringClickAccelerationSum    float64 `arff:"checkboxpathMovementDuringClickAccelerationSum" csv:"checkboxpathMovementDuringCilckAccelerationSum"`
	CheckBoxPathMovementDuringClickAccelerationMean   float64 `arff:"checkboxpathMovementDuringClickAccelerationMean" csv:"checkboxpathMovementDuringClickAccelerationMean"`
	CheckBoxPathMovementDuringClickAccelerationStdDev float64 `arff:"checkboxpathMovementDuringClickAccelerationStdDev" csv:"checkboxpathMovementDurinlgCickAccelerationStdDev"`
	CheckBoxPathMovementDuringClickAccelerationSkew   float64 `arff:"checkboxpathMovementDuringClickAccelerationSkew" csv:"checkboxpathMovementDuringClickAccelerationSkew"`
	CheckBoxPathMovementDuringClickAngleMin           float64 `arff:"checkboxpathMovementDuringClickAngleMin" csv:"checkboxpathMovementDuringClickAngleMin"`
	CheckBoxPathMovementDuringClickAngleMax           float64 `arff:"checkboxpathMovementDuringClickAngleMax" csv:"checkboxpathMovementDuringClickAngleMax"`
	CheckBoxPathMovementDuringClickAngleDiff          float64 `arff:"checkboxpathMovementDuringClickAngleDiff" csv:"checkboxpathMovementDuringClickAngleDiff"`
	CheckBoxPathMovementDuringClickAngleSum           float64 `arff:"checkboxpathMovementDuringClickAngleSum" csv:"checkboxpathMovementDuringClickAngleSum"`
	CheckBoxPathMovementDuringClickAngleMean          float64 `arff:"checkboxpathMovementDuringClickAngleMean" csv:"checkboxpathMovementDuringClickAngleMean"`
	CheckBoxPathMovementDuringClickAngleStdDev        float64 `arff:"checkboxpathMovementDuringClickAngleStdDev" csv:"checkboxpathMovementDuringClickAngleStdDev"`
	CheckBoxPathMovementDuringClickAngleSkew          float64 `arff:"checkboxpathMovementDuringClickAngleSkew" csv:"checkboxpathMovementDuringClickAngleSkew"`
	CheckBoxPathScrollDXMin                           float64 `arff:"checkboxpathScrollDXMin" csv:"checkboxpathScrollDXMin"`
	CheckBoxPathScrollDXMax                           float64 `arff:"checkboxpathScrollDXMax" csv:"checkboxpathScrollDXMax"`
	CheckBoxPathScrollDXDiff                          float64 `arff:"checkboxpathScrollDXDiff" csv:"checkboxpathScrollDXDiff"`
	CheckBoxPathScrollDXSum                           float64 `arff:"checkboxpathScrollDXSum" csv:"checkboxpathScrollDXSum"`
	CheckBoxPathScrollDXMean                          float64 `arff:"checkboxpathScrollDXMean" csv:"checkboxpathScrollDXMean"`
	CheckBoxPathScrollDXStdDev                        float64 `arff:"checkboxpathScrollDXStdDev" csv:"checkboxpathScrollDXStdDev"`
	CheckBoxPathScrollDXSkew                          float64 `arff:"checkboxpathScrollDXSkew" csv:"checkboxpathScrollDXSkew"`
	CheckBoxPathScrollDYMin                           float64 `arff:"checkboxpathScrollDYMin" csv:"checkboxpathScrollDYMin"`
	CheckBoxPathScrollDYMax                           float64 `arff:"checkboxpathScrollDYMax" csv:"checkboxpathScrollDYMax"`
	CheckBoxPathScrollDYDiff                          float64 `arff:"checkboxpathScrollDYDiff" csv:"checkboxpathScrollDYDiff"`
	CheckBoxPathScrollDYSum                           float64 `arff:"checkboxpathScrollDYSum" csv:"checkboxpathScrollDYSum"`
	CheckBoxPathScrollDYMean                          float64 `arff:"checkboxpathScrollDYMean" csv:"checkboxpathScrollDYMean"`
	CheckBoxPathScrollDYStdDev                        float64 `arff:"checkboxpathScrollDYStdDev" csv:"checkboxpathScrollDYStdDev"`
	CheckBoxPathScrollDYSkew                          float64 `arff:"checkboxpathScrollDYSkew" csv:"checkboxpathScrollDYSkew"`
	CheckBoxPathScrollDZMin                           float64 `arff:"checkboxpathScrollDZMin" csv:"checkboxpathScrollDZMin"`
	CheckBoxPathScrollDZMax                           float64 `arff:"checkboxpathScrollDZMax" csv:"checkboxpathScrollDZMax"`
	CheckBoxPathScrollDZDiff                          float64 `arff:"checkboxpathScrollDZDiff" csv:"checkboxpathScrollDZDiff"`
	CheckBoxPathScrollDZSum                           float64 `arff:"checkboxpathScrollDZSum" csv:"checkboxpathScrollDZSum"`
	CheckBoxPathScrollDZMean                          float64 `arff:"checkboxpathScrollDZMean" csv:"checkboxpathScrollDZMean"`
	CheckBoxPathScrollDZStdDev                        float64 `arff:"checkboxpathScrollDZStdDev" csv:"checkboxpathScrollDZStdDev"`
	CheckBoxPathScrollDZSkew                          float64 `arff:"checkboxpathScrollDZSkew" csv:"checkboxpathScrollDZSkew"`
	CheckBoxPathScrollDMMin                           uint8   `arff:"checkboxpathScrollDMMin" csv:"checkboxpathScrollDMMin"`
	CheckBoxPathScrollDMMax                           uint8   `arff:"checkboxpathScrollDMMax" csv:"checkboxpathScrollDMMax"`
	CheckBoxPathScrollDMDiff                          uint8   `arff:"checkboxpathScrollDMDiff" csv:"checkboxpathScrollDMDiff"`
	CheckBoxPathScrollDMSum                           uint8   `arff:"checkboxpathScrollDMSum" csv:"checkboxpathScrollDMSum"`
	CheckBoxPathScrollDMMean                          float64 `arff:"checkboxpathScrollDMMean" csv:"checkboxpathScrollDMMean"`
	CheckBoxPathScrollDMStdDev                        float64 `arff:"checkboxpathScrollDMStdDev" csv:"checkboxpathScrollDMStdDev"`
	CheckBoxPathScrollDMSkew                          float64 `arff:"checkboxpathScrollDMSkew" csv:"checkboxpathScrollDMSkew"`
	CheckBoxPathXPointsMin                            float64 `arff:"checkboxpathXPointsMin" csv:"checkboxpathXPointsMin"`
	CheckBoxPathXPointsMax                            float64 `arff:"checkboxpathXPointsMax" csv:"checkboxpathXPointsMax"`
	CheckBoxPathXPointsDiff                           float64 `arff:"checkboxpathXPointsDiff" csv:"checkboxpathXPointsDiff"`
	CheckBoxPathXPointsSum                            float64 `arff:"checkboxpathXPointsSum" csv:"checkboxpathXPointsSum"`
	CheckBoxPathXPointsMean                           float64 `arff:"checkboxpathXPointsMean" csv:"checkboxpathXPointsMean"`
	CheckBoxPathXPointsStdDev                         float64 `arff:"checkboxpathXPointsStdDev" csv:"checkboxpathXPointsStdDev"`
	CheckBoxPathXPointsSkew                           float64 `arff:"checkboxpathXPointsSkew" csv:"checkboxpathXPointsSkew"`
	CheckBoxPathYPointsMin                            float64 `arff:"checkboxpathYPointsMin" csv:"checkboxpathYPointsMin"`
	CheckBoxPathYPointsMax                            float64 `arff:"checkboxpathYPointsMax" csv:"checkboxpathYPointsMax"`
	CheckBoxPathYPointsDiff                           float64 `arff:"checkboxpathYPointsDiff" csv:"checkboxpathYPointsDiff"`
	CheckBoxPathYPointsSum                            float64 `arff:"checkboxpathYPointsSum" csv:"checkboxpathYPointsSum"`
	CheckBoxPathYPointsMean                           float64 `arff:"checkboxpathYPointsMean" csv:"checkboxpathYPointsMean"`
	CheckBoxPathYPointsStdDev                         float64 `arff:"checkboxpathYPointsStdDev" csv:"checkboxpathYPointsStdDev"`
	CheckBoxPathYPointsSkew                           float64 `arff:"checkboxpathYPointsSkew" csv:"checkboxpathYPointsSkew"`
	CheckBoxPathPairwiseXVelocityMin                  float64 `arff:"checkboxpathPairwiseXVelocityMin" csv:"checkboxpathPairwiseXVelocityMin"`
	CheckBoxPathPairwiseXVelocityMax                  float64 `arff:"checkboxpathPairwiseXVelocityMax" csv:"checkboxpathPairwiseXVelocityMax"`
	CheckBoxPathPairwiseXVelocityDiff                 float64 `arff:"checkboxpathPairwiseXVelocityDiff" csv:"checkboxpathPairwiseXVelocityDiff"`
	CheckBoxPathPairwiseXVelocitySum                  float64 `arff:"checkboxpathPairwiseXVelocitySum" csv:"checkboxpathPairwiseXVelocitySum"`
	CheckBoxPathPairwiseXVelocityMean                 float64 `arff:"checkboxpathPairwiseXVelocityMean" csv:"checkboxpathPairwiseXVelocityMean"`
	CheckBoxPathPairwiseXVelocityStdDev               float64 `arff:"checkboxpathPairwiseXVelocityStdDev" csv:"checkboxpathPairwiseXVelocityStdDev"`
	CheckBoxPathPairwiseXVelocitySkew                 float64 `arff:"checkboxpathPairwiseXVelocitySkew" csv:"checkboxpathPairwiseXVelocitySkew"`
	CheckBoxPathPairwiseYVelocityMin                  float64 `arff:"checkboxpathPairwiseYVelocityMin" csv:"checkboxpathPairwiseYVelocityMin"`
	CheckBoxPathPairwiseYVelocityMax                  float64 `arff:"checkboxpathPairwiseYVelocityMax" csv:"checkboxpathPairwiseYVelocityMax"`
	CheckBoxPathPairwiseYVelocityDiff                 float64 `arff:"checkboxpathPairwiseYVelocityDiff" csv:"checkboxpathPairwiseYVelocityDiff"`
	CheckBoxPathPairwiseYVelocitySum                  float64 `arff:"checkboxpathPairwiseYVelocitySum" csv:"checkboxpathPairwiseYVelocitySum"`
	CheckBoxPathPairwiseYVelocityMean                 float64 `arff:"checkboxpathPairwiseYVelocityMean" csv:"checkboxpathPairwiseYVelocityMean"`
	CheckBoxPathPairwiseYVelocityStdDev               float64 `arff:"checkboxpathPairwiseYVelocityStdDev" csv:"checkboxpathPairwiseYVelocityStdDev"`
	CheckBoxPathPairwiseYVelocitySkew                 float64 `arff:"checkboxpathPairwiseYVelocitySkew" csv:"checkboxpathPairwiseYVelocitySkew"`
	CheckBoxPathPairwiseXDistanceMin                  float64 `arff:"checkboxpathPairwiseXDistanceMin" csv:"checkboxpathPairwiseXDistanceMin"`
	CheckBoxPathPairwiseXDistanceMax                  float64 `arff:"checkboxpathPairwiseXDistanceMax" csv:"checkboxpathPairwiseXDistanceMax"`
	CheckBoxPathPairwiseXDistanceDiff                 float64 `arff:"checkboxpathPairwiseXDistanceDiff" csv:"checkboxpathPairwiseXDistanceDiff"`
	CheckBoxPathPairwiseXDistanceSum                  float64 `arff:"checkboxpathPairwiseXDistanceSum" csv:"checkboxpathPairwiseXDistanceSum"`
	CheckBoxPathPairwiseXDistanceMean                 float64 `arff:"checkboxpathPairwiseXDistanceMean" csv:"checkboxpathPairwiseXDistanceMean"`
	CheckBoxPathPairwiseXDistanceStdDev               float64 `arff:"checkboxpathPairwiseXDistanceStdDev" csv:"checkboxpathPairwiseXDistanceStdDev"`
	CheckBoxPathPairwiseXDistanceSkew                 float64 `arff:"checkboxpathPairwiseXDistanceSkew" csv:"checkboxpathPairwiseXDistanceSkew"`
	CheckBoxPathPairwiseYDistanceMin                  float64 `arff:"checkboxpathPairwiseYDistanceMin" csv:"checkboxpathPairwiseYDistanceMin"`
	CheckBoxPathPairwiseYDistanceMax                  float64 `arff:"checkboxpathPairwiseYDistanceMax" csv:"checkboxpathPairwiseYDistanceMax"`
	CheckBoxPathPairwiseYDistanceDiff                 float64 `arff:"checkboxpathPairwiseYDistanceDiff" csv:"checkboxpathPairwiseYDistanceDiff"`
	CheckBoxPathPairwiseYDistanceSum                  float64 `arff:"checkboxpathPairwiseYDistanceSum" csv:"checkboxpathPairwiseYDistanceSum"`
	CheckBoxPathPairwiseYDistanceMean                 float64 `arff:"checkboxpathPairwiseYDistanceMean" csv:"checkboxpathPairwiseYDistanceMean"`
	CheckBoxPathPairwiseYDistanceStdDev               float64 `arff:"checkboxpathPairwiseYDistanceStdDev" csv:"checkboxpathPairwiseYDistanceStdDev"`
	CheckBoxPathPairwiseYDistanceSkew                 float64 `arff:"checkboxpathPairwiseYDistanceSkew" csv:"checkboxpathPairwiseYDistanceSkew"`
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
	//
	// Movement Features
	//
	header.AddAttr("numberOfPaths", arff.Numeric, nil)
	// CheckBoxPath
	header.AddAttr("checkBoxPathStraightness", arff.Numeric, nil)
	header.AddAttr("checkBoxPathNoScrolls", arff.Numeric, nil)
	header.AddAttr("checkBoxPathNoMiddleClick", arff.Numeric, nil)
	header.AddAttr("checkBoxPathNoRightClick", arff.Numeric, nil)
	header.AddAttr("checkBoxPathBreakRatio", arff.Numeric, nil)
	header.AddAttr("checkBoxPathNoBreaks", arff.Numeric, nil)
	header.AddAttr("checkBoxPathTimeMovementClick", arff.Numeric, nil)
	header.AddAttr("checkBoxPathTimeClickMovement", arff.Numeric, nil)
	header.AddAttr("checkBoxPathDurPath", arff.Numeric, nil)
	header.AddAttr("checkBoxPathAngleStartEnd", arff.Numeric, nil)
	header.AddAttr("checkBoxPathNumMovementPoints", arff.Numeric, nil)
	header.AddAttr("checkBoxPathDistStartEnd", arff.Numeric, nil)
	// Auto-Generated
	//
	header.AddAttr("checkboxpathPairwiseDistanceMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDistanceMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDistanceDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDistanceSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDistanceMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDistanceStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDistanceSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseVelocityMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseVelocityMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseVelocityDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseVelocitySum", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseVelocityMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseVelocitySkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAccelerationMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAccelerationMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAccelerationDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAccelerationSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAccelerationMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAccelerationStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAccelerationSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngleMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngleMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngleDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngleSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngleMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngleStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngleSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathAngleBetweenMovementAndStartEndMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathAngleBetweenMovementAndStartEndMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathAngleBetweenMovementAndStartEndDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathAngleBetweenMovementAndStartEndSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathAngleBetweenMovementAndStartEndMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathAngleBetweenMovementAndStartEndStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathAngleBetweenMovementAndStartEndSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngularVelocityMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngularVelocityMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngularVelocityDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngularVelocitySum", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngularVelocityMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngularVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseAngularVelocitySkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDurationMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDurationMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDurationDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDurationSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDurationMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDurationStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseDurationSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathTimeBetweenClickAndReleaseMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathTimeBetweenClickAndReleaseMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathTimeBetweenClickAndReleaseDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathTimeBetweenClickAndReleaseSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathTimeBetweenClickAndReleaseMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathTimeBetweenClickAndReleaseStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathTimeBetweenClickAndReleaseSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathBreakTimesMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathBreakTimesMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathBreakTimesDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathBreakTimesSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathBreakTimesMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathBreakTimesStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathBreakTimesSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickDistanceMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickDistanceMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickDistanceDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickDistanceSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickDistanceMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickDistanceStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickDistanceSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickVelocityMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickVelocityMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickVelocityDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickVelocitySum", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickVelocityMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickVelocitySkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAccelerationMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAccelerationMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAccelerationDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAccelerationSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAccelerationMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAccelerationStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAccelerationSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAngleMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAngleMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAngleDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAngleSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAngleMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAngleStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathMovementDuringClickAngleSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDXMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDXMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDXDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDXSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDXMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDXStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDXSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDYMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDYMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDYDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDYSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDYMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDYStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDYSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDZMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDZMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDZDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDZSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDZMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDZStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDZSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDMMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDMMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDMDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDMSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDMMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDMStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathScrollDMSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathXPointsMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathXPointsMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathXPointsDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathXPointsSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathXPointsMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathXPointsStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathXPointsSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathYPointsMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathYPointsMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathYPointsDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathYPointsSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathYPointsMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathYPointsStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathYPointsSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXVelocityMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXVelocityMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXVelocityDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXVelocitySum", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXVelocityMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXVelocitySkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYVelocityMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYVelocityMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYVelocityDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYVelocitySum", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYVelocityMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYVelocitySkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXDistanceMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXDistanceMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXDistanceDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXDistanceSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXDistanceMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXDistanceStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseXDistanceSkew", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYDistanceMin", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYDistanceMax", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYDistanceDiff", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYDistanceSum", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYDistanceMean", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYDistanceStdDev", arff.Numeric, nil)
	header.AddAttr("checkboxpathPairwiseYDistanceSkew", arff.Numeric, nil)
	// RestPaths

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
	features.AddCheckBoxFeatures(checkBoxPath.GetRawFeatures())

	features.NumberOfPaths = uint16(len(otherPaths)) + 1

	restPaths := make([]*PathFeatures, len(otherPaths))
	for i, path := range otherPaths {
		pathFeature := path.GetRawFeatures()
		restPaths[i] = pathFeature
	}
	features.AddRestPathFeatures(restPaths)
}

// AddCheckBoxFeatures adds the features of the CheckboxPath to the FeatureSet
func (features *ProcessedFeatures) AddCheckBoxFeatures(path *PathFeatures) {
	features.CheckboxPathDistanceStartEndPoint = path.DistanceStartEndPoint
	features.CheckboxPathNumberOfMovementPoints = path.NumberOfMovementPoints
	features.CheckboxPathAngleStartEndPoint = path.AngleStartEndPoint
	features.CheckboxPathDurationOfPath = path.DurationOfPath
	features.CheckboxPathTimeBetweenClickAndMovement = path.TimeBetweenClickAndMovement
	features.CheckboxPathTimeBetweeenMovementAndDownClick = path.TimeBetweeenMovementAndDownClick
	features.CheckboxPathNumberOfBreaks = path.NumberOfBreaks
	features.CheckboxPathBreakTimeTotalTimeRatio = path.BreakTimeTotalTimeRatio
	features.CheckboxPathNumberOfRightClicks = path.NumberOfRightClicks
	features.CheckboxPathNumberOfMiddleClicks = path.NumberOfMiddleClicks
	features.CheckboxPathNumberOfScrolls = path.NumberOfScrolls
	features.CheckboxPathStraightness = path.Straightness
	//features.GenerateCheckboxCalculation(path)
}

// AddRestPathFeatures adds the features of the rest paths to the FeatureSet
func (features *ProcessedFeatures) AddRestPathFeatures(paths []*PathFeatures) {}

// calcCheckBoxPathPairwiseDistanceFeatures calculates the features of the PairwiseDistance vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathPairwiseDistanceFeatures(path *PathFeatures) {
	features.CheckBoxPathPairwiseDistanceMin, features.CheckBoxPathPairwiseDistanceMean, features.CheckBoxPathPairwiseDistanceMax, features.CheckBoxPathPairwiseDistanceStdDev = rnd.StatBasic(path.PairwiseDistance, true)
	features.CheckBoxPathPairwiseDistanceSum = SumFloat(path.PairwiseDistance)
	if len(path.PairwiseDistance) == 1 {
		features.CheckBoxPathPairwiseDistanceMin, features.CheckBoxPathPairwiseDistanceMean, features.CheckBoxPathPairwiseDistanceMax = path.PairwiseDistance[0], path.PairwiseDistance[0], path.PairwiseDistance[0]
	}
	features.CheckBoxPathPairwiseDistanceDiff = features.CheckBoxPathPairwiseDistanceMax - features.CheckBoxPathPairwiseDistanceMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathPairwiseDistanceSkew, _ = rnd.StatMoments(path.PairwiseDistance)
}

// calcCheckBoxPathPairwiseVelocityFeatures calculates the features of the PairwiseVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathPairwiseVelocityFeatures(path *PathFeatures) {
	features.CheckBoxPathPairwiseVelocityMin, features.CheckBoxPathPairwiseVelocityMean, features.CheckBoxPathPairwiseVelocityMax, features.CheckBoxPathPairwiseVelocityStdDev = rnd.StatBasic(path.PairwiseVelocity, true)
	features.CheckBoxPathPairwiseVelocitySum = SumFloat(path.PairwiseVelocity)
	if len(path.PairwiseVelocity) == 1 {
		features.CheckBoxPathPairwiseVelocityMin, features.CheckBoxPathPairwiseVelocityMean, features.CheckBoxPathPairwiseVelocityMax = path.PairwiseVelocity[0], path.PairwiseVelocity[0], path.PairwiseVelocity[0]
	}
	features.CheckBoxPathPairwiseVelocityDiff = features.CheckBoxPathPairwiseVelocityMax - features.CheckBoxPathPairwiseVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathPairwiseVelocitySkew, _ = rnd.StatMoments(path.PairwiseVelocity)
}

// calcCheckBoxPathPairwiseAccelerationFeatures calculates the features of the PairwiseAcceleration vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathPairwiseAccelerationFeatures(path *PathFeatures) {
	features.CheckBoxPathPairwiseAccelerationMin, features.CheckBoxPathPairwiseAccelerationMean, features.CheckBoxPathPairwiseAccelerationMax, features.CheckBoxPathPairwiseAccelerationStdDev = rnd.StatBasic(path.PairwiseAcceleration, true)
	features.CheckBoxPathPairwiseAccelerationSum = SumFloat(path.PairwiseAcceleration)
	if len(path.PairwiseAcceleration) == 1 {
		features.CheckBoxPathPairwiseAccelerationMin, features.CheckBoxPathPairwiseAccelerationMean, features.CheckBoxPathPairwiseAccelerationMax = path.PairwiseAcceleration[0], path.PairwiseAcceleration[0], path.PairwiseAcceleration[0]
	}
	features.CheckBoxPathPairwiseAccelerationDiff = features.CheckBoxPathPairwiseAccelerationMax - features.CheckBoxPathPairwiseAccelerationMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathPairwiseAccelerationSkew, _ = rnd.StatMoments(path.PairwiseAcceleration)
}

// calcCheckBoxPathPairwiseAngleFeatures calculates the features of the PairwiseAngle vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathPairwiseAngleFeatures(path *PathFeatures) {
	features.CheckBoxPathPairwiseAngleMin, features.CheckBoxPathPairwiseAngleMean, features.CheckBoxPathPairwiseAngleMax, features.CheckBoxPathPairwiseAngleStdDev = rnd.StatBasic(path.PairwiseAngle, true)
	features.CheckBoxPathPairwiseAngleSum = SumFloat(path.PairwiseAngle)
	if len(path.PairwiseAngle) == 1 {
		features.CheckBoxPathPairwiseAngleMin, features.CheckBoxPathPairwiseAngleMean, features.CheckBoxPathPairwiseAngleMax = path.PairwiseAngle[0], path.PairwiseAngle[0], path.PairwiseAngle[0]
	}
	features.CheckBoxPathPairwiseAngleDiff = features.CheckBoxPathPairwiseAngleMax - features.CheckBoxPathPairwiseAngleMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathPairwiseAngleSkew, _ = rnd.StatMoments(path.PairwiseAngle)
}

// calcCheckBoxPathAngleBetweenMovementAndStartEndFeatures calculates the features of the AngleBetweenMovementAndStartEnd vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathAngleBetweenMovementAndStartEndFeatures(path *PathFeatures) {
	features.CheckBoxPathAngleBetweenMovementAndStartEndMin, features.CheckBoxPathAngleBetweenMovementAndStartEndMean, features.CheckBoxPathAngleBetweenMovementAndStartEndMax, features.CheckBoxPathAngleBetweenMovementAndStartEndStdDev = rnd.StatBasic(path.AngleBetweenMovementAndStartEnd, true)
	features.CheckBoxPathAngleBetweenMovementAndStartEndSum = SumFloat(path.AngleBetweenMovementAndStartEnd)
	if len(path.AngleBetweenMovementAndStartEnd) == 1 {
		features.CheckBoxPathAngleBetweenMovementAndStartEndMin, features.CheckBoxPathAngleBetweenMovementAndStartEndMean, features.CheckBoxPathAngleBetweenMovementAndStartEndMax = path.AngleBetweenMovementAndStartEnd[0], path.AngleBetweenMovementAndStartEnd[0], path.AngleBetweenMovementAndStartEnd[0]
	}
	features.CheckBoxPathAngleBetweenMovementAndStartEndDiff = features.CheckBoxPathAngleBetweenMovementAndStartEndMax - features.CheckBoxPathAngleBetweenMovementAndStartEndMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathAngleBetweenMovementAndStartEndSkew, _ = rnd.StatMoments(path.AngleBetweenMovementAndStartEnd)
}

// calcCheckBoxPathPairwiseAngularVelocityFeatures calculates the features of the PairwiseAngularVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathPairwiseAngularVelocityFeatures(path *PathFeatures) {
	features.CheckBoxPathPairwiseAngularVelocityMin, features.CheckBoxPathPairwiseAngularVelocityMean, features.CheckBoxPathPairwiseAngularVelocityMax, features.CheckBoxPathPairwiseAngularVelocityStdDev = rnd.StatBasic(path.PairwiseAngularVelocity, true)
	features.CheckBoxPathPairwiseAngularVelocitySum = SumFloat(path.PairwiseAngularVelocity)
	if len(path.PairwiseAngularVelocity) == 1 {
		features.CheckBoxPathPairwiseAngularVelocityMin, features.CheckBoxPathPairwiseAngularVelocityMean, features.CheckBoxPathPairwiseAngularVelocityMax = path.PairwiseAngularVelocity[0], path.PairwiseAngularVelocity[0], path.PairwiseAngularVelocity[0]
	}
	features.CheckBoxPathPairwiseAngularVelocityDiff = features.CheckBoxPathPairwiseAngularVelocityMax - features.CheckBoxPathPairwiseAngularVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathPairwiseAngularVelocitySkew, _ = rnd.StatMoments(path.PairwiseAngularVelocity)
}

// calcCheckBoxPathPairwiseDurationFeatures calculates the features of the PairwiseDuration vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathPairwiseDurationFeatures(path *PathFeatures) {
	features.CheckBoxPathPairwiseDurationMax = MaxUint64(path.PairwiseDuration)
	features.CheckBoxPathPairwiseDurationMin = MinUint64(path.PairwiseDuration)
	features.CheckBoxPathPairwiseDurationDiff = features.CheckBoxPathPairwiseDurationMax - features.CheckBoxPathPairwiseDurationMin
	features.CheckBoxPathPairwiseDurationSum = SumUint64(path.PairwiseDuration)
	features.CheckBoxPathPairwiseDurationMean = Mean(ConvertUint64ToFloat64(path.PairwiseDuration), float64(features.CheckBoxPathPairwiseDurationSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.CheckBoxPathPairwiseDurationStdDev, features.CheckBoxPathPairwiseDurationSkew, _ = rnd.StatMoments(ConvertUint64ToFloat64(path.PairwiseDuration))
}

// calcCheckBoxPathTimeBetweenClickAndReleaseFeatures calculates the features of the TimeBetweenClickAndRelease vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathTimeBetweenClickAndReleaseFeatures(path *PathFeatures) {
	features.CheckBoxPathTimeBetweenClickAndReleaseMax = MaxUint64(path.TimeBetweenClickAndRelease)
	features.CheckBoxPathTimeBetweenClickAndReleaseMin = MinUint64(path.TimeBetweenClickAndRelease)
	features.CheckBoxPathTimeBetweenClickAndReleaseDiff = features.CheckBoxPathTimeBetweenClickAndReleaseMax - features.CheckBoxPathTimeBetweenClickAndReleaseMin
	features.CheckBoxPathTimeBetweenClickAndReleaseSum = SumUint64(path.TimeBetweenClickAndRelease)
	features.CheckBoxPathTimeBetweenClickAndReleaseMean = Mean(ConvertUint64ToFloat64(path.TimeBetweenClickAndRelease), float64(features.CheckBoxPathTimeBetweenClickAndReleaseSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.CheckBoxPathTimeBetweenClickAndReleaseStdDev, features.CheckBoxPathTimeBetweenClickAndReleaseSkew, _ = rnd.StatMoments(ConvertUint64ToFloat64(path.TimeBetweenClickAndRelease))
}

// calcCheckBoxPathBreakTimesFeatures calculates the features of the BreakTimes vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathBreakTimesFeatures(path *PathFeatures) {
	features.CheckBoxPathBreakTimesMax = MaxUint64(path.BreakTimes)
	features.CheckBoxPathBreakTimesMin = MinUint64(path.BreakTimes)
	features.CheckBoxPathBreakTimesDiff = features.CheckBoxPathBreakTimesMax - features.CheckBoxPathBreakTimesMin
	features.CheckBoxPathBreakTimesSum = SumUint64(path.BreakTimes)
	features.CheckBoxPathBreakTimesMean = Mean(ConvertUint64ToFloat64(path.BreakTimes), float64(features.CheckBoxPathBreakTimesSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.CheckBoxPathBreakTimesStdDev, features.CheckBoxPathBreakTimesSkew, _ = rnd.StatMoments(ConvertUint64ToFloat64(path.BreakTimes))
}

// calcCheckBoxPathMovementDuringClickDistanceFeatures calculates the features of the MovementDuringClickDistance vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathMovementDuringClickDistanceFeatures(path *PathFeatures) {
	features.CheckBoxPathMovementDuringClickDistanceMin, features.CheckBoxPathMovementDuringClickDistanceMean, features.CheckBoxPathMovementDuringClickDistanceMax, features.CheckBoxPathMovementDuringClickDistanceStdDev = rnd.StatBasic(path.MovementDuringClickDistance, true)
	features.CheckBoxPathMovementDuringClickDistanceSum = SumFloat(path.MovementDuringClickDistance)
	if len(path.MovementDuringClickDistance) == 1 {
		features.CheckBoxPathMovementDuringClickDistanceMin, features.CheckBoxPathMovementDuringClickDistanceMean, features.CheckBoxPathMovementDuringClickDistanceMax = path.MovementDuringClickDistance[0], path.MovementDuringClickDistance[0], path.MovementDuringClickDistance[0]
	}
	features.CheckBoxPathMovementDuringClickDistanceDiff = features.CheckBoxPathMovementDuringClickDistanceMax - features.CheckBoxPathMovementDuringClickDistanceMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathMovementDuringClickDistanceSkew, _ = rnd.StatMoments(path.MovementDuringClickDistance)
}

// calcCheckBoxPathMovementDuringClickVelocityFeatures calculates the features of the MovementDuringClickVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathMovementDuringClickVelocityFeatures(path *PathFeatures) {
	features.CheckBoxPathMovementDuringClickVelocityMin, features.CheckBoxPathMovementDuringClickVelocityMean, features.CheckBoxPathMovementDuringClickVelocityMax, features.CheckBoxPathMovementDuringClickVelocityStdDev = rnd.StatBasic(path.MovementDuringClickVelocity, true)
	features.CheckBoxPathMovementDuringClickVelocitySum = SumFloat(path.MovementDuringClickVelocity)
	if len(path.MovementDuringClickVelocity) == 1 {
		features.CheckBoxPathMovementDuringClickVelocityMin, features.CheckBoxPathMovementDuringClickVelocityMean, features.CheckBoxPathMovementDuringClickVelocityMax = path.MovementDuringClickVelocity[0], path.MovementDuringClickVelocity[0], path.MovementDuringClickVelocity[0]
	}
	features.CheckBoxPathMovementDuringClickVelocityDiff = features.CheckBoxPathMovementDuringClickVelocityMax - features.CheckBoxPathMovementDuringClickVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathMovementDuringClickVelocitySkew, _ = rnd.StatMoments(path.MovementDuringClickVelocity)
}

// calcCheckBoxPathMovementDuringClickAccelerationFeatures calculates the features of the MovementDuringClickAcceleration vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathMovementDuringClickAccelerationFeatures(path *PathFeatures) {
	features.CheckBoxPathMovementDuringClickAccelerationMin, features.CheckBoxPathMovementDuringClickAccelerationMean, features.CheckBoxPathMovementDuringClickAccelerationMax, features.CheckBoxPathMovementDuringClickAccelerationStdDev = rnd.StatBasic(path.MovementDuringClickAcceleration, true)
	features.CheckBoxPathMovementDuringClickAccelerationSum = SumFloat(path.MovementDuringClickAcceleration)
	if len(path.MovementDuringClickAcceleration) == 1 {
		features.CheckBoxPathMovementDuringClickAccelerationMin, features.CheckBoxPathMovementDuringClickAccelerationMean, features.CheckBoxPathMovementDuringClickAccelerationMax = path.MovementDuringClickAcceleration[0], path.MovementDuringClickAcceleration[0], path.MovementDuringClickAcceleration[0]
	}
	features.CheckBoxPathMovementDuringClickAccelerationDiff = features.CheckBoxPathMovementDuringClickAccelerationMax - features.CheckBoxPathMovementDuringClickAccelerationMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathMovementDuringClickAccelerationSkew, _ = rnd.StatMoments(path.MovementDuringClickAcceleration)
}

// calcCheckBoxPathMovementDuringClickAngleFeatures calculates the features of the MovementDuringClickAngle vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathMovementDuringClickAngleFeatures(path *PathFeatures) {
	features.CheckBoxPathMovementDuringClickAngleMin, features.CheckBoxPathMovementDuringClickAngleMean, features.CheckBoxPathMovementDuringClickAngleMax, features.CheckBoxPathMovementDuringClickAngleStdDev = rnd.StatBasic(path.MovementDuringClickAngle, true)
	features.CheckBoxPathMovementDuringClickAngleSum = SumFloat(path.MovementDuringClickAngle)
	if len(path.MovementDuringClickAngle) == 1 {
		features.CheckBoxPathMovementDuringClickAngleMin, features.CheckBoxPathMovementDuringClickAngleMean, features.CheckBoxPathMovementDuringClickAngleMax = path.MovementDuringClickAngle[0], path.MovementDuringClickAngle[0], path.MovementDuringClickAngle[0]
	}
	features.CheckBoxPathMovementDuringClickAngleDiff = features.CheckBoxPathMovementDuringClickAngleMax - features.CheckBoxPathMovementDuringClickAngleMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathMovementDuringClickAngleSkew, _ = rnd.StatMoments(path.MovementDuringClickAngle)
}

// calcCheckBoxPathScrollDXFeatures calculates the features of the ScrollDX vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathScrollDXFeatures(path *PathFeatures) {
	features.CheckBoxPathScrollDXMin, features.CheckBoxPathScrollDXMean, features.CheckBoxPathScrollDXMax, features.CheckBoxPathScrollDXStdDev = rnd.StatBasic(path.ScrollDX, true)
	features.CheckBoxPathScrollDXSum = SumFloat(path.ScrollDX)
	if len(path.ScrollDX) == 1 {
		features.CheckBoxPathScrollDXMin, features.CheckBoxPathScrollDXMean, features.CheckBoxPathScrollDXMax = path.ScrollDX[0], path.ScrollDX[0], path.ScrollDX[0]
	}
	features.CheckBoxPathScrollDXDiff = features.CheckBoxPathScrollDXMax - features.CheckBoxPathScrollDXMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathScrollDXSkew, _ = rnd.StatMoments(path.ScrollDX)
}

// calcCheckBoxPathScrollDYFeatures calculates the features of the ScrollDY vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathScrollDYFeatures(path *PathFeatures) {
	features.CheckBoxPathScrollDYMin, features.CheckBoxPathScrollDYMean, features.CheckBoxPathScrollDYMax, features.CheckBoxPathScrollDYStdDev = rnd.StatBasic(path.ScrollDY, true)
	features.CheckBoxPathScrollDYSum = SumFloat(path.ScrollDY)
	if len(path.ScrollDY) == 1 {
		features.CheckBoxPathScrollDYMin, features.CheckBoxPathScrollDYMean, features.CheckBoxPathScrollDYMax = path.ScrollDY[0], path.ScrollDY[0], path.ScrollDY[0]
	}
	features.CheckBoxPathScrollDYDiff = features.CheckBoxPathScrollDYMax - features.CheckBoxPathScrollDYMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathScrollDYSkew, _ = rnd.StatMoments(path.ScrollDY)
}

// calcCheckBoxPathScrollDZFeatures calculates the features of the ScrollDZ vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathScrollDZFeatures(path *PathFeatures) {
	features.CheckBoxPathScrollDZMin, features.CheckBoxPathScrollDZMean, features.CheckBoxPathScrollDZMax, features.CheckBoxPathScrollDZStdDev = rnd.StatBasic(path.ScrollDZ, true)
	features.CheckBoxPathScrollDZSum = SumFloat(path.ScrollDZ)
	if len(path.ScrollDZ) == 1 {
		features.CheckBoxPathScrollDZMin, features.CheckBoxPathScrollDZMean, features.CheckBoxPathScrollDZMax = path.ScrollDZ[0], path.ScrollDZ[0], path.ScrollDZ[0]
	}
	features.CheckBoxPathScrollDZDiff = features.CheckBoxPathScrollDZMax - features.CheckBoxPathScrollDZMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathScrollDZSkew, _ = rnd.StatMoments(path.ScrollDZ)
}

// calcCheckBoxPathScrollDMFeatures calculates the features of the ScrollDM vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathScrollDMFeatures(path *PathFeatures) {
	features.CheckBoxPathScrollDMMax = MaxUint8(path.ScrollDM)
	features.CheckBoxPathScrollDMMin = MinUint8(path.ScrollDM)
	features.CheckBoxPathScrollDMDiff = features.CheckBoxPathScrollDMMax - features.CheckBoxPathScrollDMMin
	features.CheckBoxPathScrollDMSum = SumUint8(path.ScrollDM)
	features.CheckBoxPathScrollDMMean = Mean(ConvertUint8ToFloat64(path.ScrollDM), float64(features.CheckBoxPathScrollDMSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.CheckBoxPathScrollDMStdDev, features.CheckBoxPathScrollDMSkew, _ = rnd.StatMoments(ConvertUint8ToFloat64(path.ScrollDM))
}

// calcCheckBoxPathXPointsFeatures calculates the features of the XPoints vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathXPointsFeatures(path *PathFeatures) {
	features.CheckBoxPathXPointsMin, features.CheckBoxPathXPointsMean, features.CheckBoxPathXPointsMax, features.CheckBoxPathXPointsStdDev = rnd.StatBasic(path.XPoints, true)
	features.CheckBoxPathXPointsSum = SumFloat(path.XPoints)
	if len(path.XPoints) == 1 {
		features.CheckBoxPathXPointsMin, features.CheckBoxPathXPointsMean, features.CheckBoxPathXPointsMax = path.XPoints[0], path.XPoints[0], path.XPoints[0]
	}
	features.CheckBoxPathXPointsDiff = features.CheckBoxPathXPointsMax - features.CheckBoxPathXPointsMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathXPointsSkew, _ = rnd.StatMoments(path.XPoints)
}

// calcCheckBoxPathYPointsFeatures calculates the features of the YPoints vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathYPointsFeatures(path *PathFeatures) {
	features.CheckBoxPathYPointsMin, features.CheckBoxPathYPointsMean, features.CheckBoxPathYPointsMax, features.CheckBoxPathYPointsStdDev = rnd.StatBasic(path.YPoints, true)
	features.CheckBoxPathYPointsSum = SumFloat(path.YPoints)
	if len(path.YPoints) == 1 {
		features.CheckBoxPathYPointsMin, features.CheckBoxPathYPointsMean, features.CheckBoxPathYPointsMax = path.YPoints[0], path.YPoints[0], path.YPoints[0]
	}
	features.CheckBoxPathYPointsDiff = features.CheckBoxPathYPointsMax - features.CheckBoxPathYPointsMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathYPointsSkew, _ = rnd.StatMoments(path.YPoints)
}

// calcCheckBoxPathPairwiseXVelocityFeatures calculates the features of the PairwiseXVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathPairwiseXVelocityFeatures(path *PathFeatures) {
	features.CheckBoxPathPairwiseXVelocityMin, features.CheckBoxPathPairwiseXVelocityMean, features.CheckBoxPathPairwiseXVelocityMax, features.CheckBoxPathPairwiseXVelocityStdDev = rnd.StatBasic(path.PairwiseXVelocity, true)
	features.CheckBoxPathPairwiseXVelocitySum = SumFloat(path.PairwiseXVelocity)
	if len(path.PairwiseXVelocity) == 1 {
		features.CheckBoxPathPairwiseXVelocityMin, features.CheckBoxPathPairwiseXVelocityMean, features.CheckBoxPathPairwiseXVelocityMax = path.PairwiseXVelocity[0], path.PairwiseXVelocity[0], path.PairwiseXVelocity[0]
	}
	features.CheckBoxPathPairwiseXVelocityDiff = features.CheckBoxPathPairwiseXVelocityMax - features.CheckBoxPathPairwiseXVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathPairwiseXVelocitySkew, _ = rnd.StatMoments(path.PairwiseXVelocity)
}

// calcCheckBoxPathPairwiseYVelocityFeatures calculates the features of the PairwiseYVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathPairwiseYVelocityFeatures(path *PathFeatures) {
	features.CheckBoxPathPairwiseYVelocityMin, features.CheckBoxPathPairwiseYVelocityMean, features.CheckBoxPathPairwiseYVelocityMax, features.CheckBoxPathPairwiseYVelocityStdDev = rnd.StatBasic(path.PairwiseYVelocity, true)
	features.CheckBoxPathPairwiseYVelocitySum = SumFloat(path.PairwiseYVelocity)
	if len(path.PairwiseYVelocity) == 1 {
		features.CheckBoxPathPairwiseYVelocityMin, features.CheckBoxPathPairwiseYVelocityMean, features.CheckBoxPathPairwiseYVelocityMax = path.PairwiseYVelocity[0], path.PairwiseYVelocity[0], path.PairwiseYVelocity[0]
	}
	features.CheckBoxPathPairwiseYVelocityDiff = features.CheckBoxPathPairwiseYVelocityMax - features.CheckBoxPathPairwiseYVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathPairwiseYVelocitySkew, _ = rnd.StatMoments(path.PairwiseYVelocity)
}

// calcCheckBoxPathPairwiseXDistanceFeatures calculates the features of the PairwiseXDistance vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathPairwiseXDistanceFeatures(path *PathFeatures) {
	features.CheckBoxPathPairwiseXDistanceMin, features.CheckBoxPathPairwiseXDistanceMean, features.CheckBoxPathPairwiseXDistanceMax, features.CheckBoxPathPairwiseXDistanceStdDev = rnd.StatBasic(path.PairwiseXDistance, true)
	features.CheckBoxPathPairwiseXDistanceSum = SumFloat(path.PairwiseXDistance)
	if len(path.PairwiseXDistance) == 1 {
		features.CheckBoxPathPairwiseXDistanceMin, features.CheckBoxPathPairwiseXDistanceMean, features.CheckBoxPathPairwiseXDistanceMax = path.PairwiseXDistance[0], path.PairwiseXDistance[0], path.PairwiseXDistance[0]
	}
	features.CheckBoxPathPairwiseXDistanceDiff = features.CheckBoxPathPairwiseXDistanceMax - features.CheckBoxPathPairwiseXDistanceMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathPairwiseXDistanceSkew, _ = rnd.StatMoments(path.PairwiseXDistance)
}

// calcCheckBoxPathPairwiseYDistanceFeatures calculates the features of the PairwiseYDistance vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcCheckBoxPathPairwiseYDistanceFeatures(path *PathFeatures) {
	features.CheckBoxPathPairwiseYDistanceMin, features.CheckBoxPathPairwiseYDistanceMean, features.CheckBoxPathPairwiseYDistanceMax, features.CheckBoxPathPairwiseYDistanceStdDev = rnd.StatBasic(path.PairwiseYDistance, true)
	features.CheckBoxPathPairwiseYDistanceSum = SumFloat(path.PairwiseYDistance)
	if len(path.PairwiseYDistance) == 1 {
		features.CheckBoxPathPairwiseYDistanceMin, features.CheckBoxPathPairwiseYDistanceMean, features.CheckBoxPathPairwiseYDistanceMax = path.PairwiseYDistance[0], path.PairwiseYDistance[0], path.PairwiseYDistance[0]
	}
	features.CheckBoxPathPairwiseYDistanceDiff = features.CheckBoxPathPairwiseYDistanceMax - features.CheckBoxPathPairwiseYDistanceMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.CheckBoxPathPairwiseYDistanceSkew, _ = rnd.StatMoments(path.PairwiseYDistance)
}

// GenerateCheckboxCalculation calculates the features from the checkbox movement path; Auto-generated
func (features *ProcessedFeatures) GenerateCheckboxCalculation(path *PathFeatures) {
	features.calcCheckBoxPathPairwiseDistanceFeatures(path)
	features.calcCheckBoxPathPairwiseVelocityFeatures(path)
	features.calcCheckBoxPathPairwiseAccelerationFeatures(path)
	features.calcCheckBoxPathPairwiseAngleFeatures(path)
	features.calcCheckBoxPathAngleBetweenMovementAndStartEndFeatures(path)
	features.calcCheckBoxPathPairwiseAngularVelocityFeatures(path)
	features.calcCheckBoxPathPairwiseDurationFeatures(path)
	features.calcCheckBoxPathTimeBetweenClickAndReleaseFeatures(path)
	features.calcCheckBoxPathBreakTimesFeatures(path)
	features.calcCheckBoxPathMovementDuringClickDistanceFeatures(path)
	features.calcCheckBoxPathMovementDuringClickVelocityFeatures(path)
	features.calcCheckBoxPathMovementDuringClickAccelerationFeatures(path)
	features.calcCheckBoxPathMovementDuringClickAngleFeatures(path)
	features.calcCheckBoxPathScrollDXFeatures(path)
	features.calcCheckBoxPathScrollDYFeatures(path)
	features.calcCheckBoxPathScrollDZFeatures(path)
	features.calcCheckBoxPathScrollDMFeatures(path)
	features.calcCheckBoxPathXPointsFeatures(path)
	features.calcCheckBoxPathYPointsFeatures(path)
	features.calcCheckBoxPathPairwiseXVelocityFeatures(path)
	features.calcCheckBoxPathPairwiseYVelocityFeatures(path)
	features.calcCheckBoxPathPairwiseXDistanceFeatures(path)
	features.calcCheckBoxPathPairwiseYDistanceFeatures(path)
}
