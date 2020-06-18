package data

import (
	"math"
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
	// Plugins
	PluginChromePDF         uint8 `arff:"plugin-chrome-pdf" csv:"plugin-chrome-pdf"`
	PluginChromePDFViewer   uint8 `arff:"plugin-chrome-pdf-viewer" csv:"plugin-chrome-pdf-viewer"`
	PluginChromiumPDF       uint8 `arff:"plugin-chromium-pdf" csv:"plugin-chromium-pdf"`
	PluginChromiumPDFViewer uint8 `arff:"plugin-chromium-pdf-viewer" csv:"plugin-chromium-pdf-viewer"`
	PluginNativeClient      uint8 `arff:"plugin-native-client" csv:"plugin-native-client"`
	PluginMore              uint8 `arff:"plugin-more" csv:"plugin-more"`
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
	WebGLGCModel                 string `arff:"webglGCModel" csv:"webglGCModel"`
	WebGLGCVendor                string `arff:"webglGCVendor" csv:"webglGCVendor"`
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
	// Task Data
	AmountFailedTasks                  uint8   `arff:"failedTasks" csv:"failedTasks"`
	AmountTasks                        uint8   `arff:"taskAmount" csv:"taskAmount"`
	AmountNo                           uint8   `arff:"noCaptchaAmount" csv:"noCaptchaAmount"`
	AmountSimImg                       uint8   `arff:"similarImageAmount" csv:"similarImageAmount"`
	AmountDynSimImg                    uint8   `arff:"dynamicSimilarImageAmount" csv:"dynamicSimilarImageAmount"`
	AmountObjIdent                     uint8   `arff:"objectIdentificationAmount" csv:"objectIdentificationAmount"`
	AmountCandidatesStopSign           uint8   `arff:"candidatesStopSign" csv:"candidatesStopSign"`
	AmountCandidatesSpeedLimit         uint8   `arff:"candidatesSpeedLimit" csv:"candidatesSpeedLimit"`
	AmountCandidatesStreetName         uint8   `arff:"candidatesStreetName" csv:"candidatesStreetName"`
	AmountCandidatesOther              uint8   `arff:"candidatesOther" csv:"candidatesOther"`
	AmountCandidatesCar                uint8   `arff:"candidatesCar" csv:"candidatesCar"`
	AmountCandidatesBridge             uint8   `arff:"candidatesBridge" csv:"candidatesBridge"`
	AmountCandidatesRoad               uint8   `arff:"candidatesRoad" csv:"candidatesRoad"`
	AmountCandidatesBoundingbox        uint8   `arff:"candidatesBoundingBox" csv:"candidatesBoundingBox"`
	DurationLastCellEventToStop        uint64  `arff:"durationLastCellEventToStop" csv:"durationLastCellEventToStop"`
	TaskPicturesSelectedMin            uint16  `arff:"taskPicturesSelectedMin" csv:"taskPicturesSelectedMin"`
	TaskPicturesSelectedMax            uint16  `arff:"taskPicturesSelectedMax" csv:"taskPicturesSelectedMax"`
	TaskPicturesSelectedDiff           uint16  `arff:"taskPicturesSelectedDiff" csv:"taskPicturesSelectedDiff"`
	TaskPicturesSelectedSum            uint16  `arff:"taskPicturesSelectedSum" csv:"taskPicturesSelectedSum"`
	TaskPicturesSelectedMean           float64 `arff:"taskPicturesSelectedMean" csv:"taskPicturesSelectedMean"`
	TaskPicturesSelectedStdDev         float64 `arff:"taskPicturesSelectedStdDev" csv:"taskPicturesSelectedStdDev"`
	TaskPicturesSelectedSkew           float64 `arff:"taskPicturesSelectedSkew" csv:"taskPicturesSelectedSkew"`
	AmountCellEventsMin                uint16  `arff:"cellEventsMin" csv:"cellEventsMin"`
	AmountCellEventsMax                uint16  `arff:"cellEventsMax" csv:"cellEventsMax"`
	AmountCellEventsDiff               uint16  `arff:"cellEventsDiff" csv:"cellEventsDiff"`
	AmountCellEventsSum                uint16  `arff:"cellEventsSum" csv:"cellEventsSum"`
	AmountCellEventsMean               float64 `arff:"cellEventsMean" csv:"cellEventsMean"`
	AmountCellEventsStdDev             float64 `arff:"cellEventsStdDev" csv:"cellEventsStdDev"`
	AmountCellEventsSkew               float64 `arff:"cellEventsSkew" csv:"cellEventsSkew"`
	TaskDurationFirstLastEventMin      uint64  `arff:"taskDurationFirstLastEventMin" csv:"taskDurationFirstLastEventMin"`
	TaskDurationFirstLastEventMax      uint64  `arff:"taskDurationFirstLastEventMax" csv:"taskDurationFirstLastEventMax"`
	TaskDurationFirstLastEventDiff     uint64  `arff:"taskDurationFirstLastEventDiff" csv:"taskDurationFirstLastEventDiff"`
	TaskDurationFirstLastEventSum      uint64  `arff:"taskDurationFirstLastEventSum" csv:"taskDurationFirstLastEventSum"`
	TaskDurationFirstLastEventMean     float64 `arff:"taskDurationFirstLastEventMean" csv:"taskDurationFirstLastEventMean"`
	TaskDurationFirstLastEventStdDev   float64 `arff:"taskDurationFirstLastEventStdDev" csv:"taskDurationFirstLastEventStdDev"`
	TaskDurationFirstLastEventSkew     float64 `arff:"taskDurationFirstLastEventSkew" csv:"taskDurationFirstLastEventSkew"`
	TaskChangedSelectionsMin           uint16  `arff:"taskChangedSelectionsMin" csv:"taskChangedSelectionsMin"`
	TaskChangedSelectionsMax           uint16  `arff:"taskChangedSelectionsMax" csv:"taskChangedSelectionsMax"`
	TaskChangedSelectionsDiff          uint16  `arff:"taskChangedSelectionsDiff" csv:"taskChangedSelectionsDiff"`
	TaskChangedSelectionsSum           uint16  `arff:"taskChangedSelectionsSum" csv:"taskChangedSelectionsSum"`
	TaskChangedSelectionsMean          float64 `arff:"taskChangedSelectionsMean" csv:"taskChangedSelectionsMean"`
	TaskChangedSelectionsStdDev        float64 `arff:"taskChangedSelectionsStdDev" csv:"taskChangedSelectionsStdDev"`
	TaskChangedSelectionsSkew          float64 `arff:"taskChangedSelectionsSkew" csv:"taskChangedSelectionsSkew"`
	AmountDisappearingImagesMin        uint8   `arff:"amountDisappearingImagesMin" csv:"amountDisappearingImagesMin"`
	AmountDisappearingImagesMax        uint8   `arff:"amountDisappearingImagesMax" csv:"amountDisappearingImagesMax"`
	AmountDisappearingImagesDiff       uint8   `arff:"amountDisappearingImagesDiff" csv:"amountDisappearingImagesDiff"`
	AmountDisappearingImagesSum        uint8   `arff:"amountDisappearingImagesSum" csv:"amountDisappearingImagesSum"`
	AmountDisappearingImagesMean       float64 `arff:"amountDisappearingImagesMean" csv:"amountDisappearingImagesMean"`
	AmountDisappearingImagesStdDev     float64 `arff:"amountDisappearingImagesStdDev" csv:"amountDisappearingImagesStdDev"`
	AmountDisappearingImagesSkew       float64 `arff:"amountDisappearingImagesSkew" csv:"amountDisappearingImagesSkew"`
	ChosenTaskImagesXMin               uint8   `arff:"chosenTaskImagesXMin" csv:"chosenTaskImagesXMin"`
	ChosenTaskImagesXMax               uint8   `arff:"chosenTaskImagesXMax" csv:"chosenTaskImagesXMax"`
	ChosenTaskImagesXDiff              uint8   `arff:"chosenTaskImagesXDiff" csv:"chosenTaskImagesXDiff"`
	ChosenTaskImagesXSum               uint8   `arff:"chosenTaskImagesXSum" csv:"chosenTaskImagesXSum"`
	ChosenTaskImagesXMean              float64 `arff:"chosenTaskImagesXMean" csv:"chosenTaskImagesXMean"`
	ChosenTaskImagesXStdDev            float64 `arff:"chosenTaskImagesXStdDev" csv:"chosenTaskImagesXStdDev"`
	ChosenTaskImagesXSkew              float64 `arff:"chosenTaskImagesXSkew" csv:"chosenTaskImagesImagesXSkew"`
	ChosenTaskImagesYMin               uint8   `arff:"chosenTaskImagesYMin" csv:"chosenTaskImagesYMin"`
	ChosenTaskImagesYMax               uint8   `arff:"chosenTaskImagesYMax" csv:"chosenTaskImagesYMax"`
	ChosenTaskImagesYDiff              uint8   `arff:"chosenTaskImagesYDiff" csv:"chosenTaskImagesYDiff"`
	ChosenTaskImagesYSum               uint8   `arff:"chosenTaskImagesYSum" csv:"chosenTaskImagesYSum"`
	ChosenTaskImagesYMean              float64 `arff:"chosenTaskImagesYMean" csv:"chosenTaskImagesYMean"`
	ChosenTaskImagesYStdDev            float64 `arff:"chosenTaskImagesYStdDev" csv:"chosenTaskImagesYStdDev"`
	ChosenTaskImagesYSkew              float64 `arff:"chosenTaskImagesYSkew" csv:"chosenTaskImagesImagesYSkew"`
	ChosenTaskImagesYDistMin           int16   `arff:"chosenTaskImagesYDistMin" csv:"chosenTaskImagesYDistMin"`
	ChosenTaskImagesYDistMax           int16   `arff:"chosenTaskImagesYDistMax" csv:"chosenTaskImagesYDistMax"`
	ChosenTaskImagesYDistDiff          int16   `arff:"chosenTaskImagesYDistDiff" csv:"chosenTaskImagesYDistDiff"`
	ChosenTaskImagesYDistSum           int16   `arff:"chosenTaskImagesYDistSum" csv:"chosenTaskImagesYDistSum"`
	ChosenTaskImagesYDistMean          float64 `arff:"chosenTaskImagesYDistMean" csv:"chosenTaskImagesYDistMean"`
	ChosenTaskImagesYDistStdDev        float64 `arff:"chosenTaskImagesYDistStdDev" csv:"chosenTaskImagesYDistStdDev"`
	ChosenTaskImagesYDistSkew          float64 `arff:"chosenTaskImagesYDistSkew" csv:"chosenTaskImagesImagesYDistSkew"`
	ChosenTaskImagesXDistMin           int16   `arff:"chosenTaskImagesXDistMin" csv:"chosenTaskImagesXDistMin"`
	ChosenTaskImagesXDistMax           int16   `arff:"chosenTaskImagesXDistMax" csv:"chosenTaskImagesXDistMax"`
	ChosenTaskImagesXDistDiff          int16   `arff:"chosenTaskImagesXDistDiff" csv:"chosenTaskImagesXDistDiff"`
	ChosenTaskImagesXDistSum           int16   `arff:"chosenTaskImagesXDistSum" csv:"chosenTaskImagesXDistSum"`
	ChosenTaskImagesXDistMean          float64 `arff:"chosenTaskImagesXDistMean" csv:"chosenTaskImagesXDistMean"`
	ChosenTaskImagesXDistStdDev        float64 `arff:"chosenTaskImagesXDistStdDev" csv:"chosenTaskImagesXDistStdDev"`
	ChosenTaskImagesXDistSkew          float64 `arff:"chosenTaskImagesXDistSkew" csv:"chosenTaskImagesImagesXDistSkew"`
	ChosenTaskImagesPairwiseDistMin    float64 `arff:"chosenTaskImagesPairwiseDistMin" csv:"chosenTaskImagesPairwiseDistMin"`
	ChosenTaskImagesPairwiseDistMax    float64 `arff:"chosenTaskImagesPairwiseDistMax" csv:"chosenTaskImagesPairwiseDistMax"`
	ChosenTaskImagesPairwiseDistDiff   float64 `arff:"chosenTaskImagesPairwiseDistDiff" csv:"chosenTaskImagesPairwiseDistDiff"`
	ChosenTaskImagesPairwiseDistSum    float64 `arff:"chosenTaskImagesPairwiseDistSum" csv:"chosenTaskImagesPairwiseDistSum"`
	ChosenTaskImagesPairwiseDistMean   float64 `arff:"chosenTaskImagesPairwiseDistMean" csv:"chosenTaskImagesPairwiseDistMean"`
	ChosenTaskImagesPairwiseDistStdDev float64 `arff:"chosenTaskImagesPairwiseDistStdDev" csv:"chosenTaskImagesPairwiseDistStdDev"`
	ChosenTaskImagesPairwiseDistSkew   float64 `arff:"chosenTaskImagesPairwiseDistSkew" csv:"chosenTaskImagesImagesDistStartEndSkew"`
	ChosenTaskImagesDistStartEndMin    float64 `arff:"chosenTaskImagesDistStartEndMin" csv:"chosenTaskImagesDistStartEndMin"`
	ChosenTaskImagesDistStartEndMax    float64 `arff:"chosenTaskImagesDistStartEndMax" csv:"chosenTaskImagesDistStartEndMax"`
	ChosenTaskImagesDistStartEndDiff   float64 `arff:"chosenTaskImagesDistStartEndDiff" csv:"chosenTaskImagesDistStartEndDiff"`
	ChosenTaskImagesDistStartEndSum    float64 `arff:"chosenTaskImagesDistStartEndSum" csv:"chosenTaskImagesDistStartEndSum"`
	ChosenTaskImagesDistStartEndMean   float64 `arff:"chosenTaskImagesDistStartEndMean" csv:"chosenTaskImagesDistStartEndMean"`
	ChosenTaskImagesDistStartEndStdDev float64 `arff:"chosenTaskImagesDistStartEndStdDev" csv:"chosenTaskImagesDistStartEndStdDev"`
	ChosenTaskImagesDistStartEndSkew   float64 `arff:"chosenTaskImagesDistStartEndSkew" csv:"chosenTaskImagesImagesDistStartEndSkew"`
	ChosenTaskImagesDistSumMin         float64 `arff:"chosenTaskImagesDistSumMin" csv:"chosenTaskImagesDistSumMin"`
	ChosenTaskImagesDistSumMax         float64 `arff:"chosenTaskImagesDistSumMax" csv:"chosenTaskImagesDistSumMax"`
	ChosenTaskImagesDistSumDiff        float64 `arff:"chosenTaskImagesDistSumDiff" csv:"chosenTaskImagesDistSumDiff"`
	ChosenTaskImagesDistSumSum         float64 `arff:"chosenTaskImagesDistSumSum" csv:"chosenTaskImagesDistSumSum"`
	ChosenTaskImagesDistSumMean        float64 `arff:"chosenTaskImagesDistSumMean" csv:"chosenTaskImagesDistSumMean"`
	ChosenTaskImagesDistSumStdDev      float64 `arff:"chosenTaskImagesDistSumStdDev" csv:"chosenTaskImagesDistSumStdDev"`
	ChosenTaskImagesDistSumSkew        float64 `arff:"chosenTaskImagesDistSumSkew" csv:"chosenTaskImagesImagesDistSumSkew"`
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
	// Rest Paths
	RestPathPairwiseDistanceMin                    float64 `arff:"restpathPairwiseDistanceMin" csv:"restpathPairwiseDistanceMin"`
	RestPathPairwiseDistanceMax                    float64 `arff:"restpathPairwiseDistanceMax" csv:"restpathPairwiseDistanceMax"`
	RestPathPairwiseDistanceDiff                   float64 `arff:"restpathPairwiseDistanceDiff" csv:"restpathPairwiseDistanceDiff"`
	RestPathPairwiseDistanceSum                    float64 `arff:"restpathPairwiseDistanceSum" csv:"restpathPairwiseDistanceSum"`
	RestPathPairwiseDistanceMean                   float64 `arff:"restpathPairwiseDistanceMean" csv:"restpathPairwiseDistanceMean"`
	RestPathPairwiseDistanceStdDev                 float64 `arff:"restpathPairwiseDistanceStdDev" csv:"restpathPairwiseDistanceStdDev"`
	RestPathPairwiseDistanceSkew                   float64 `arff:"restpathPairwiseDistanceSkew" csv:"restpathPairwiseDistanceSkew"`
	RestPathPairwiseVelocityMin                    float64 `arff:"restpathPairwiseVelocityMin" csv:"restpathPairwiseVelocityMin"`
	RestPathPairwiseVelocityMax                    float64 `arff:"restpathPairwiseVelocityMax" csv:"restpathPairwiseVelocityMax"`
	RestPathPairwiseVelocityDiff                   float64 `arff:"restpathPairwiseVelocityDiff" csv:"restpathPairwiseVelocityDiff"`
	RestPathPairwiseVelocitySum                    float64 `arff:"restpathPairwiseVelocitySum" csv:"restpathPairwiseVelocitySum"`
	RestPathPairwiseVelocityMean                   float64 `arff:"restpathPairwiseVelocityMean" csv:"restpathPairwiseVelocityMean"`
	RestPathPairwiseVelocityStdDev                 float64 `arff:"restpathPairwiseVelocityStdDev" csv:"restpathPairwiseVelocityStdDev"`
	RestPathPairwiseVelocitySkew                   float64 `arff:"restpathPairwiseVelocitySkew" csv:"restpathPairwiseVelocitySkew"`
	RestPathPairwiseAccelerationMin                float64 `arff:"restpathPairwiseAccelerationMin" csv:"restpathPairwiseAccelerationMin"`
	RestPathPairwiseAccelerationMax                float64 `arff:"restpathPairwiseAccelerationMax" csv:"restpathPairwiseAccelerationMax"`
	RestPathPairwiseAccelerationDiff               float64 `arff:"restpathPairwiseAccelerationDiff" csv:"restpathPairwiseAccelerationDiff"`
	RestPathPairwiseAccelerationSum                float64 `arff:"restpathPairwiseAccelerationSum" csv:"restpathPairwiseAccelerationSum"`
	RestPathPairwiseAccelerationMean               float64 `arff:"restpathPairwiseAccelerationMean" csv:"restpathPairwiseAccelerationMean"`
	RestPathPairwiseAccelerationStdDev             float64 `arff:"restpathPairwiseAccelerationStdDev" csv:"restpathPairwiseAccelerationStdDev"`
	RestPathPairwiseAccelerationSkew               float64 `arff:"restpathPairwiseAccelerationSkew" csv:"restpathPairwiseAccelerationSkew"`
	RestPathPairwiseAngleMin                       float64 `arff:"restpathPairwiseAngleMin" csv:"restpathPairwiseAngleMin"`
	RestPathPairwiseAngleMax                       float64 `arff:"restpathPairwiseAngleMax" csv:"restpathPairwiseAngleMax"`
	RestPathPairwiseAngleDiff                      float64 `arff:"restpathPairwiseAngleDiff" csv:"restpathPairwiseAngleDiff"`
	RestPathPairwiseAngleSum                       float64 `arff:"restpathPairwiseAngleSum" csv:"restpathPairwiseAngleSum"`
	RestPathPairwiseAngleMean                      float64 `arff:"restpathPairwiseAngleMean" csv:"restpathPairwiseAngleMean"`
	RestPathPairwiseAngleStdDev                    float64 `arff:"restpathPairwiseAngleStdDev" csv:"restpathPairwiseAngleStdDev"`
	RestPathPairwiseAngleSkew                      float64 `arff:"restpathPairwiseAngleSkew" csv:"restpathPairwiseAngleSkew"`
	RestPathAngleBetweenMovementAndStartEndMin     float64 `arff:"restpathAngleBetweenMovementAndStartEndMin" csv:"restpathAngleBetweenMovementAndStartEndMin"`
	RestPathAngleBetweenMovementAndStartEndMax     float64 `arff:"restpathAngleBetweenMovementAndStartEndMax" csv:"restpathAngleBetweenMovementAndStartEndMax"`
	RestPathAngleBetweenMovementAndStartEndDiff    float64 `arff:"restpathAngleBetweenMovementAndStartEndDiff" csv:"restpathAngleBetweenMovementAndStartEndDiff"`
	RestPathAngleBetweenMovementAndStartEndSum     float64 `arff:"restpathAngleBetweenMovementAndStartEndSum" csv:"restpathAngleBetweenMovementAndStartEndSum"`
	RestPathAngleBetweenMovementAndStartEndMean    float64 `arff:"restpathAngleBetweenMovementAndStartEndMean" csv:"restpathAngleBetweenMovementAndStartEndMean"`
	RestPathAngleBetweenMovementAndStartEndStdDev  float64 `arff:"restpathAngleBetweenMovementAndStartEndStdDev" csv:"restpathAngleBetweenMovementAndStartEndStdDev"`
	RestPathAngleBetweenMovementAndStartEndSkew    float64 `arff:"restpathAngleBetweenMovementAndStartEndSkew" csv:"restpathAngleBetweenMovementAndStartEndSkew"`
	RestPathPairwiseAngularVelocityMin             float64 `arff:"restpathPairwiseAngularVelocityMin" csv:"restpathPairwiseAngularVelocityMin"`
	RestPathPairwiseAngularVelocityMax             float64 `arff:"restpathPairwiseAngularVelocityMax" csv:"restpathPairwiseAngularVelocityMax"`
	RestPathPairwiseAngularVelocityDiff            float64 `arff:"restpathPairwiseAngularVelocityDiff" csv:"restpathPairwiseAngularVelocityDiff"`
	RestPathPairwiseAngularVelocitySum             float64 `arff:"restpathPairwiseAngularVelocitySum" csv:"restpathPairwiseAngularVelocitySum"`
	RestPathPairwiseAngularVelocityMean            float64 `arff:"restpathPairwiseAngularVelocityMean" csv:"restpathPairwiseAngularVelocityMean"`
	RestPathPairwiseAngularVelocityStdDev          float64 `arff:"restpathPairwiseAngularVelocityStdDev" csv:"restpathPairwiseAngularVelocityStdDev"`
	RestPathPairwiseAngularVelocitySkew            float64 `arff:"restpathPairwiseAngularVelocitySkew" csv:"restpathPairwiseAngularVelocitySkew"`
	RestPathPairwiseDurationMin                    uint64  `arff:"restpathPairwiseDurationMin" csv:"restpathPairwiseDurationMin"`
	RestPathPairwiseDurationMax                    uint64  `arff:"restpathPairwiseDurationMax" csv:"restpathPairwiseDurationMax"`
	RestPathPairwiseDurationDiff                   uint64  `arff:"restpathPairwiseDurationDiff" csv:"restpathPairwiseDurationDiff"`
	RestPathPairwiseDurationSum                    uint64  `arff:"restpathPairwiseDurationSum" csv:"restpathPairwiseDurationSum"`
	RestPathPairwiseDurationMean                   float64 `arff:"restpathPairwiseDurationMean" csv:"restpathPairwiseDurationMean"`
	RestPathPairwiseDurationStdDev                 float64 `arff:"restpathPairwiseDurationStdDev" csv:"restpathPairwiseDurationStdDev"`
	RestPathPairwiseDurationSkew                   float64 `arff:"restpathPairwiseDurationSkew" csv:"restpathPairwiseDurationSkew"`
	RestPathTimeBetweenClickAndReleaseMin          uint64  `arff:"restpathTimeBetweenClickAndReleaseMin" csv:"restpathTimeBetweenClickAndReleaseMin"`
	RestPathTimeBetweenClickAndReleaseMax          uint64  `arff:"restpathTimeBetweenClickAndReleaseMax" csv:"restpathTimeBetweenClickAndReleaseMax"`
	RestPathTimeBetweenClickAndReleaseDiff         uint64  `arff:"restpathTimeBetweenClickAndReleaseDiff" csv:"restpathTimeBetweenClickAndReleaseDiff"`
	RestPathTimeBetweenClickAndReleaseSum          uint64  `arff:"restpathTimeBetweenClickAndReleaseSum" csv:"restpathTimeBetweenClickAndReleaseSum"`
	RestPathTimeBetweenClickAndReleaseMean         float64 `arff:"restpathTimeBetweenClickAndReleaseMean" csv:"restpathTimeBetweenClickAndReleaseMean"`
	RestPathTimeBetweenClickAndReleaseStdDev       float64 `arff:"restpathTimeBetweenClickAndReleaseStdDev" csv:"restpathTimeBetweenClickAndReleaseStdDev"`
	RestPathTimeBetweenClickAndReleaseSkew         float64 `arff:"restpathTimeBetweenClickAndReleaseSkew" csv:"restpathTimeBetweenClickAndReleaseSkew"`
	RestPathBreakTimesMin                          uint64  `arff:"restpathBreakTimesMin" csv:"restpathBreakTimesMin"`
	RestPathBreakTimesMax                          uint64  `arff:"restpathBreakTimesMax" csv:"restpathBreakTimesMax"`
	RestPathBreakTimesDiff                         uint64  `arff:"restpathBreakTimesDiff" csv:"restpathBreakTimesDiff"`
	RestPathBreakTimesSum                          uint64  `arff:"restpathBreakTimesSum" csv:"restpathBreakTimesSum"`
	RestPathBreakTimesMean                         float64 `arff:"restpathBreakTimesMean" csv:"restpathBreakTimesMean"`
	RestPathBreakTimesStdDev                       float64 `arff:"restpathBreakTimesStdDev" csv:"restpathBreakTimesStdDev"`
	RestPathBreakTimesSkew                         float64 `arff:"restpathBreakTimesSkew" csv:"restpathBreakTimesSkew"`
	RestPathMovementDuringClickDistanceMin         float64 `arff:"restpathMovementDuringClickDistanceMin" csv:"restpathMovementDuringClickDistanceMin"`
	RestPathMovementDuringClickDistanceMax         float64 `arff:"restpathMovementDuringClickDistanceMax" csv:"restpathMovementDuringClickDistanceMax"`
	RestPathMovementDuringClickDistanceDiff        float64 `arff:"restpathMovementDuringClickDistanceDiff" csv:"restpathMovementDuringClickDistanceDiff"`
	RestPathMovementDuringClickDistanceSum         float64 `arff:"restpathMovementDuringClickDistanceSum" csv:"restpathMovementDuringClickDistanceSum"`
	RestPathMovementDuringClickDistanceMean        float64 `arff:"restpathMovementDuringClickDistanceMean" csv:"restpathMovementDuringClickDistanceMean"`
	RestPathMovementDuringClickDistanceStdDev      float64 `arff:"restpathMovementDuringClickDistanceStdDev" csv:"restpathMovementDuringClickDistanceStdDev"`
	RestPathMovementDuringClickDistanceSkew        float64 `arff:"restpathMovementDuringClickDistanceSkew" csv:"restpathMovementDuringClickDistanceSkew"`
	RestPathMovementDuringClickVelocityMin         float64 `arff:"restpathMovementDuringClickVelocityMin" csv:"restpathMovementDuringClickVelocityMin"`
	RestPathMovementDuringClickVelocityMax         float64 `arff:"restpathMovementDuringClickVelocityMax" csv:"restpathMovementDuringClickVelocityMax"`
	RestPathMovementDuringClickVelocityDiff        float64 `arff:"restpathMovementDuringClickVelocityDiff" csv:"restpathMovementDuringClickVelocityDiff"`
	RestPathMovementDuringClickVelocitySum         float64 `arff:"restpathMovementDuringClickVelocitySum" csv:"restpathMovementDuringClickVelocitySum"`
	RestPathMovementDuringClickVelocityMean        float64 `arff:"restpathMovementDuringClickVelocityMean" csv:"restpathMovementDuringClickVelocityMean"`
	RestPathMovementDuringClickVelocityStdDev      float64 `arff:"restpathMovementDuringClickVelocityStdDev" csv:"restpathMovementDuringClickVelocityStdDev"`
	RestPathMovementDuringClickVelocitySkew        float64 `arff:"restpathMovementDuringClickVelocitySkew" csv:"restpathMovementDuringClickVelocitySkew"`
	RestPathMovementDuringClickAccelerationMin     float64 `arff:"restpathMovementDuringClickAccelerationMin" csv:"restpathMovementDuringClickAccelerationMin"`
	RestPathMovementDuringClickAccelerationMax     float64 `arff:"restpathMovementDuringClickAccelerationMax" csv:"restpathMovementDuringClickAccelerationMax"`
	RestPathMovementDuringClickAccelerationDiff    float64 `arff:"restpathMovementDuringClickAccelerationDiff" csv:"restpathMovementDuringClickAccelerationDiff"`
	RestPathMovementDuringClickAccelerationSum     float64 `arff:"restpathMovementDuringClickAccelerationSum" csv:"restpathMovementDuringClickAccelerationSum"`
	RestPathMovementDuringClickAccelerationMean    float64 `arff:"restpathMovementDuringClickAccelerationMean" csv:"restpathMovementDuringClickAccelerationMean"`
	RestPathMovementDuringClickAccelerationStdDev  float64 `arff:"restpathMovementDuringClickAccelerationStdDev" csv:"restpathMovementDuringClickAccelerationStdDev"`
	RestPathMovementDuringClickAccelerationSkew    float64 `arff:"restpathMovementDuringClickAccelerationSkew" csv:"restpathMovementDuringClickAccelerationSkew"`
	RestPathMovementDuringClickAngleMin            float64 `arff:"restpathMovementDuringClickAngleMin" csv:"restpathMovementDuringClickAngleMin"`
	RestPathMovementDuringClickAngleMax            float64 `arff:"restpathMovementDuringClickAngleMax" csv:"restpathMovementDuringClickAngleMax"`
	RestPathMovementDuringClickAngleDiff           float64 `arff:"restpathMovementDuringClickAngleDiff" csv:"restpathMovementDuringClickAngleDiff"`
	RestPathMovementDuringClickAngleSum            float64 `arff:"restpathMovementDuringClickAngleSum" csv:"restpathMovementDuringClickAngleSum"`
	RestPathMovementDuringClickAngleMean           float64 `arff:"restpathMovementDuringClickAngleMean" csv:"restpathMovementDuringClickAngleMean"`
	RestPathMovementDuringClickAngleStdDev         float64 `arff:"restpathMovementDuringClickAngleStdDev" csv:"restpathMovementDuringClickAngleStdDev"`
	RestPathMovementDuringClickAngleSkew           float64 `arff:"restpathMovementDuringClickAngleSkew" csv:"restpathMovementDuringClickAngleSkew"`
	RestPathScrollDXMin                            float64 `arff:"restpathScrollDXMin" csv:"restpathScrollDXMin"`
	RestPathScrollDXMax                            float64 `arff:"restpathScrollDXMax" csv:"restpathScrollDXMax"`
	RestPathScrollDXDiff                           float64 `arff:"restpathScrollDXDiff" csv:"restpathScrollDXDiff"`
	RestPathScrollDXSum                            float64 `arff:"restpathScrollDXSum" csv:"restpathScrollDXSum"`
	RestPathScrollDXMean                           float64 `arff:"restpathScrollDXMean" csv:"restpathScrollDXMean"`
	RestPathScrollDXStdDev                         float64 `arff:"restpathScrollDXStdDev" csv:"restpathScrollDXStdDev"`
	RestPathScrollDXSkew                           float64 `arff:"restpathScrollDXSkew" csv:"restpathScrollDXSkew"`
	RestPathScrollDYMin                            float64 `arff:"restpathScrollDYMin" csv:"restpathScrollDYMin"`
	RestPathScrollDYMax                            float64 `arff:"restpathScrollDYMax" csv:"restpathScrollDYMax"`
	RestPathScrollDYDiff                           float64 `arff:"restpathScrollDYDiff" csv:"restpathScrollDYDiff"`
	RestPathScrollDYSum                            float64 `arff:"restpathScrollDYSum" csv:"restpathScrollDYSum"`
	RestPathScrollDYMean                           float64 `arff:"restpathScrollDYMean" csv:"restpathScrollDYMean"`
	RestPathScrollDYStdDev                         float64 `arff:"restpathScrollDYStdDev" csv:"restpathScrollDYStdDev"`
	RestPathScrollDYSkew                           float64 `arff:"restpathScrollDYSkew" csv:"restpathScrollDYSkew"`
	RestPathScrollDZMin                            float64 `arff:"restpathScrollDZMin" csv:"restpathScrollDZMin"`
	RestPathScrollDZMax                            float64 `arff:"restpathScrollDZMax" csv:"restpathScrollDZMax"`
	RestPathScrollDZDiff                           float64 `arff:"restpathScrollDZDiff" csv:"restpathScrollDZDiff"`
	RestPathScrollDZSum                            float64 `arff:"restpathScrollDZSum" csv:"restpathScrollDZSum"`
	RestPathScrollDZMean                           float64 `arff:"restpathScrollDZMean" csv:"restpathScrollDZMean"`
	RestPathScrollDZStdDev                         float64 `arff:"restpathScrollDZStdDev" csv:"restpathScrollDZStdDev"`
	RestPathScrollDZSkew                           float64 `arff:"restpathScrollDZSkew" csv:"restpathScrollDZSkew"`
	RestPathScrollDMMin                            uint8   `arff:"restpathScrollDMMin" csv:"restpathScrollDMMin"`
	RestPathScrollDMMax                            uint8   `arff:"restpathScrollDMMax" csv:"restpathScrollDMMax"`
	RestPathScrollDMDiff                           uint8   `arff:"restpathScrollDMDiff" csv:"restpathScrollDMDiff"`
	RestPathScrollDMSum                            uint8   `arff:"restpathScrollDMSum" csv:"restpathScrollDMSum"`
	RestPathScrollDMMean                           float64 `arff:"restpathScrollDMMean" csv:"restpathScrollDMMean"`
	RestPathScrollDMStdDev                         float64 `arff:"restpathScrollDMStdDev" csv:"restpathScrollDMStdDev"`
	RestPathScrollDMSkew                           float64 `arff:"restpathScrollDMSkew" csv:"restpathScrollDMSkew"`
	RestPathXPointsMin                             float64 `arff:"restpathXPointsMin" csv:"restpathXPointsMin"`
	RestPathXPointsMax                             float64 `arff:"restpathXPointsMax" csv:"restpathXPointsMax"`
	RestPathXPointsDiff                            float64 `arff:"restpathXPointsDiff" csv:"restpathXPointsDiff"`
	RestPathXPointsSum                             float64 `arff:"restpathXPointsSum" csv:"restpathXPointsSum"`
	RestPathXPointsMean                            float64 `arff:"restpathXPointsMean" csv:"restpathXPointsMean"`
	RestPathXPointsStdDev                          float64 `arff:"restpathXPointsStdDev" csv:"restpathXPointsStdDev"`
	RestPathXPointsSkew                            float64 `arff:"restpathXPointsSkew" csv:"restpathXPointsSkew"`
	RestPathYPointsMin                             float64 `arff:"restpathYPointsMin" csv:"restpathYPointsMin"`
	RestPathYPointsMax                             float64 `arff:"restpathYPointsMax" csv:"restpathYPointsMax"`
	RestPathYPointsDiff                            float64 `arff:"restpathYPointsDiff" csv:"restpathYPointsDiff"`
	RestPathYPointsSum                             float64 `arff:"restpathYPointsSum" csv:"restpathYPointsSum"`
	RestPathYPointsMean                            float64 `arff:"restpathYPointsMean" csv:"restpathYPointsMean"`
	RestPathYPointsStdDev                          float64 `arff:"restpathYPointsStdDev" csv:"restpathYPointsStdDev"`
	RestPathYPointsSkew                            float64 `arff:"restpathYPointsSkew" csv:"restpathYPointsSkew"`
	RestPathPairwiseXVelocityMin                   float64 `arff:"restpathPairwiseXVelocityMin" csv:"restpathPairwiseXVelocityMin"`
	RestPathPairwiseXVelocityMax                   float64 `arff:"restpathPairwiseXVelocityMax" csv:"restpathPairwiseXVelocityMax"`
	RestPathPairwiseXVelocityDiff                  float64 `arff:"restpathPairwiseXVelocityDiff" csv:"restpathPairwiseXVelocityDiff"`
	RestPathPairwiseXVelocitySum                   float64 `arff:"restpathPairwiseXVelocitySum" csv:"restpathPairwiseXVelocitySum"`
	RestPathPairwiseXVelocityMean                  float64 `arff:"restpathPairwiseXVelocityMean" csv:"restpathPairwiseXVelocityMean"`
	RestPathPairwiseXVelocityStdDev                float64 `arff:"restpathPairwiseXVelocityStdDev" csv:"restpathPairwiseXVelocityStdDev"`
	RestPathPairwiseXVelocitySkew                  float64 `arff:"restpathPairwiseXVelocitySkew" csv:"restpathPairwiseXVelocitySkew"`
	RestPathPairwiseYVelocityMin                   float64 `arff:"restpathPairwiseYVelocityMin" csv:"restpathPairwiseYVelocityMin"`
	RestPathPairwiseYVelocityMax                   float64 `arff:"restpathPairwiseYVelocityMax" csv:"restpathPairwiseYVelocityMax"`
	RestPathPairwiseYVelocityDiff                  float64 `arff:"restpathPairwiseYVelocityDiff" csv:"restpathPairwiseYVelocityDiff"`
	RestPathPairwiseYVelocitySum                   float64 `arff:"restpathPairwiseYVelocitySum" csv:"restpathPairwiseYVelocitySum"`
	RestPathPairwiseYVelocityMean                  float64 `arff:"restpathPairwiseYVelocityMean" csv:"restpathPairwiseYVelocityMean"`
	RestPathPairwiseYVelocityStdDev                float64 `arff:"restpathPairwiseYVelocityStdDev" csv:"restpathPairwiseYVelocityStdDev"`
	RestPathPairwiseYVelocitySkew                  float64 `arff:"restpathPairwiseYVelocitySkew" csv:"restpathPairwiseYVelocitySkew"`
	RestPathPairwiseXDistanceMin                   float64 `arff:"restpathPairwiseXDistanceMin" csv:"restpathPairwiseXDistanceMin"`
	RestPathPairwiseXDistanceMax                   float64 `arff:"restpathPairwiseXDistanceMax" csv:"restpathPairwiseXDistanceMax"`
	RestPathPairwiseXDistanceDiff                  float64 `arff:"restpathPairwiseXDistanceDiff" csv:"restpathPairwiseXDistanceDiff"`
	RestPathPairwiseXDistanceSum                   float64 `arff:"restpathPairwiseXDistanceSum" csv:"restpathPairwiseXDistanceSum"`
	RestPathPairwiseXDistanceMean                  float64 `arff:"restpathPairwiseXDistanceMean" csv:"restpathPairwiseXDistanceMean"`
	RestPathPairwiseXDistanceStdDev                float64 `arff:"restpathPairwiseXDistanceStdDev" csv:"restpathPairwiseXDistanceStdDev"`
	RestPathPairwiseXDistanceSkew                  float64 `arff:"restpathPairwiseXDistanceSkew" csv:"restpathPairwiseXDistanceSkew"`
	RestPathPairwiseYDistanceMin                   float64 `arff:"restpathPairwiseYDistanceMin" csv:"restpathPairwiseYDistanceMin"`
	RestPathPairwiseYDistanceMax                   float64 `arff:"restpathPairwiseYDistanceMax" csv:"restpathPairwiseYDistanceMax"`
	RestPathPairwiseYDistanceDiff                  float64 `arff:"restpathPairwiseYDistanceDiff" csv:"restpathPairwiseYDistanceDiff"`
	RestPathPairwiseYDistanceSum                   float64 `arff:"restpathPairwiseYDistanceSum" csv:"restpathPairwiseYDistanceSum"`
	RestPathPairwiseYDistanceMean                  float64 `arff:"restpathPairwiseYDistanceMean" csv:"restpathPairwiseYDistanceMean"`
	RestPathPairwiseYDistanceStdDev                float64 `arff:"restpathPairwiseYDistanceStdDev" csv:"restpathPairwiseYDistanceStdDev"`
	RestPathPairwiseYDistanceSkew                  float64 `arff:"restpathPairwiseYDistanceSkew" csv:"restpathPairwiseYDistanceSkew"`
	RestPathSumYDistanceMin                        float64 `arff:"restpathSumYDistanceMin" csv:"restpathSumYDistanceMin"`
	RestPathSumYDistanceMax                        float64 `arff:"restpathSumYDistanceMax" csv:"restpathSumYDistanceMax"`
	RestPathSumYDistanceDiff                       float64 `arff:"restpathSumYDistanceDiff" csv:"restpathSumYDistanceDiff"`
	RestPathSumYDistanceSum                        float64 `arff:"restpathSumYDistanceSum" csv:"restpathSumYDistanceSum"`
	RestPathSumYDistanceMean                       float64 `arff:"restpathSumYDistanceMean" csv:"restpathSumYDistanceMean"`
	RestPathSumYDistanceStdDev                     float64 `arff:"restpathSumYDistanceStdDev" csv:"restpathSumYDistanceStdDev"`
	RestPathSumYDistanceSkew                       float64 `arff:"restpathSumYDistanceSkew" csv:"restpathSumYDistanceSkew"`
	RestPathMeanYVelocityMin                       float64 `arff:"restpathMeanYVelocityMin" csv:"restpathMeanYVelocityMin"`
	RestPathMeanYVelocityMax                       float64 `arff:"restpathMeanYVelocityMax" csv:"restpathMeanYVelocityMax"`
	RestPathMeanYVelocityDiff                      float64 `arff:"restpathMeanYVelocityDiff" csv:"restpathMeanYVelocityDiff"`
	RestPathMeanYVelocitySum                       float64 `arff:"restpathMeanYVelocitySum" csv:"restpathMeanYVelocitySum"`
	RestPathMeanYVelocityMean                      float64 `arff:"restpathMeanYVelocityMean" csv:"restpathMeanYVelocityMean"`
	RestPathMeanYVelocityStdDev                    float64 `arff:"restpathMeanYVelocityStdDev" csv:"restpathMeanYVelocityStdDev"`
	RestPathMeanYVelocitySkew                      float64 `arff:"restpathMeanYVelocitySkew" csv:"restpathMeanYVelocitySkew"`
	RestPathSumXDistanceMin                        float64 `arff:"restpathSumXDistanceMin" csv:"restpathSumXDistanceMin"`
	RestPathSumXDistanceMax                        float64 `arff:"restpathSumXDistanceMax" csv:"restpathSumXDistanceMax"`
	RestPathSumXDistanceDiff                       float64 `arff:"restpathSumXDistanceDiff" csv:"restpathSumXDistanceDiff"`
	RestPathSumXDistanceSum                        float64 `arff:"restpathSumXDistanceSum" csv:"restpathSumXDistanceSum"`
	RestPathSumXDistanceMean                       float64 `arff:"restpathSumXDistanceMean" csv:"restpathSumXDistanceMean"`
	RestPathSumXDistanceStdDev                     float64 `arff:"restpathSumXDistanceStdDev" csv:"restpathSumXDistanceStdDev"`
	RestPathSumXDistanceSkew                       float64 `arff:"restpathSumXDistanceSkew" csv:"restpathSumXDistanceSkew"`
	RestPathMeanXVelocityMin                       float64 `arff:"restpathMeanXVelocityMin" csv:"restpathMeanXVelocityMin"`
	RestPathMeanXVelocityMax                       float64 `arff:"restpathMeanXVelocityMax" csv:"restpathMeanXVelocityMax"`
	RestPathMeanXVelocityDiff                      float64 `arff:"restpathMeanXVelocityDiff" csv:"restpathMeanXVelocityDiff"`
	RestPathMeanXVelocitySum                       float64 `arff:"restpathMeanXVelocitySum" csv:"restpathMeanXVelocitySum"`
	RestPathMeanXVelocityMean                      float64 `arff:"restpathMeanXVelocityMean" csv:"restpathMeanXVelocityMean"`
	RestPathMeanXVelocityStdDev                    float64 `arff:"restpathMeanXVelocityStdDev" csv:"restpathMeanXVelocityStdDev"`
	RestPathMeanXVelocitySkew                      float64 `arff:"restpathMeanXVelocitySkew" csv:"restpathMeanXVelocitySkew"`
	RestPathStraightnessMin                        float64 `arff:"restpathStraightnessMin" csv:"restpathStraightnessMin"`
	RestPathStraightnessMax                        float64 `arff:"restpathStraightnessMax" csv:"restpathStraightnessMax"`
	RestPathStraightnessDiff                       float64 `arff:"restpathStraightnessDiff" csv:"restpathStraightnessDiff"`
	RestPathStraightnessSum                        float64 `arff:"restpathStraightnessSum" csv:"restpathStraightnessSum"`
	RestPathStraightnessMean                       float64 `arff:"restpathStraightnessMean" csv:"restpathStraightnessMean"`
	RestPathStraightnessStdDev                     float64 `arff:"restpathStraightnessStdDev" csv:"restpathStraightnessStdDev"`
	RestPathStraightnessSkew                       float64 `arff:"restpathStraightnessSkew" csv:"restpathStraightnessSkew"`
	RestPathNumberOfRightClicksMin                 uint8   `arff:"restpathNumberOfRightClicksMin" csv:"restpathNumberOfRightClicksMin"`
	RestPathNumberOfRightClicksMax                 uint8   `arff:"restpathNumberOfRightClicksMax" csv:"restpathNumberOfRightClicksMax"`
	RestPathNumberOfRightClicksDiff                uint8   `arff:"restpathNumberOfRightClicksDiff" csv:"restpathNumberOfRightClicksDiff"`
	RestPathNumberOfRightClicksSum                 uint8   `arff:"restpathNumberOfRightClicksSum" csv:"restpathNumberOfRightClicksSum"`
	RestPathNumberOfRightClicksMean                float64 `arff:"restpathNumberOfRightClicksMean" csv:"restpathNumberOfRightClicksMean"`
	RestPathNumberOfRightClicksStdDev              float64 `arff:"restpathNumberOfRightClicksStdDev" csv:"restpathNumberOfRightClicksStdDev"`
	RestPathNumberOfRightClicksSkew                float64 `arff:"restpathNumberOfRightClicksSkew" csv:"restpathNumberOfRightClicksSkew"`
	RestPathNumberOfMiddleClicksMin                uint8   `arff:"restpathNumberOfMiddleClicksMin" csv:"restpathNumberOfMiddleClicksMin"`
	RestPathNumberOfMiddleClicksMax                uint8   `arff:"restpathNumberOfMiddleClicksMax" csv:"restpathNumberOfMiddleClicksMax"`
	RestPathNumberOfMiddleClicksDiff               uint8   `arff:"restpathNumberOfMiddleClicksDiff" csv:"restpathNumberOfMiddleClicksDiff"`
	RestPathNumberOfMiddleClicksSum                uint8   `arff:"restpathNumberOfMiddleClicksSum" csv:"restpathNumberOfMiddleClicksSum"`
	RestPathNumberOfMiddleClicksMean               float64 `arff:"restpathNumberOfMiddleClicksMean" csv:"restpathNumberOfMiddleClicksMean"`
	RestPathNumberOfMiddleClicksStdDev             float64 `arff:"restpathNumberOfMiddleClicksStdDev" csv:"restpathNumberOfMiddleClicksStdDev"`
	RestPathNumberOfMiddleClicksSkew               float64 `arff:"restpathNumberOfMiddleClicksSkew" csv:"restpathNumberOfMiddleClicksSkew"`
	RestPathNumberOfScrollsMin                     uint8   `arff:"restpathNumberOfScrollsMin" csv:"restpathNumberOfScrollsMin"`
	RestPathNumberOfScrollsMax                     uint8   `arff:"restpathNumberOfScrollsMax" csv:"restpathNumberOfScrollsMax"`
	RestPathNumberOfScrollsDiff                    uint8   `arff:"restpathNumberOfScrollsDiff" csv:"restpathNumberOfScrollsDiff"`
	RestPathNumberOfScrollsSum                     uint8   `arff:"restpathNumberOfScrollsSum" csv:"restpathNumberOfScrollsSum"`
	RestPathNumberOfScrollsMean                    float64 `arff:"restpathNumberOfScrollsMean" csv:"restpathNumberOfScrollsMean"`
	RestPathNumberOfScrollsStdDev                  float64 `arff:"restpathNumberOfScrollsStdDev" csv:"restpathNumberOfScrollsStdDev"`
	RestPathNumberOfScrollsSkew                    float64 `arff:"restpathNumberOfScrollsSkew" csv:"restpathNumberOfScrollsSkew"`
	RestPathBreakTimeTotalTimeRatioMin             float64 `arff:"restpathBreakTimeTotalTimeRatioMin" csv:"restpathBreakTimeTotalTimeRatioMin"`
	RestPathBreakTimeTotalTimeRatioMax             float64 `arff:"restpathBreakTimeTotalTimeRatioMax" csv:"restpathBreakTimeTotalTimeRatioMax"`
	RestPathBreakTimeTotalTimeRatioDiff            float64 `arff:"restpathBreakTimeTotalTimeRatioDiff" csv:"restpathBreakTimeTotalTimeRatioDiff"`
	RestPathBreakTimeTotalTimeRatioSum             float64 `arff:"restpathBreakTimeTotalTimeRatioSum" csv:"restpathBreakTimeTotalTimeRatioSum"`
	RestPathBreakTimeTotalTimeRatioMean            float64 `arff:"restpathBreakTimeTotalTimeRatioMean" csv:"restpathBreakTimeTotalTimeRatioMean"`
	RestPathBreakTimeTotalTimeRatioStdDev          float64 `arff:"restpathBreakTimeTotalTimeRatioStdDev" csv:"restpathBreakTimeTotalTimeRatioStdDev"`
	RestPathBreakTimeTotalTimeRatioSkew            float64 `arff:"restpathBreakTimeTotalTimeRatioSkew" csv:"restpathBreakTimeTotalTimeRatioSkew"`
	RestPathNumberOfBreaksMin                      uint16  `arff:"restpathNumberOfBreaksMin" csv:"restpathNumberOfBreaksMin"`
	RestPathNumberOfBreaksMax                      uint16  `arff:"restpathNumberOfBreaksMax" csv:"restpathNumberOfBreaksMax"`
	RestPathNumberOfBreaksDiff                     uint16  `arff:"restpathNumberOfBreaksDiff" csv:"restpathNumberOfBreaksDiff"`
	RestPathNumberOfBreaksSum                      uint16  `arff:"restpathNumberOfBreaksSum" csv:"restpathNumberOfBreaksSum"`
	RestPathNumberOfBreaksMean                     float64 `arff:"restpathNumberOfBreaksMean" csv:"restpathNumberOfBreaksMean"`
	RestPathNumberOfBreaksStdDev                   float64 `arff:"restpathNumberOfBreaksStdDev" csv:"restpathNumberOfBreaksStdDev"`
	RestPathNumberOfBreaksSkew                     float64 `arff:"restpathNumberOfBreaksSkew" csv:"restpathNumberOfBreaksSkew"`
	RestPathDurationOfPathMin                      uint64  `arff:"restpathDurationOfPathMin" csv:"restpathDurationOfPathMin"`
	RestPathDurationOfPathMax                      uint64  `arff:"restpathDurationOfPathMax" csv:"restpathDurationOfPathMax"`
	RestPathDurationOfPathDiff                     uint64  `arff:"restpathDurationOfPathDiff" csv:"restpathDurationOfPathDiff"`
	RestPathDurationOfPathSum                      uint64  `arff:"restpathDurationOfPathSum" csv:"restpathDurationOfPathSum"`
	RestPathDurationOfPathMean                     float64 `arff:"restpathDurationOfPathMean" csv:"restpathDurationOfPathMean"`
	RestPathDurationOfPathStdDev                   float64 `arff:"restpathDurationOfPathStdDev" csv:"restpathDurationOfPathStdDev"`
	RestPathDurationOfPathSkew                     float64 `arff:"restpathDurationOfPathSkew" csv:"restpathDurationOfPathSkew"`
	RestPathTimeBetweenClickAndMovementMin         uint64  `arff:"restpathTimeBetweenClickAndMovementMin" csv:"restpathTimeBetweenClickAndMovementMin"`
	RestPathTimeBetweenClickAndMovementMax         uint64  `arff:"restpathTimeBetweenClickAndMovementMax" csv:"restpathTimeBetweenClickAndMovementMax"`
	RestPathTimeBetweenClickAndMovementDiff        uint64  `arff:"restpathTimeBetweenClickAndMovementDiff" csv:"restpathTimeBetweenClickAndMovementDiff"`
	RestPathTimeBetweenClickAndMovementSum         uint64  `arff:"restpathTimeBetweenClickAndMovementSum" csv:"restpathTimeBetweenClickAndMovementSum"`
	RestPathTimeBetweenClickAndMovementMean        float64 `arff:"restpathTimeBetweenClickAndMovementMean" csv:"restpathTimeBetweenClickAndMovementMean"`
	RestPathTimeBetweenClickAndMovementStdDev      float64 `arff:"restpathTimeBetweenClickAndMovementStdDev" csv:"restpathTimeBetweenClickAndMovementStdDev"`
	RestPathTimeBetweenClickAndMovementSkew        float64 `arff:"restpathTimeBetweenClickAndMovementSkew" csv:"restpathTimeBetweenClickAndMovementSkew"`
	RestPathTimeBetweeenMovementAndDownClickMin    uint64  `arff:"restpathTimeBetweeenMovementAndDownClickMin" csv:"restpathTimeBetweeenMovementAndDownClickMin"`
	RestPathTimeBetweeenMovementAndDownClickMax    uint64  `arff:"restpathTimeBetweeenMovementAndDownClickMax" csv:"restpathTimeBetweeenMovementAndDownClickMax"`
	RestPathTimeBetweeenMovementAndDownClickDiff   uint64  `arff:"restpathTimeBetweeenMovementAndDownClickDiff" csv:"restpathTimeBetweeenMovementAndDownClickDiff"`
	RestPathTimeBetweeenMovementAndDownClickSum    uint64  `arff:"restpathTimeBetweeenMovementAndDownClickSum" csv:"restpathTimeBetweeenMovementAndDownClickSum"`
	RestPathTimeBetweeenMovementAndDownClickMean   float64 `arff:"restpathTimeBetweeenMovementAndDownClickMean" csv:"restpathTimeBetweeenMovementAndDownClickMean"`
	RestPathTimeBetweeenMovementAndDownClickStdDev float64 `arff:"restpathTimeBetweeenMovementAndDownClickStdDev" csv:"restpathTimeBetweeenMovementAndDownClickStdDev"`
	RestPathTimeBetweeenMovementAndDownClickSkew   float64 `arff:"restpathTimeBetweeenMovementAndDownClickSkew" csv:"restpathTimeBetweeenMovementAndDownClickSkew"`
	RestPathAngleStartEndPointMin                  float64 `arff:"restpathAngleStartEndPointMin" csv:"restpathAngleStartEndPointMin"`
	RestPathAngleStartEndPointMax                  float64 `arff:"restpathAngleStartEndPointMax" csv:"restpathAngleStartEndPointMax"`
	RestPathAngleStartEndPointDiff                 float64 `arff:"restpathAngleStartEndPointDiff" csv:"restpathAngleStartEndPointDiff"`
	RestPathAngleStartEndPointSum                  float64 `arff:"restpathAngleStartEndPointSum" csv:"restpathAngleStartEndPointSum"`
	RestPathAngleStartEndPointMean                 float64 `arff:"restpathAngleStartEndPointMean" csv:"restpathAngleStartEndPointMean"`
	RestPathAngleStartEndPointStdDev               float64 `arff:"restpathAngleStartEndPointStdDev" csv:"restpathAngleStartEndPointStdDev"`
	RestPathAngleStartEndPointSkew                 float64 `arff:"restpathAngleStartEndPointSkew" csv:"restpathAngleStartEndPointSkew"`
	RestPathMeanAccelerationMin                    float64 `arff:"restpathMeanAccelerationMin" csv:"restpathMeanAccelerationMin"`
	RestPathMeanAccelerationMax                    float64 `arff:"restpathMeanAccelerationMax" csv:"restpathMeanAccelerationMax"`
	RestPathMeanAccelerationDiff                   float64 `arff:"restpathMeanAccelerationDiff" csv:"restpathMeanAccelerationDiff"`
	RestPathMeanAccelerationSum                    float64 `arff:"restpathMeanAccelerationSum" csv:"restpathMeanAccelerationSum"`
	RestPathMeanAccelerationMean                   float64 `arff:"restpathMeanAccelerationMean" csv:"restpathMeanAccelerationMean"`
	RestPathMeanAccelerationStdDev                 float64 `arff:"restpathMeanAccelerationStdDev" csv:"restpathMeanAccelerationStdDev"`
	RestPathMeanAccelerationSkew                   float64 `arff:"restpathMeanAccelerationSkew" csv:"restpathMeanAccelerationSkew"`
	RestPathMeanVelocityMin                        float64 `arff:"restpathMeanVelocityMin" csv:"restpathMeanVelocityMin"`
	RestPathMeanVelocityMax                        float64 `arff:"restpathMeanVelocityMax" csv:"restpathMeanVelocityMax"`
	RestPathMeanVelocityDiff                       float64 `arff:"restpathMeanVelocityDiff" csv:"restpathMeanVelocityDiff"`
	RestPathMeanVelocitySum                        float64 `arff:"restpathMeanVelocitySum" csv:"restpathMeanVelocitySum"`
	RestPathMeanVelocityMean                       float64 `arff:"restpathMeanVelocityMean" csv:"restpathMeanVelocityMean"`
	RestPathMeanVelocityStdDev                     float64 `arff:"restpathMeanVelocityStdDev" csv:"restpathMeanVelocityStdDev"`
	RestPathMeanVelocitySkew                       float64 `arff:"restpathMeanVelocitySkew" csv:"restpathMeanVelocitySkew"`
	RestPathDistanceStartEndPointMin               float64 `arff:"restpathDistanceStartEndPointMin" csv:"restpathDistanceStartEndPointMin"`
	RestPathDistanceStartEndPointMax               float64 `arff:"restpathDistanceStartEndPointMax" csv:"restpathDistanceStartEndPointMax"`
	RestPathDistanceStartEndPointDiff              float64 `arff:"restpathDistanceStartEndPointDiff" csv:"restpathDistanceStartEndPointDiff"`
	RestPathDistanceStartEndPointSum               float64 `arff:"restpathDistanceStartEndPointSum" csv:"restpathDistanceStartEndPointSum"`
	RestPathDistanceStartEndPointMean              float64 `arff:"restpathDistanceStartEndPointMean" csv:"restpathDistanceStartEndPointMean"`
	RestPathDistanceStartEndPointStdDev            float64 `arff:"restpathDistanceStartEndPointStdDev" csv:"restpathDistanceStartEndPointStdDev"`
	RestPathDistanceStartEndPointSkew              float64 `arff:"restpathDistanceStartEndPointSkew" csv:"restpathDistanceStartEndPointSkew"`
	RestPathDistanceSumMin                         float64 `arff:"restpathDistanceSumMin" csv:"restpathDistanceSumMin"`
	RestPathDistanceSumMax                         float64 `arff:"restpathDistanceSumMax" csv:"restpathDistanceSumMax"`
	RestPathDistanceSumDiff                        float64 `arff:"restpathDistanceSumDiff" csv:"restpathDistanceSumDiff"`
	RestPathDistanceSumSum                         float64 `arff:"restpathDistanceSumSum" csv:"restpathDistanceSumSum"`
	RestPathDistanceSumMean                        float64 `arff:"restpathDistanceSumMean" csv:"restpathDistanceSumMean"`
	RestPathDistanceSumStdDev                      float64 `arff:"restpathDistanceSumStdDev" csv:"restpathDistanceSumStdDev"`
	RestPathDistanceSumSkew                        float64 `arff:"restpathDistanceSumSkew" csv:"restpathDistanceSumSkew"`
	RestPathNumberOfMovementPointsMin              uint16  `arff:"restpathNumberOfMovementPointsMin" csv:"restpathNumberOfMovementPointsMin"`
	RestPathNumberOfMovementPointsMax              uint16  `arff:"restpathNumberOfMovementPointsMax" csv:"restpathNumberOfMovementPointsMax"`
	RestPathNumberOfMovementPointsDiff             uint16  `arff:"restpathNumberOfMovementPointsDiff" csv:"restpathNumberOfMovementPointsDiff"`
	RestPathNumberOfMovementPointsSum              uint16  `arff:"restpathNumberOfMovementPointsSum" csv:"restpathNumberOfMovementPointsSum"`
	RestPathNumberOfMovementPointsMean             float64 `arff:"restpathNumberOfMovementPointsMean" csv:"restpathNumberOfMovementPointsMean"`
	RestPathNumberOfMovementPointsStdDev           float64 `arff:"restpathNumberOfMovementPointsStdDev" csv:"restpathNumberOfMovementPointsStdDev"`
	RestPathNumberOfMovementPointsSkew             float64 `arff:"restpathNumberOfMovementPointsSkew" csv:"restpathNumberOfMovementPointsSkew"`
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
	// Plugins
	header.AddAttr("plugin-chrome-pdf", arff.Numeric, nil)
	header.AddAttr("plugin-chrome-pdf-viewer", arff.Numeric, nil)
	header.AddAttr("plugin-chromium-pdf", arff.Numeric, nil)
	header.AddAttr("plugin-chromium-pdf-viewer", arff.Numeric, nil)
	header.AddAttr("plugin-native-client", arff.Numeric, nil)
	header.AddAttr("plugin-more", arff.Numeric, nil)
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
	header.AddAttr("webglGCModel", arff.Nominal, gcModels)
	header.AddAttr("webglGCVendor", arff.Nominal, gcVendors)
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
	// Task Features
	header.AddAttr("failedTasks", arff.Numeric, nil)
	header.AddAttr("taskAmount", arff.Numeric, nil)
	header.AddAttr("noCaptchaAmount", arff.Numeric, nil)
	header.AddAttr("similarImageAmount", arff.Numeric, nil)
	header.AddAttr("dynamicSimilarImageAmount", arff.Numeric, nil)
	header.AddAttr("objectIdentificationAmount", arff.Numeric, nil)
	header.AddAttr("candidatesStopSign", arff.Numeric, nil)
	header.AddAttr("candidatesSpeedLimit", arff.Numeric, nil)
	header.AddAttr("candidatesStreetName", arff.Numeric, nil)
	header.AddAttr("candidatesOther", arff.Numeric, nil)
	header.AddAttr("candidatesCar", arff.Numeric, nil)
	header.AddAttr("candidatesBridge", arff.Numeric, nil)
	header.AddAttr("candidatesRoad", arff.Numeric, nil)
	header.AddAttr("candidatesBoundingBox", arff.Numeric, nil)
	header.AddAttr("durationLastCellEventToStop", arff.Numeric, nil)
	header.AddAttr("taskPicturesSelectedMin", arff.Numeric, nil)
	header.AddAttr("taskPicturesSelectedMax", arff.Numeric, nil)
	header.AddAttr("taskPicturesSelectedDiff", arff.Numeric, nil)
	header.AddAttr("taskPicturesSelectedSum", arff.Numeric, nil)
	header.AddAttr("taskPicturesSelectedMean", arff.Numeric, nil)
	header.AddAttr("taskPicturesSelectedStdDev", arff.Numeric, nil)
	header.AddAttr("taskPicturesSelectedSkew", arff.Numeric, nil)
	header.AddAttr("cellEventsMin", arff.Numeric, nil)
	header.AddAttr("cellEventsMax", arff.Numeric, nil)
	header.AddAttr("cellEventsDiff", arff.Numeric, nil)
	header.AddAttr("cellEventsSum", arff.Numeric, nil)
	header.AddAttr("cellEventsMean", arff.Numeric, nil)
	header.AddAttr("cellEventsStdDev", arff.Numeric, nil)
	header.AddAttr("cellEventsSkew", arff.Numeric, nil)
	header.AddAttr("taskDurationFirstLastEventMin", arff.Numeric, nil)
	header.AddAttr("taskDurationFirstLastEventMax", arff.Numeric, nil)
	header.AddAttr("taskDurationFirstLastEventDiff", arff.Numeric, nil)
	header.AddAttr("taskDurationFirstLastEventSum", arff.Numeric, nil)
	header.AddAttr("taskDurationFirstLastEventMean", arff.Numeric, nil)
	header.AddAttr("taskDurationFirstLastEventStdDev", arff.Numeric, nil)
	header.AddAttr("taskDurationFirstLastEventSkew", arff.Numeric, nil)
	header.AddAttr("taskChangedSelectionsMin", arff.Numeric, nil)
	header.AddAttr("taskChangedSelectionsMax", arff.Numeric, nil)
	header.AddAttr("taskChangedSelectionsDiff", arff.Numeric, nil)
	header.AddAttr("taskChangedSelectionsSum", arff.Numeric, nil)
	header.AddAttr("taskChangedSelectionsMean", arff.Numeric, nil)
	header.AddAttr("taskChangedSelectionsStdDev", arff.Numeric, nil)
	header.AddAttr("taskChangedSelectionsSkew", arff.Numeric, nil)
	header.AddAttr("amountDisappearingImagesMin", arff.Numeric, nil)
	header.AddAttr("amountDisappearingImagesMax", arff.Numeric, nil)
	header.AddAttr("amountDisappearingImagesDiff", arff.Numeric, nil)
	header.AddAttr("amountDisappearingImagesSum", arff.Numeric, nil)
	header.AddAttr("amountDisappearingImagesMean", arff.Numeric, nil)
	header.AddAttr("amountDisappearingImagesStdDev", arff.Numeric, nil)
	header.AddAttr("amountDisappearingImagesSkew", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXMin", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXMax", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXDiff", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXSum", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXMean", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXStdDev", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXSkew", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYMin", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYMax", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYDiff", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYSum", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYMean", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYStdDev", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYSkew", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXDistMin", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXDistMax", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXDistDiff", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXDistSum", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXDistMean", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXDistStdDev", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesXDistSkew", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYDistMin", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYDistMax", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYDistDiff", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYDistSum", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYDistMean", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYDistStdDev", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesYDistSkew", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesPairwiseDistMin", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesPairwiseDistMax", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesPairwiseDistDiff", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesPairwiseDistSum", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesPairwiseDistMean", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesPairwiseDistStdDev", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesPairwiseDistSkew", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistStartEndMin", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistStartEndMax", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistStartEndDiff", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistStartEndSum", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistStartEndMean", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistStartEndStdDev", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistStartEndSkew", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistSumMin", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistSumMax", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistSumDiff", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistSumSum", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistSumMean", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistSumStdDev", arff.Numeric, nil)
	header.AddAttr("chosenTaskImagesDistSumSkew", arff.Numeric, nil)
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
	header.AddAttr("restpathPairwiseDistanceMin", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDistanceMax", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDistanceDiff", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDistanceSum", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDistanceMean", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDistanceStdDev", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDistanceSkew", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseVelocityMin", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseVelocityMax", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseVelocityDiff", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseVelocitySum", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseVelocityMean", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseVelocitySkew", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAccelerationMin", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAccelerationMax", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAccelerationDiff", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAccelerationSum", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAccelerationMean", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAccelerationStdDev", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAccelerationSkew", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngleMin", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngleMax", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngleDiff", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngleSum", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngleMean", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngleStdDev", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngleSkew", arff.Numeric, nil)
	header.AddAttr("restpathAngleBetweenMovementAndStartEndMin", arff.Numeric, nil)
	header.AddAttr("restpathAngleBetweenMovementAndStartEndMax", arff.Numeric, nil)
	header.AddAttr("restpathAngleBetweenMovementAndStartEndDiff", arff.Numeric, nil)
	header.AddAttr("restpathAngleBetweenMovementAndStartEndSum", arff.Numeric, nil)
	header.AddAttr("restpathAngleBetweenMovementAndStartEndMean", arff.Numeric, nil)
	header.AddAttr("restpathAngleBetweenMovementAndStartEndStdDev", arff.Numeric, nil)
	header.AddAttr("restpathAngleBetweenMovementAndStartEndSkew", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngularVelocityMin", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngularVelocityMax", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngularVelocityDiff", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngularVelocitySum", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngularVelocityMean", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngularVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseAngularVelocitySkew", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDurationMin", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDurationMax", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDurationDiff", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDurationSum", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDurationMean", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDurationStdDev", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseDurationSkew", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndReleaseMin", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndReleaseMax", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndReleaseDiff", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndReleaseSum", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndReleaseMean", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndReleaseStdDev", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndReleaseSkew", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimesMin", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimesMax", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimesDiff", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimesSum", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimesMean", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimesStdDev", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimesSkew", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickDistanceMin", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickDistanceMax", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickDistanceDiff", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickDistanceSum", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickDistanceMean", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickDistanceStdDev", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickDistanceSkew", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickVelocityMin", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickVelocityMax", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickVelocityDiff", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickVelocitySum", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickVelocityMean", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickVelocitySkew", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAccelerationMin", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAccelerationMax", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAccelerationDiff", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAccelerationSum", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAccelerationMean", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAccelerationStdDev", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAccelerationSkew", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAngleMin", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAngleMax", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAngleDiff", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAngleSum", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAngleMean", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAngleStdDev", arff.Numeric, nil)
	header.AddAttr("restpathMovementDuringClickAngleSkew", arff.Numeric, nil)
	header.AddAttr("restpathScrollDXMin", arff.Numeric, nil)
	header.AddAttr("restpathScrollDXMax", arff.Numeric, nil)
	header.AddAttr("restpathScrollDXDiff", arff.Numeric, nil)
	header.AddAttr("restpathScrollDXSum", arff.Numeric, nil)
	header.AddAttr("restpathScrollDXMean", arff.Numeric, nil)
	header.AddAttr("restpathScrollDXStdDev", arff.Numeric, nil)
	header.AddAttr("restpathScrollDXSkew", arff.Numeric, nil)
	header.AddAttr("restpathScrollDYMin", arff.Numeric, nil)
	header.AddAttr("restpathScrollDYMax", arff.Numeric, nil)
	header.AddAttr("restpathScrollDYDiff", arff.Numeric, nil)
	header.AddAttr("restpathScrollDYSum", arff.Numeric, nil)
	header.AddAttr("restpathScrollDYMean", arff.Numeric, nil)
	header.AddAttr("restpathScrollDYStdDev", arff.Numeric, nil)
	header.AddAttr("restpathScrollDYSkew", arff.Numeric, nil)
	header.AddAttr("restpathScrollDZMin", arff.Numeric, nil)
	header.AddAttr("restpathScrollDZMax", arff.Numeric, nil)
	header.AddAttr("restpathScrollDZDiff", arff.Numeric, nil)
	header.AddAttr("restpathScrollDZSum", arff.Numeric, nil)
	header.AddAttr("restpathScrollDZMean", arff.Numeric, nil)
	header.AddAttr("restpathScrollDZStdDev", arff.Numeric, nil)
	header.AddAttr("restpathScrollDZSkew", arff.Numeric, nil)
	header.AddAttr("restpathScrollDMMin", arff.Numeric, nil)
	header.AddAttr("restpathScrollDMMax", arff.Numeric, nil)
	header.AddAttr("restpathScrollDMDiff", arff.Numeric, nil)
	header.AddAttr("restpathScrollDMSum", arff.Numeric, nil)
	header.AddAttr("restpathScrollDMMean", arff.Numeric, nil)
	header.AddAttr("restpathScrollDMStdDev", arff.Numeric, nil)
	header.AddAttr("restpathScrollDMSkew", arff.Numeric, nil)
	header.AddAttr("restpathXPointsMin", arff.Numeric, nil)
	header.AddAttr("restpathXPointsMax", arff.Numeric, nil)
	header.AddAttr("restpathXPointsDiff", arff.Numeric, nil)
	header.AddAttr("restpathXPointsSum", arff.Numeric, nil)
	header.AddAttr("restpathXPointsMean", arff.Numeric, nil)
	header.AddAttr("restpathXPointsStdDev", arff.Numeric, nil)
	header.AddAttr("restpathXPointsSkew", arff.Numeric, nil)
	header.AddAttr("restpathYPointsMin", arff.Numeric, nil)
	header.AddAttr("restpathYPointsMax", arff.Numeric, nil)
	header.AddAttr("restpathYPointsDiff", arff.Numeric, nil)
	header.AddAttr("restpathYPointsSum", arff.Numeric, nil)
	header.AddAttr("restpathYPointsMean", arff.Numeric, nil)
	header.AddAttr("restpathYPointsStdDev", arff.Numeric, nil)
	header.AddAttr("restpathYPointsSkew", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXVelocityMin", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXVelocityMax", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXVelocityDiff", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXVelocitySum", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXVelocityMean", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXVelocitySkew", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYVelocityMin", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYVelocityMax", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYVelocityDiff", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYVelocitySum", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYVelocityMean", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYVelocitySkew", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXDistanceMin", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXDistanceMax", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXDistanceDiff", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXDistanceSum", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXDistanceMean", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXDistanceStdDev", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseXDistanceSkew", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYDistanceMin", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYDistanceMax", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYDistanceDiff", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYDistanceSum", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYDistanceMean", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYDistanceStdDev", arff.Numeric, nil)
	header.AddAttr("restpathPairwiseYDistanceSkew", arff.Numeric, nil)
	header.AddAttr("restpathSumYDistanceMin", arff.Numeric, nil)
	header.AddAttr("restpathSumYDistanceMax", arff.Numeric, nil)
	header.AddAttr("restpathSumYDistanceDiff", arff.Numeric, nil)
	header.AddAttr("restpathSumYDistanceSum", arff.Numeric, nil)
	header.AddAttr("restpathSumYDistanceMean", arff.Numeric, nil)
	header.AddAttr("restpathSumYDistanceStdDev", arff.Numeric, nil)
	header.AddAttr("restpathSumYDistanceSkew", arff.Numeric, nil)
	header.AddAttr("restpathMeanYVelocityMin", arff.Numeric, nil)
	header.AddAttr("restpathMeanYVelocityMax", arff.Numeric, nil)
	header.AddAttr("restpathMeanYVelocityDiff", arff.Numeric, nil)
	header.AddAttr("restpathMeanYVelocitySum", arff.Numeric, nil)
	header.AddAttr("restpathMeanYVelocityMean", arff.Numeric, nil)
	header.AddAttr("restpathMeanYVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("restpathMeanYVelocitySkew", arff.Numeric, nil)
	header.AddAttr("restpathSumXDistanceMin", arff.Numeric, nil)
	header.AddAttr("restpathSumXDistanceMax", arff.Numeric, nil)
	header.AddAttr("restpathSumXDistanceDiff", arff.Numeric, nil)
	header.AddAttr("restpathSumXDistanceSum", arff.Numeric, nil)
	header.AddAttr("restpathSumXDistanceMean", arff.Numeric, nil)
	header.AddAttr("restpathSumXDistanceStdDev", arff.Numeric, nil)
	header.AddAttr("restpathSumXDistanceSkew", arff.Numeric, nil)
	header.AddAttr("restpathMeanXVelocityMin", arff.Numeric, nil)
	header.AddAttr("restpathMeanXVelocityMax", arff.Numeric, nil)
	header.AddAttr("restpathMeanXVelocityDiff", arff.Numeric, nil)
	header.AddAttr("restpathMeanXVelocitySum", arff.Numeric, nil)
	header.AddAttr("restpathMeanXVelocityMean", arff.Numeric, nil)
	header.AddAttr("restpathMeanXVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("restpathMeanXVelocitySkew", arff.Numeric, nil)
	header.AddAttr("restpathStraightnessMin", arff.Numeric, nil)
	header.AddAttr("restpathStraightnessMax", arff.Numeric, nil)
	header.AddAttr("restpathStraightnessDiff", arff.Numeric, nil)
	header.AddAttr("restpathStraightnessSum", arff.Numeric, nil)
	header.AddAttr("restpathStraightnessMean", arff.Numeric, nil)
	header.AddAttr("restpathStraightnessStdDev", arff.Numeric, nil)
	header.AddAttr("restpathStraightnessSkew", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfRightClicksMin", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfRightClicksMax", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfRightClicksDiff", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfRightClicksSum", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfRightClicksMean", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfRightClicksStdDev", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfRightClicksSkew", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMiddleClicksMin", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMiddleClicksMax", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMiddleClicksDiff", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMiddleClicksSum", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMiddleClicksMean", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMiddleClicksStdDev", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMiddleClicksSkew", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfScrollsMin", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfScrollsMax", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfScrollsDiff", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfScrollsSum", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfScrollsMean", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfScrollsStdDev", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfScrollsSkew", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimeTotalTimeRatioMin", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimeTotalTimeRatioMax", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimeTotalTimeRatioDiff", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimeTotalTimeRatioSum", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimeTotalTimeRatioMean", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimeTotalTimeRatioStdDev", arff.Numeric, nil)
	header.AddAttr("restpathBreakTimeTotalTimeRatioSkew", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfBreaksMin", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfBreaksMax", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfBreaksDiff", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfBreaksSum", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfBreaksMean", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfBreaksStdDev", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfBreaksSkew", arff.Numeric, nil)
	header.AddAttr("restpathDurationOfPathMin", arff.Numeric, nil)
	header.AddAttr("restpathDurationOfPathMax", arff.Numeric, nil)
	header.AddAttr("restpathDurationOfPathDiff", arff.Numeric, nil)
	header.AddAttr("restpathDurationOfPathSum", arff.Numeric, nil)
	header.AddAttr("restpathDurationOfPathMean", arff.Numeric, nil)
	header.AddAttr("restpathDurationOfPathStdDev", arff.Numeric, nil)
	header.AddAttr("restpathDurationOfPathSkew", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndMovementMin", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndMovementMax", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndMovementDiff", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndMovementSum", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndMovementMean", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndMovementStdDev", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweenClickAndMovementSkew", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweeenMovementAndDownClickMin", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweeenMovementAndDownClickMax", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweeenMovementAndDownClickDiff", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweeenMovementAndDownClickSum", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweeenMovementAndDownClickMean", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweeenMovementAndDownClickStdDev", arff.Numeric, nil)
	header.AddAttr("restpathTimeBetweeenMovementAndDownClickSkew", arff.Numeric, nil)
	header.AddAttr("restpathAngleStartEndPointMin", arff.Numeric, nil)
	header.AddAttr("restpathAngleStartEndPointMax", arff.Numeric, nil)
	header.AddAttr("restpathAngleStartEndPointDiff", arff.Numeric, nil)
	header.AddAttr("restpathAngleStartEndPointSum", arff.Numeric, nil)
	header.AddAttr("restpathAngleStartEndPointMean", arff.Numeric, nil)
	header.AddAttr("restpathAngleStartEndPointStdDev", arff.Numeric, nil)
	header.AddAttr("restpathAngleStartEndPointSkew", arff.Numeric, nil)
	header.AddAttr("restpathMeanAccelerationMin", arff.Numeric, nil)
	header.AddAttr("restpathMeanAccelerationMax", arff.Numeric, nil)
	header.AddAttr("restpathMeanAccelerationDiff", arff.Numeric, nil)
	header.AddAttr("restpathMeanAccelerationSum", arff.Numeric, nil)
	header.AddAttr("restpathMeanAccelerationMean", arff.Numeric, nil)
	header.AddAttr("restpathMeanAccelerationStdDev", arff.Numeric, nil)
	header.AddAttr("restpathMeanAccelerationSkew", arff.Numeric, nil)
	header.AddAttr("restpathMeanVelocityMin", arff.Numeric, nil)
	header.AddAttr("restpathMeanVelocityMax", arff.Numeric, nil)
	header.AddAttr("restpathMeanVelocityDiff", arff.Numeric, nil)
	header.AddAttr("restpathMeanVelocitySum", arff.Numeric, nil)
	header.AddAttr("restpathMeanVelocityMean", arff.Numeric, nil)
	header.AddAttr("restpathMeanVelocityStdDev", arff.Numeric, nil)
	header.AddAttr("restpathMeanVelocitySkew", arff.Numeric, nil)
	header.AddAttr("restpathDistanceStartEndPointMin", arff.Numeric, nil)
	header.AddAttr("restpathDistanceStartEndPointMax", arff.Numeric, nil)
	header.AddAttr("restpathDistanceStartEndPointDiff", arff.Numeric, nil)
	header.AddAttr("restpathDistanceStartEndPointSum", arff.Numeric, nil)
	header.AddAttr("restpathDistanceStartEndPointMean", arff.Numeric, nil)
	header.AddAttr("restpathDistanceStartEndPointStdDev", arff.Numeric, nil)
	header.AddAttr("restpathDistanceStartEndPointSkew", arff.Numeric, nil)
	header.AddAttr("restpathDistanceSumMin", arff.Numeric, nil)
	header.AddAttr("restpathDistanceSumMax", arff.Numeric, nil)
	header.AddAttr("restpathDistanceSumDiff", arff.Numeric, nil)
	header.AddAttr("restpathDistanceSumSum", arff.Numeric, nil)
	header.AddAttr("restpathDistanceSumMean", arff.Numeric, nil)
	header.AddAttr("restpathDistanceSumStdDev", arff.Numeric, nil)
	header.AddAttr("restpathDistanceSumSkew", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMovementPointsMin", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMovementPointsMax", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMovementPointsDiff", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMovementPointsSum", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMovementPointsMean", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMovementPointsStdDev", arff.Numeric, nil)
	header.AddAttr("restpathNumberOfMovementPointsSkew", arff.Numeric, nil)

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
	features.ExtractGraphicCardInformation(data.WebGL)
	features.ExtractScreenInfo(data.Screen)
	features.ExtractFramePosition(data.FramePosition)
	features.ExtractMouseData(data)
	features.ExtractTaskData(data.TaskEvents, data.Time)
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
	features.ExtractPlugins(binfo)
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
	features.GenerateRestCalculation(restPaths)
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
	features.GenerateCheckboxCalculation(path)
}

// ExtractTaskData extracts relevant Features from the given tasks
func (features *ProcessedFeatures) ExtractTaskData(data []TaskData, t Time) {
	l := len(data)
	noSelected := make([]uint16, l)
	noEvents := make([]uint16, l)
	timesFirstLastEvent := make([]uint64, l)
	noChanged := make([]uint16, l) //Number of deselections/changed rectangles
	noDiss := make([]uint8, 0)
	xCoords := make([]uint8, 0)
	yCoords := make([]uint8, 0)
	xDist := make([]int16, 0)
	yDist := make([]int16, 0)
	dists := make([]float64, 0)
	distStartEnd := make([]float64, l)
	distSums := make([]float64, 0)

	for i, task := range data {
		size := uint8(3)

		if task.Failed {
			features.AmountFailedTasks++
		}
		switch task.Type {
		case TaskTypeNO:
			features.AmountNo++

		case TaskTypeSIM:
			features.AmountSimImg++
		case TaskTypeDYN:
			features.AmountDynSimImg++
			noDiss = append(noDiss, uint8(len(task.Images))-1)
		case TaskTypeOBJ:
			features.AmountObjIdent++
			size = uint8(4)
		}
		switch task.Candidate {
		case "stop-sign":
			features.AmountCandidatesStopSign++
		case "speed-limit":
			features.AmountCandidatesSpeedLimit++
		case "street-name":
			features.AmountCandidatesStreetName++
		case "other":
			features.AmountCandidatesOther++
		case "car":
			features.AmountCandidatesCar++
		case "bridge":
			features.AmountCandidatesBridge++
		case "road":
			features.AmountCandidatesRoad++
		case "bounding-box":
			features.AmountCandidatesBoundingbox++
		}

		for _, row := range task.Selected {
			for _, cell := range row {
				if cell {
					noSelected[i]++
				}
			}
		}

		noEvents[i] = uint16(len(task.Events))

		if noEvents[i] > 0 {
			lastEvent := task.Events[noEvents[i]-1]
			firstEvent := task.Events[0]
			timesFirstLastEvent[i] = lastEvent.Time - firstEvent.Time
			xL, yL := Convert1DTo2D(lastEvent.ID, size)
			xF, yF := Convert1DTo2D(firstEvent.ID, size)
			distStartEnd[i] = math.Abs(float64(xL)-float64(xF)) + math.Abs(float64(yL)-float64(yF))
		}

		distSum := 0.0

		for j, event := range task.Events {
			x, y := Convert1DTo2D(event.ID, size)

			if j > 0 {
				prevX := xCoords[len(xCoords)-1]
				prevY := yCoords[len(yCoords)-1]
				xDist = append(xDist, int16(x)-int16(prevX))
				yDist = append(yDist, int16(y)-int16(prevY))
				manhattanDist := math.Abs(float64(x)-float64(prevX)) + math.Abs(float64(y)-float64(prevY))
				dists = append(dists, manhattanDist)
				distSum += manhattanDist
			}
			xCoords = append(xCoords, x)
			yCoords = append(yCoords, y)

			if !event.Selected {
				noChanged[i]++
			}
		}

		distSums = append(distSums, distSum)

	}

	if l > 0 {
		lastTask := data[l-1]
		evL := len(lastTask.Events)
		if evL > 0 {
			lastEvent := lastTask.Events[evL-1]
			features.DurationLastCellEventToStop = t.End - lastEvent.Time
		}
	}
	features.AmountTasks = uint8(l)
	features.calcTaskPictureSelected(noSelected)
	features.calcAmountCellEvents(noEvents)
	features.calcTaskChangedSelections(noChanged)
	features.calcTaskDurationFirstLastEvent(timesFirstLastEvent)
	features.calcAmountDisappearingImages(noDiss)
	features.calcChosenTaskImagesX(xCoords)
	features.calcChosenTaskImagesY(yCoords)
	features.calcChosenTaskImagesXDist(xDist)
	features.calcChosenTaskImagesYDist(yDist)
	features.calcChosenTaskImagesPairwiseDistance(dists)
	features.calcChosenTaskImagesDistStartEnd(distStartEnd)
	features.calcChosenTaskImagesDistSum(distSums)
}

// Convert1DTo2D converts the given position number to the (x,y) coordinates in the table with the given size
func Convert1DTo2D(position, size uint8) (uint8, uint8) {
	return position % size, position / size
}

// calcTaskPictureSelected calculates the features of the TaskPicturesSelected vector from the given vector
func (features *ProcessedFeatures) calcTaskPictureSelected(vec []uint16) {
	features.TaskPicturesSelectedMax = MaxUint16(vec)
	features.TaskPicturesSelectedMin = MinUint16(vec)
	features.TaskPicturesSelectedDiff = features.TaskPicturesSelectedMax - features.TaskPicturesSelectedMin
	features.TaskPicturesSelectedSum = SumUint16(vec)
	features.TaskPicturesSelectedMean = Mean(ConvertUint16ToFloat64(vec), float64(features.TaskPicturesSelectedSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.TaskPicturesSelectedStdDev, features.TaskPicturesSelectedSkew, _ = rnd.StatMoments(ConvertUint16ToFloat64(vec))
}

// calcAmountCellEvents calculates the features of the AmountCellEvents vector from the given vector
func (features *ProcessedFeatures) calcAmountCellEvents(vec []uint16) {
	features.AmountCellEventsMax = MaxUint16(vec)
	features.AmountCellEventsMin = MinUint16(vec)
	features.AmountCellEventsDiff = features.AmountCellEventsMax - features.AmountCellEventsMin
	features.AmountCellEventsSum = SumUint16(vec)
	features.AmountCellEventsMean = Mean(ConvertUint16ToFloat64(vec), float64(features.AmountCellEventsSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.AmountCellEventsStdDev, features.AmountCellEventsSkew, _ = rnd.StatMoments(ConvertUint16ToFloat64(vec))
}

// calcTaskChangedSelections calculates the features of the TaskChangedSelections vector from the given vector
func (features *ProcessedFeatures) calcTaskChangedSelections(vec []uint16) {
	features.TaskChangedSelectionsMax = MaxUint16(vec)
	features.TaskChangedSelectionsMin = MinUint16(vec)
	features.TaskChangedSelectionsDiff = features.TaskChangedSelectionsMax - features.TaskChangedSelectionsMin
	features.TaskChangedSelectionsSum = SumUint16(vec)
	features.TaskChangedSelectionsMean = Mean(ConvertUint16ToFloat64(vec), float64(features.TaskChangedSelectionsSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.TaskChangedSelectionsStdDev, features.TaskChangedSelectionsSkew, _ = rnd.StatMoments(ConvertUint16ToFloat64(vec))
}

// calcTaskDurationFirstLastEvent calculates the features of the TaskDurationFirstLastEvent vector from the given vector
func (features *ProcessedFeatures) calcTaskDurationFirstLastEvent(vec []uint64) {
	features.TaskDurationFirstLastEventMax = MaxUint64(vec)
	features.TaskDurationFirstLastEventMin = MinUint64(vec)
	features.TaskDurationFirstLastEventDiff = features.TaskDurationFirstLastEventMax - features.TaskDurationFirstLastEventMin
	features.TaskDurationFirstLastEventSum = SumUint64(vec)
	features.TaskDurationFirstLastEventMean = Mean(ConvertUint64ToFloat64(vec), float64(features.TaskDurationFirstLastEventSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.TaskDurationFirstLastEventStdDev, features.TaskDurationFirstLastEventSkew, _ = rnd.StatMoments(ConvertUint64ToFloat64(vec))
}

// calcAmountDisappearingImages calculates the features of the AmountDisappearingImages vector from the given vector
func (features *ProcessedFeatures) calcAmountDisappearingImages(vec []uint8) {
	features.AmountDisappearingImagesMax = MaxUint8(vec)
	features.AmountDisappearingImagesMin = MinUint8(vec)
	features.AmountDisappearingImagesDiff = features.AmountDisappearingImagesMax - features.AmountDisappearingImagesMin
	features.AmountDisappearingImagesSum = SumUint8(vec)
	features.AmountDisappearingImagesMean = Mean(ConvertUint8ToFloat64(vec), float64(features.AmountDisappearingImagesSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.AmountDisappearingImagesStdDev, features.AmountDisappearingImagesSkew, _ = rnd.StatMoments(ConvertUint8ToFloat64(vec))
}

// calcChosenTaskImagesX calculates the features of the ChosenTaskImagesX vector from the given vector
func (features *ProcessedFeatures) calcChosenTaskImagesX(vec []uint8) {
	features.ChosenTaskImagesXMax = MaxUint8(vec)
	features.ChosenTaskImagesXMin = MinUint8(vec)
	features.ChosenTaskImagesXDiff = features.ChosenTaskImagesXMax - features.ChosenTaskImagesXMin
	features.ChosenTaskImagesXSum = SumUint8(vec)
	features.ChosenTaskImagesXMean = Mean(ConvertUint8ToFloat64(vec), float64(features.ChosenTaskImagesXSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.ChosenTaskImagesXStdDev, features.ChosenTaskImagesXSkew, _ = rnd.StatMoments(ConvertUint8ToFloat64(vec))
}

// calcChosenTaskImagesY calculates the features of the ChosenTaskImagesY vector from the given vector
func (features *ProcessedFeatures) calcChosenTaskImagesY(vec []uint8) {
	features.ChosenTaskImagesYMax = MaxUint8(vec)
	features.ChosenTaskImagesYMin = MinUint8(vec)
	features.ChosenTaskImagesYDiff = features.ChosenTaskImagesYMax - features.ChosenTaskImagesYMin
	features.ChosenTaskImagesYSum = SumUint8(vec)
	features.ChosenTaskImagesYMean = Mean(ConvertUint8ToFloat64(vec), float64(features.ChosenTaskImagesYSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.ChosenTaskImagesYStdDev, features.ChosenTaskImagesYSkew, _ = rnd.StatMoments(ConvertUint8ToFloat64(vec))
}

// calcChosenTaskImagesYDist calculates the features of the ChosenTaskImagesYDist vector from the given vector
func (features *ProcessedFeatures) calcChosenTaskImagesYDist(vec []int16) {
	features.ChosenTaskImagesYDistMax = MaxInt(vec)
	features.ChosenTaskImagesYDistMin = MinInt(vec)
	features.ChosenTaskImagesYDistDiff = features.ChosenTaskImagesYDistMax - features.ChosenTaskImagesYDistMin
	features.ChosenTaskImagesYDistSum = SumInt(vec)
	features.ChosenTaskImagesYDistMean = Mean(ConvertInt16ToFloat64(vec), float64(features.ChosenTaskImagesYDistSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.ChosenTaskImagesYDistStdDev, features.ChosenTaskImagesYDistSkew, _ = rnd.StatMoments(ConvertInt16ToFloat64(vec))
}

// calcChosenTaskImagesXDist calculates the features of the ChosenTaskImagesXDist vector from the given vector
func (features *ProcessedFeatures) calcChosenTaskImagesXDist(vec []int16) {
	features.ChosenTaskImagesXDistMax = MaxInt(vec)
	features.ChosenTaskImagesXDistMin = MinInt(vec)
	features.ChosenTaskImagesXDistDiff = features.ChosenTaskImagesXDistMax - features.ChosenTaskImagesXDistMin
	features.ChosenTaskImagesXDistSum = SumInt(vec)
	features.ChosenTaskImagesXDistMean = Mean(ConvertInt16ToFloat64(vec), float64(features.ChosenTaskImagesXDistSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.ChosenTaskImagesXDistStdDev, features.ChosenTaskImagesXDistSkew, _ = rnd.StatMoments(ConvertInt16ToFloat64(vec))
}

// calcChosenTaskImagesPairwiseDist calculates the features of the ChosenTaskImagesPairwiseDistance vector from the given vector
func (features *ProcessedFeatures) calcChosenTaskImagesPairwiseDistance(vec []float64) {
	features.ChosenTaskImagesPairwiseDistMin, features.ChosenTaskImagesPairwiseDistMean, features.ChosenTaskImagesPairwiseDistMax, features.ChosenTaskImagesPairwiseDistStdDev = rnd.StatBasic(vec, true)
	features.ChosenTaskImagesPairwiseDistSum = SumFloat(vec)
	if len(vec) == 1 {
		features.ChosenTaskImagesPairwiseDistMin, features.ChosenTaskImagesPairwiseDistMean, features.ChosenTaskImagesPairwiseDistMax = vec[0], vec[0], vec[0]
	}
	features.ChosenTaskImagesPairwiseDistDiff = features.ChosenTaskImagesPairwiseDistMax - features.ChosenTaskImagesPairwiseDistMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.ChosenTaskImagesPairwiseDistSkew, _ = rnd.StatMoments(vec)
}

// calcChosenTaskImagesDistStartEnd calculates the features of the ChosenTaskImagesDistStartEnd vector from the given vector
func (features *ProcessedFeatures) calcChosenTaskImagesDistStartEnd(vec []float64) {
	features.ChosenTaskImagesDistStartEndMin, features.ChosenTaskImagesDistStartEndMean, features.ChosenTaskImagesDistStartEndMax, features.ChosenTaskImagesDistStartEndStdDev = rnd.StatBasic(vec, true)
	features.ChosenTaskImagesDistStartEndSum = SumFloat(vec)
	if len(vec) == 1 {
		features.ChosenTaskImagesDistStartEndMin, features.ChosenTaskImagesDistStartEndMean, features.ChosenTaskImagesDistStartEndMax = vec[0], vec[0], vec[0]
	}
	features.ChosenTaskImagesDistStartEndDiff = features.ChosenTaskImagesDistStartEndMax - features.ChosenTaskImagesDistStartEndMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.ChosenTaskImagesDistStartEndSkew, _ = rnd.StatMoments(vec)
}

// calcChosenTaskImagesDistSum calculates the features of the ChosenTaskImagesDistSum vector from the given vector
func (features *ProcessedFeatures) calcChosenTaskImagesDistSum(vec []float64) {
	features.ChosenTaskImagesDistSumMin, features.ChosenTaskImagesDistSumMean, features.ChosenTaskImagesDistSumMax, features.ChosenTaskImagesDistSumStdDev = rnd.StatBasic(vec, true)
	features.ChosenTaskImagesDistSumSum = SumFloat(vec)
	if len(vec) == 1 {
		features.ChosenTaskImagesDistSumMin, features.ChosenTaskImagesDistSumMean, features.ChosenTaskImagesDistSumMax = vec[0], vec[0], vec[0]
	}
	features.ChosenTaskImagesDistSumDiff = features.ChosenTaskImagesDistSumMax - features.ChosenTaskImagesDistSumMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.ChosenTaskImagesDistSumSkew, _ = rnd.StatMoments(vec)
}

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

// calcRestPathPairwiseDistanceFeatures calculates the features of the PairwiseDistance vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathPairwiseDistanceFeatures(col *collection) {
	features.RestPathPairwiseDistanceMin, features.RestPathPairwiseDistanceMean, features.RestPathPairwiseDistanceMax, features.RestPathPairwiseDistanceStdDev = rnd.StatBasic(col.PairwiseDistance, true)
	features.RestPathPairwiseDistanceSum = SumFloat(col.PairwiseDistance)
	if len(col.PairwiseDistance) == 1 {
		features.RestPathPairwiseDistanceMin, features.RestPathPairwiseDistanceMean, features.RestPathPairwiseDistanceMax = col.PairwiseDistance[0], col.PairwiseDistance[0], col.PairwiseDistance[0]
	}
	features.RestPathPairwiseDistanceDiff = features.RestPathPairwiseDistanceMax - features.RestPathPairwiseDistanceMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathPairwiseDistanceSkew, _ = rnd.StatMoments(col.PairwiseDistance)
}

// calcRestPathPairwiseVelocityFeatures calculates the features of the PairwiseVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathPairwiseVelocityFeatures(col *collection) {
	features.RestPathPairwiseVelocityMin, features.RestPathPairwiseVelocityMean, features.RestPathPairwiseVelocityMax, features.RestPathPairwiseVelocityStdDev = rnd.StatBasic(col.PairwiseVelocity, true)
	features.RestPathPairwiseVelocitySum = SumFloat(col.PairwiseVelocity)
	if len(col.PairwiseVelocity) == 1 {
		features.RestPathPairwiseVelocityMin, features.RestPathPairwiseVelocityMean, features.RestPathPairwiseVelocityMax = col.PairwiseVelocity[0], col.PairwiseVelocity[0], col.PairwiseVelocity[0]
	}
	features.RestPathPairwiseVelocityDiff = features.RestPathPairwiseVelocityMax - features.RestPathPairwiseVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathPairwiseVelocitySkew, _ = rnd.StatMoments(col.PairwiseVelocity)
}

// calcRestPathPairwiseAccelerationFeatures calculates the features of the PairwiseAcceleration vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathPairwiseAccelerationFeatures(col *collection) {
	features.RestPathPairwiseAccelerationMin, features.RestPathPairwiseAccelerationMean, features.RestPathPairwiseAccelerationMax, features.RestPathPairwiseAccelerationStdDev = rnd.StatBasic(col.PairwiseAcceleration, true)
	features.RestPathPairwiseAccelerationSum = SumFloat(col.PairwiseAcceleration)
	if len(col.PairwiseAcceleration) == 1 {
		features.RestPathPairwiseAccelerationMin, features.RestPathPairwiseAccelerationMean, features.RestPathPairwiseAccelerationMax = col.PairwiseAcceleration[0], col.PairwiseAcceleration[0], col.PairwiseAcceleration[0]
	}
	features.RestPathPairwiseAccelerationDiff = features.RestPathPairwiseAccelerationMax - features.RestPathPairwiseAccelerationMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathPairwiseAccelerationSkew, _ = rnd.StatMoments(col.PairwiseAcceleration)
}

// calcRestPathPairwiseAngleFeatures calculates the features of the PairwiseAngle vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathPairwiseAngleFeatures(col *collection) {
	features.RestPathPairwiseAngleMin, features.RestPathPairwiseAngleMean, features.RestPathPairwiseAngleMax, features.RestPathPairwiseAngleStdDev = rnd.StatBasic(col.PairwiseAngle, true)
	features.RestPathPairwiseAngleSum = SumFloat(col.PairwiseAngle)
	if len(col.PairwiseAngle) == 1 {
		features.RestPathPairwiseAngleMin, features.RestPathPairwiseAngleMean, features.RestPathPairwiseAngleMax = col.PairwiseAngle[0], col.PairwiseAngle[0], col.PairwiseAngle[0]
	}
	features.RestPathPairwiseAngleDiff = features.RestPathPairwiseAngleMax - features.RestPathPairwiseAngleMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathPairwiseAngleSkew, _ = rnd.StatMoments(col.PairwiseAngle)
}

// calcRestPathAngleBetweenMovementAndStartEndFeatures calculates the features of the AngleBetweenMovementAndStartEnd vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathAngleBetweenMovementAndStartEndFeatures(col *collection) {
	features.RestPathAngleBetweenMovementAndStartEndMin, features.RestPathAngleBetweenMovementAndStartEndMean, features.RestPathAngleBetweenMovementAndStartEndMax, features.RestPathAngleBetweenMovementAndStartEndStdDev = rnd.StatBasic(col.AngleBetweenMovementAndStartEnd, true)
	features.RestPathAngleBetweenMovementAndStartEndSum = SumFloat(col.AngleBetweenMovementAndStartEnd)
	if len(col.AngleBetweenMovementAndStartEnd) == 1 {
		features.RestPathAngleBetweenMovementAndStartEndMin, features.RestPathAngleBetweenMovementAndStartEndMean, features.RestPathAngleBetweenMovementAndStartEndMax = col.AngleBetweenMovementAndStartEnd[0], col.AngleBetweenMovementAndStartEnd[0], col.AngleBetweenMovementAndStartEnd[0]
	}
	features.RestPathAngleBetweenMovementAndStartEndDiff = features.RestPathAngleBetweenMovementAndStartEndMax - features.RestPathAngleBetweenMovementAndStartEndMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathAngleBetweenMovementAndStartEndSkew, _ = rnd.StatMoments(col.AngleBetweenMovementAndStartEnd)
}

// calcRestPathPairwiseAngularVelocityFeatures calculates the features of the PairwiseAngularVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathPairwiseAngularVelocityFeatures(col *collection) {
	features.RestPathPairwiseAngularVelocityMin, features.RestPathPairwiseAngularVelocityMean, features.RestPathPairwiseAngularVelocityMax, features.RestPathPairwiseAngularVelocityStdDev = rnd.StatBasic(col.PairwiseAngularVelocity, true)
	features.RestPathPairwiseAngularVelocitySum = SumFloat(col.PairwiseAngularVelocity)
	if len(col.PairwiseAngularVelocity) == 1 {
		features.RestPathPairwiseAngularVelocityMin, features.RestPathPairwiseAngularVelocityMean, features.RestPathPairwiseAngularVelocityMax = col.PairwiseAngularVelocity[0], col.PairwiseAngularVelocity[0], col.PairwiseAngularVelocity[0]
	}
	features.RestPathPairwiseAngularVelocityDiff = features.RestPathPairwiseAngularVelocityMax - features.RestPathPairwiseAngularVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathPairwiseAngularVelocitySkew, _ = rnd.StatMoments(col.PairwiseAngularVelocity)
}

// calcRestPathPairwiseDurationFeatures calculates the features of the PairwiseDuration vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathPairwiseDurationFeatures(col *collection) {
	features.RestPathPairwiseDurationMax = MaxUint64(col.PairwiseDuration)
	features.RestPathPairwiseDurationMin = MinUint64(col.PairwiseDuration)
	features.RestPathPairwiseDurationDiff = features.RestPathPairwiseDurationMax - features.RestPathPairwiseDurationMin
	features.RestPathPairwiseDurationSum = SumUint64(col.PairwiseDuration)
	features.RestPathPairwiseDurationMean = Mean(ConvertUint64ToFloat64(col.PairwiseDuration), float64(features.RestPathPairwiseDurationSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathPairwiseDurationStdDev, features.RestPathPairwiseDurationSkew, _ = rnd.StatMoments(ConvertUint64ToFloat64(col.PairwiseDuration))
}

// calcRestPathTimeBetweenClickAndReleaseFeatures calculates the features of the TimeBetweenClickAndRelease vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathTimeBetweenClickAndReleaseFeatures(col *collection) {
	features.RestPathTimeBetweenClickAndReleaseMax = MaxUint64(col.TimeBetweenClickAndRelease)
	features.RestPathTimeBetweenClickAndReleaseMin = MinUint64(col.TimeBetweenClickAndRelease)
	features.RestPathTimeBetweenClickAndReleaseDiff = features.RestPathTimeBetweenClickAndReleaseMax - features.RestPathTimeBetweenClickAndReleaseMin
	features.RestPathTimeBetweenClickAndReleaseSum = SumUint64(col.TimeBetweenClickAndRelease)
	features.RestPathTimeBetweenClickAndReleaseMean = Mean(ConvertUint64ToFloat64(col.TimeBetweenClickAndRelease), float64(features.RestPathTimeBetweenClickAndReleaseSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathTimeBetweenClickAndReleaseStdDev, features.RestPathTimeBetweenClickAndReleaseSkew, _ = rnd.StatMoments(ConvertUint64ToFloat64(col.TimeBetweenClickAndRelease))
}

// calcRestPathBreakTimesFeatures calculates the features of the BreakTimes vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathBreakTimesFeatures(col *collection) {
	features.RestPathBreakTimesMax = MaxUint64(col.BreakTimes)
	features.RestPathBreakTimesMin = MinUint64(col.BreakTimes)
	features.RestPathBreakTimesDiff = features.RestPathBreakTimesMax - features.RestPathBreakTimesMin
	features.RestPathBreakTimesSum = SumUint64(col.BreakTimes)
	features.RestPathBreakTimesMean = Mean(ConvertUint64ToFloat64(col.BreakTimes), float64(features.RestPathBreakTimesSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathBreakTimesStdDev, features.RestPathBreakTimesSkew, _ = rnd.StatMoments(ConvertUint64ToFloat64(col.BreakTimes))
}

// calcRestPathMovementDuringClickDistanceFeatures calculates the features of the MovementDuringClickDistance vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathMovementDuringClickDistanceFeatures(col *collection) {
	features.RestPathMovementDuringClickDistanceMin, features.RestPathMovementDuringClickDistanceMean, features.RestPathMovementDuringClickDistanceMax, features.RestPathMovementDuringClickDistanceStdDev = rnd.StatBasic(col.MovementDuringClickDistance, true)
	features.RestPathMovementDuringClickDistanceSum = SumFloat(col.MovementDuringClickDistance)
	if len(col.MovementDuringClickDistance) == 1 {
		features.RestPathMovementDuringClickDistanceMin, features.RestPathMovementDuringClickDistanceMean, features.RestPathMovementDuringClickDistanceMax = col.MovementDuringClickDistance[0], col.MovementDuringClickDistance[0], col.MovementDuringClickDistance[0]
	}
	features.RestPathMovementDuringClickDistanceDiff = features.RestPathMovementDuringClickDistanceMax - features.RestPathMovementDuringClickDistanceMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathMovementDuringClickDistanceSkew, _ = rnd.StatMoments(col.MovementDuringClickDistance)
}

// calcRestPathMovementDuringClickVelocityFeatures calculates the features of the MovementDuringClickVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathMovementDuringClickVelocityFeatures(col *collection) {
	features.RestPathMovementDuringClickVelocityMin, features.RestPathMovementDuringClickVelocityMean, features.RestPathMovementDuringClickVelocityMax, features.RestPathMovementDuringClickVelocityStdDev = rnd.StatBasic(col.MovementDuringClickVelocity, true)
	features.RestPathMovementDuringClickVelocitySum = SumFloat(col.MovementDuringClickVelocity)
	if len(col.MovementDuringClickVelocity) == 1 {
		features.RestPathMovementDuringClickVelocityMin, features.RestPathMovementDuringClickVelocityMean, features.RestPathMovementDuringClickVelocityMax = col.MovementDuringClickVelocity[0], col.MovementDuringClickVelocity[0], col.MovementDuringClickVelocity[0]
	}
	features.RestPathMovementDuringClickVelocityDiff = features.RestPathMovementDuringClickVelocityMax - features.RestPathMovementDuringClickVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathMovementDuringClickVelocitySkew, _ = rnd.StatMoments(col.MovementDuringClickVelocity)
}

// calcRestPathMovementDuringClickAccelerationFeatures calculates the features of the MovementDuringClickAcceleration vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathMovementDuringClickAccelerationFeatures(col *collection) {
	features.RestPathMovementDuringClickAccelerationMin, features.RestPathMovementDuringClickAccelerationMean, features.RestPathMovementDuringClickAccelerationMax, features.RestPathMovementDuringClickAccelerationStdDev = rnd.StatBasic(col.MovementDuringClickAcceleration, true)
	features.RestPathMovementDuringClickAccelerationSum = SumFloat(col.MovementDuringClickAcceleration)
	if len(col.MovementDuringClickAcceleration) == 1 {
		features.RestPathMovementDuringClickAccelerationMin, features.RestPathMovementDuringClickAccelerationMean, features.RestPathMovementDuringClickAccelerationMax = col.MovementDuringClickAcceleration[0], col.MovementDuringClickAcceleration[0], col.MovementDuringClickAcceleration[0]
	}
	features.RestPathMovementDuringClickAccelerationDiff = features.RestPathMovementDuringClickAccelerationMax - features.RestPathMovementDuringClickAccelerationMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathMovementDuringClickAccelerationSkew, _ = rnd.StatMoments(col.MovementDuringClickAcceleration)
}

// calcRestPathMovementDuringClickAngleFeatures calculates the features of the MovementDuringClickAngle vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathMovementDuringClickAngleFeatures(col *collection) {
	features.RestPathMovementDuringClickAngleMin, features.RestPathMovementDuringClickAngleMean, features.RestPathMovementDuringClickAngleMax, features.RestPathMovementDuringClickAngleStdDev = rnd.StatBasic(col.MovementDuringClickAngle, true)
	features.RestPathMovementDuringClickAngleSum = SumFloat(col.MovementDuringClickAngle)
	if len(col.MovementDuringClickAngle) == 1 {
		features.RestPathMovementDuringClickAngleMin, features.RestPathMovementDuringClickAngleMean, features.RestPathMovementDuringClickAngleMax = col.MovementDuringClickAngle[0], col.MovementDuringClickAngle[0], col.MovementDuringClickAngle[0]
	}
	features.RestPathMovementDuringClickAngleDiff = features.RestPathMovementDuringClickAngleMax - features.RestPathMovementDuringClickAngleMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathMovementDuringClickAngleSkew, _ = rnd.StatMoments(col.MovementDuringClickAngle)
}

// calcRestPathScrollDXFeatures calculates the features of the ScrollDX vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathScrollDXFeatures(col *collection) {
	features.RestPathScrollDXMin, features.RestPathScrollDXMean, features.RestPathScrollDXMax, features.RestPathScrollDXStdDev = rnd.StatBasic(col.ScrollDX, true)
	features.RestPathScrollDXSum = SumFloat(col.ScrollDX)
	if len(col.ScrollDX) == 1 {
		features.RestPathScrollDXMin, features.RestPathScrollDXMean, features.RestPathScrollDXMax = col.ScrollDX[0], col.ScrollDX[0], col.ScrollDX[0]
	}
	features.RestPathScrollDXDiff = features.RestPathScrollDXMax - features.RestPathScrollDXMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathScrollDXSkew, _ = rnd.StatMoments(col.ScrollDX)
}

// calcRestPathScrollDYFeatures calculates the features of the ScrollDY vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathScrollDYFeatures(col *collection) {
	features.RestPathScrollDYMin, features.RestPathScrollDYMean, features.RestPathScrollDYMax, features.RestPathScrollDYStdDev = rnd.StatBasic(col.ScrollDY, true)
	features.RestPathScrollDYSum = SumFloat(col.ScrollDY)
	if len(col.ScrollDY) == 1 {
		features.RestPathScrollDYMin, features.RestPathScrollDYMean, features.RestPathScrollDYMax = col.ScrollDY[0], col.ScrollDY[0], col.ScrollDY[0]
	}
	features.RestPathScrollDYDiff = features.RestPathScrollDYMax - features.RestPathScrollDYMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathScrollDYSkew, _ = rnd.StatMoments(col.ScrollDY)
}

// calcRestPathScrollDZFeatures calculates the features of the ScrollDZ vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathScrollDZFeatures(col *collection) {
	features.RestPathScrollDZMin, features.RestPathScrollDZMean, features.RestPathScrollDZMax, features.RestPathScrollDZStdDev = rnd.StatBasic(col.ScrollDZ, true)
	features.RestPathScrollDZSum = SumFloat(col.ScrollDZ)
	if len(col.ScrollDZ) == 1 {
		features.RestPathScrollDZMin, features.RestPathScrollDZMean, features.RestPathScrollDZMax = col.ScrollDZ[0], col.ScrollDZ[0], col.ScrollDZ[0]
	}
	features.RestPathScrollDZDiff = features.RestPathScrollDZMax - features.RestPathScrollDZMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathScrollDZSkew, _ = rnd.StatMoments(col.ScrollDZ)
}

// calcRestPathScrollDMFeatures calculates the features of the ScrollDM vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathScrollDMFeatures(col *collection) {
	features.RestPathScrollDMMax = MaxUint8(col.ScrollDM)
	features.RestPathScrollDMMin = MinUint8(col.ScrollDM)
	features.RestPathScrollDMDiff = features.RestPathScrollDMMax - features.RestPathScrollDMMin
	features.RestPathScrollDMSum = SumUint8(col.ScrollDM)
	features.RestPathScrollDMMean = Mean(ConvertUint8ToFloat64(col.ScrollDM), float64(features.RestPathScrollDMSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathScrollDMStdDev, features.RestPathScrollDMSkew, _ = rnd.StatMoments(ConvertUint8ToFloat64(col.ScrollDM))
}

// calcRestPathXPointsFeatures calculates the features of the XPoints vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathXPointsFeatures(col *collection) {
	features.RestPathXPointsMin, features.RestPathXPointsMean, features.RestPathXPointsMax, features.RestPathXPointsStdDev = rnd.StatBasic(col.XPoints, true)
	features.RestPathXPointsSum = SumFloat(col.XPoints)
	if len(col.XPoints) == 1 {
		features.RestPathXPointsMin, features.RestPathXPointsMean, features.RestPathXPointsMax = col.XPoints[0], col.XPoints[0], col.XPoints[0]
	}
	features.RestPathXPointsDiff = features.RestPathXPointsMax - features.RestPathXPointsMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathXPointsSkew, _ = rnd.StatMoments(col.XPoints)
}

// calcRestPathYPointsFeatures calculates the features of the YPoints vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathYPointsFeatures(col *collection) {
	features.RestPathYPointsMin, features.RestPathYPointsMean, features.RestPathYPointsMax, features.RestPathYPointsStdDev = rnd.StatBasic(col.YPoints, true)
	features.RestPathYPointsSum = SumFloat(col.YPoints)
	if len(col.YPoints) == 1 {
		features.RestPathYPointsMin, features.RestPathYPointsMean, features.RestPathYPointsMax = col.YPoints[0], col.YPoints[0], col.YPoints[0]
	}
	features.RestPathYPointsDiff = features.RestPathYPointsMax - features.RestPathYPointsMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathYPointsSkew, _ = rnd.StatMoments(col.YPoints)
}

// calcRestPathPairwiseXVelocityFeatures calculates the features of the PairwiseXVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathPairwiseXVelocityFeatures(col *collection) {
	features.RestPathPairwiseXVelocityMin, features.RestPathPairwiseXVelocityMean, features.RestPathPairwiseXVelocityMax, features.RestPathPairwiseXVelocityStdDev = rnd.StatBasic(col.PairwiseXVelocity, true)
	features.RestPathPairwiseXVelocitySum = SumFloat(col.PairwiseXVelocity)
	if len(col.PairwiseXVelocity) == 1 {
		features.RestPathPairwiseXVelocityMin, features.RestPathPairwiseXVelocityMean, features.RestPathPairwiseXVelocityMax = col.PairwiseXVelocity[0], col.PairwiseXVelocity[0], col.PairwiseXVelocity[0]
	}
	features.RestPathPairwiseXVelocityDiff = features.RestPathPairwiseXVelocityMax - features.RestPathPairwiseXVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathPairwiseXVelocitySkew, _ = rnd.StatMoments(col.PairwiseXVelocity)
}

// calcRestPathPairwiseYVelocityFeatures calculates the features of the PairwiseYVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathPairwiseYVelocityFeatures(col *collection) {
	features.RestPathPairwiseYVelocityMin, features.RestPathPairwiseYVelocityMean, features.RestPathPairwiseYVelocityMax, features.RestPathPairwiseYVelocityStdDev = rnd.StatBasic(col.PairwiseYVelocity, true)
	features.RestPathPairwiseYVelocitySum = SumFloat(col.PairwiseYVelocity)
	if len(col.PairwiseYVelocity) == 1 {
		features.RestPathPairwiseYVelocityMin, features.RestPathPairwiseYVelocityMean, features.RestPathPairwiseYVelocityMax = col.PairwiseYVelocity[0], col.PairwiseYVelocity[0], col.PairwiseYVelocity[0]
	}
	features.RestPathPairwiseYVelocityDiff = features.RestPathPairwiseYVelocityMax - features.RestPathPairwiseYVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathPairwiseYVelocitySkew, _ = rnd.StatMoments(col.PairwiseYVelocity)
}

// calcRestPathPairwiseXDistanceFeatures calculates the features of the PairwiseXDistance vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathPairwiseXDistanceFeatures(col *collection) {
	features.RestPathPairwiseXDistanceMin, features.RestPathPairwiseXDistanceMean, features.RestPathPairwiseXDistanceMax, features.RestPathPairwiseXDistanceStdDev = rnd.StatBasic(col.PairwiseXDistance, true)
	features.RestPathPairwiseXDistanceSum = SumFloat(col.PairwiseXDistance)
	if len(col.PairwiseXDistance) == 1 {
		features.RestPathPairwiseXDistanceMin, features.RestPathPairwiseXDistanceMean, features.RestPathPairwiseXDistanceMax = col.PairwiseXDistance[0], col.PairwiseXDistance[0], col.PairwiseXDistance[0]
	}
	features.RestPathPairwiseXDistanceDiff = features.RestPathPairwiseXDistanceMax - features.RestPathPairwiseXDistanceMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathPairwiseXDistanceSkew, _ = rnd.StatMoments(col.PairwiseXDistance)
}

// calcRestPathPairwiseYDistanceFeatures calculates the features of the PairwiseYDistance vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathPairwiseYDistanceFeatures(col *collection) {
	features.RestPathPairwiseYDistanceMin, features.RestPathPairwiseYDistanceMean, features.RestPathPairwiseYDistanceMax, features.RestPathPairwiseYDistanceStdDev = rnd.StatBasic(col.PairwiseYDistance, true)
	features.RestPathPairwiseYDistanceSum = SumFloat(col.PairwiseYDistance)
	if len(col.PairwiseYDistance) == 1 {
		features.RestPathPairwiseYDistanceMin, features.RestPathPairwiseYDistanceMean, features.RestPathPairwiseYDistanceMax = col.PairwiseYDistance[0], col.PairwiseYDistance[0], col.PairwiseYDistance[0]
	}
	features.RestPathPairwiseYDistanceDiff = features.RestPathPairwiseYDistanceMax - features.RestPathPairwiseYDistanceMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathPairwiseYDistanceSkew, _ = rnd.StatMoments(col.PairwiseYDistance)
}

// calcRestPathSumYDistanceFeatures calculates the features of the SumYDistance vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathSumYDistanceFeatures(col *collection) {
	features.RestPathSumYDistanceMin, features.RestPathSumYDistanceMean, features.RestPathSumYDistanceMax, features.RestPathSumYDistanceStdDev = rnd.StatBasic(col.SumYDistance, true)
	features.RestPathSumYDistanceSum = SumFloat(col.SumYDistance)
	if len(col.SumYDistance) == 1 {
		features.RestPathSumYDistanceMin, features.RestPathSumYDistanceMean, features.RestPathSumYDistanceMax = col.SumYDistance[0], col.SumYDistance[0], col.SumYDistance[0]
	}
	features.RestPathSumYDistanceDiff = features.RestPathSumYDistanceMax - features.RestPathSumYDistanceMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathSumYDistanceSkew, _ = rnd.StatMoments(col.SumYDistance)
}

// calcRestPathMeanYVelocityFeatures calculates the features of the MeanYVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathMeanYVelocityFeatures(col *collection) {
	features.RestPathMeanYVelocityMin, features.RestPathMeanYVelocityMean, features.RestPathMeanYVelocityMax, features.RestPathMeanYVelocityStdDev = rnd.StatBasic(col.MeanYVelocity, true)
	features.RestPathMeanYVelocitySum = SumFloat(col.MeanYVelocity)
	if len(col.MeanYVelocity) == 1 {
		features.RestPathMeanYVelocityMin, features.RestPathMeanYVelocityMean, features.RestPathMeanYVelocityMax = col.MeanYVelocity[0], col.MeanYVelocity[0], col.MeanYVelocity[0]
	}
	features.RestPathMeanYVelocityDiff = features.RestPathMeanYVelocityMax - features.RestPathMeanYVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathMeanYVelocitySkew, _ = rnd.StatMoments(col.MeanYVelocity)
}

// calcRestPathSumXDistanceFeatures calculates the features of the SumXDistance vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathSumXDistanceFeatures(col *collection) {
	features.RestPathSumXDistanceMin, features.RestPathSumXDistanceMean, features.RestPathSumXDistanceMax, features.RestPathSumXDistanceStdDev = rnd.StatBasic(col.SumXDistance, true)
	features.RestPathSumXDistanceSum = SumFloat(col.SumXDistance)
	if len(col.SumXDistance) == 1 {
		features.RestPathSumXDistanceMin, features.RestPathSumXDistanceMean, features.RestPathSumXDistanceMax = col.SumXDistance[0], col.SumXDistance[0], col.SumXDistance[0]
	}
	features.RestPathSumXDistanceDiff = features.RestPathSumXDistanceMax - features.RestPathSumXDistanceMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathSumXDistanceSkew, _ = rnd.StatMoments(col.SumXDistance)
}

// calcRestPathMeanXVelocityFeatures calculates the features of the MeanXVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathMeanXVelocityFeatures(col *collection) {
	features.RestPathMeanXVelocityMin, features.RestPathMeanXVelocityMean, features.RestPathMeanXVelocityMax, features.RestPathMeanXVelocityStdDev = rnd.StatBasic(col.MeanXVelocity, true)
	features.RestPathMeanXVelocitySum = SumFloat(col.MeanXVelocity)
	if len(col.MeanXVelocity) == 1 {
		features.RestPathMeanXVelocityMin, features.RestPathMeanXVelocityMean, features.RestPathMeanXVelocityMax = col.MeanXVelocity[0], col.MeanXVelocity[0], col.MeanXVelocity[0]
	}
	features.RestPathMeanXVelocityDiff = features.RestPathMeanXVelocityMax - features.RestPathMeanXVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathMeanXVelocitySkew, _ = rnd.StatMoments(col.MeanXVelocity)
}

// calcRestPathStraightnessFeatures calculates the features of the Straightness vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathStraightnessFeatures(col *collection) {
	features.RestPathStraightnessMin, features.RestPathStraightnessMean, features.RestPathStraightnessMax, features.RestPathStraightnessStdDev = rnd.StatBasic(col.Straightness, true)
	features.RestPathStraightnessSum = SumFloat(col.Straightness)
	if len(col.Straightness) == 1 {
		features.RestPathStraightnessMin, features.RestPathStraightnessMean, features.RestPathStraightnessMax = col.Straightness[0], col.Straightness[0], col.Straightness[0]
	}
	features.RestPathStraightnessDiff = features.RestPathStraightnessMax - features.RestPathStraightnessMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathStraightnessSkew, _ = rnd.StatMoments(col.Straightness)
}

// calcRestPathNumberOfRightClicksFeatures calculates the features of the NumberOfRightClicks vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathNumberOfRightClicksFeatures(col *collection) {
	features.RestPathNumberOfRightClicksMax = MaxUint8(col.NumberOfRightClicks)
	features.RestPathNumberOfRightClicksMin = MinUint8(col.NumberOfRightClicks)
	features.RestPathNumberOfRightClicksDiff = features.RestPathNumberOfRightClicksMax - features.RestPathNumberOfRightClicksMin
	features.RestPathNumberOfRightClicksSum = SumUint8(col.NumberOfRightClicks)
	features.RestPathNumberOfRightClicksMean = Mean(ConvertUint8ToFloat64(col.NumberOfRightClicks), float64(features.RestPathNumberOfRightClicksSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathNumberOfRightClicksStdDev, features.RestPathNumberOfRightClicksSkew, _ = rnd.StatMoments(ConvertUint8ToFloat64(col.NumberOfRightClicks))
}

// calcRestPathNumberOfMiddleClicksFeatures calculates the features of the NumberOfMiddleClicks vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathNumberOfMiddleClicksFeatures(col *collection) {
	features.RestPathNumberOfMiddleClicksMax = MaxUint8(col.NumberOfMiddleClicks)
	features.RestPathNumberOfMiddleClicksMin = MinUint8(col.NumberOfMiddleClicks)
	features.RestPathNumberOfMiddleClicksDiff = features.RestPathNumberOfMiddleClicksMax - features.RestPathNumberOfMiddleClicksMin
	features.RestPathNumberOfMiddleClicksSum = SumUint8(col.NumberOfMiddleClicks)
	features.RestPathNumberOfMiddleClicksMean = Mean(ConvertUint8ToFloat64(col.NumberOfMiddleClicks), float64(features.RestPathNumberOfMiddleClicksSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathNumberOfMiddleClicksStdDev, features.RestPathNumberOfMiddleClicksSkew, _ = rnd.StatMoments(ConvertUint8ToFloat64(col.NumberOfMiddleClicks))
}

// calcRestPathNumberOfScrollsFeatures calculates the features of the NumberOfScrolls vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathNumberOfScrollsFeatures(col *collection) {
	features.RestPathNumberOfScrollsMax = MaxUint8(col.NumberOfScrolls)
	features.RestPathNumberOfScrollsMin = MinUint8(col.NumberOfScrolls)
	features.RestPathNumberOfScrollsDiff = features.RestPathNumberOfScrollsMax - features.RestPathNumberOfScrollsMin
	features.RestPathNumberOfScrollsSum = SumUint8(col.NumberOfScrolls)
	features.RestPathNumberOfScrollsMean = Mean(ConvertUint8ToFloat64(col.NumberOfScrolls), float64(features.RestPathNumberOfScrollsSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathNumberOfScrollsStdDev, features.RestPathNumberOfScrollsSkew, _ = rnd.StatMoments(ConvertUint8ToFloat64(col.NumberOfScrolls))
}

// calcRestPathBreakTimeTotalTimeRatioFeatures calculates the features of the BreakTimeTotalTimeRatio vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathBreakTimeTotalTimeRatioFeatures(col *collection) {
	features.RestPathBreakTimeTotalTimeRatioMin, features.RestPathBreakTimeTotalTimeRatioMean, features.RestPathBreakTimeTotalTimeRatioMax, features.RestPathBreakTimeTotalTimeRatioStdDev = rnd.StatBasic(col.BreakTimeTotalTimeRatio, true)
	features.RestPathBreakTimeTotalTimeRatioSum = SumFloat(col.BreakTimeTotalTimeRatio)
	if len(col.BreakTimeTotalTimeRatio) == 1 {
		features.RestPathBreakTimeTotalTimeRatioMin, features.RestPathBreakTimeTotalTimeRatioMean, features.RestPathBreakTimeTotalTimeRatioMax = col.BreakTimeTotalTimeRatio[0], col.BreakTimeTotalTimeRatio[0], col.BreakTimeTotalTimeRatio[0]
	}
	features.RestPathBreakTimeTotalTimeRatioDiff = features.RestPathBreakTimeTotalTimeRatioMax - features.RestPathBreakTimeTotalTimeRatioMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathBreakTimeTotalTimeRatioSkew, _ = rnd.StatMoments(col.BreakTimeTotalTimeRatio)
}

// calcRestPathNumberOfBreaksFeatures calculates the features of the NumberOfBreaks vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathNumberOfBreaksFeatures(col *collection) {
	features.RestPathNumberOfBreaksMax = MaxUint16(col.NumberOfBreaks)
	features.RestPathNumberOfBreaksMin = MinUint16(col.NumberOfBreaks)
	features.RestPathNumberOfBreaksDiff = features.RestPathNumberOfBreaksMax - features.RestPathNumberOfBreaksMin
	features.RestPathNumberOfBreaksSum = SumUint16(col.NumberOfBreaks)
	features.RestPathNumberOfBreaksMean = Mean(ConvertUint16ToFloat64(col.NumberOfBreaks), float64(features.RestPathNumberOfBreaksSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathNumberOfBreaksStdDev, features.RestPathNumberOfBreaksSkew, _ = rnd.StatMoments(ConvertUint16ToFloat64(col.NumberOfBreaks))
}

// calcRestPathDurationOfPathFeatures calculates the features of the DurationOfPath vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathDurationOfPathFeatures(col *collection) {
	features.RestPathDurationOfPathMax = MaxUint64(col.DurationOfPath)
	features.RestPathDurationOfPathMin = MinUint64(col.DurationOfPath)
	features.RestPathDurationOfPathDiff = features.RestPathDurationOfPathMax - features.RestPathDurationOfPathMin
	features.RestPathDurationOfPathSum = SumUint64(col.DurationOfPath)
	features.RestPathDurationOfPathMean = Mean(ConvertUint64ToFloat64(col.DurationOfPath), float64(features.RestPathDurationOfPathSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathDurationOfPathStdDev, features.RestPathDurationOfPathSkew, _ = rnd.StatMoments(ConvertUint64ToFloat64(col.DurationOfPath))
}

// calcRestPathTimeBetweenClickAndMovementFeatures calculates the features of the TimeBetweenClickAndMovement vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathTimeBetweenClickAndMovementFeatures(col *collection) {
	features.RestPathTimeBetweenClickAndMovementMax = MaxUint64(col.TimeBetweenClickAndMovement)
	features.RestPathTimeBetweenClickAndMovementMin = MinUint64(col.TimeBetweenClickAndMovement)
	features.RestPathTimeBetweenClickAndMovementDiff = features.RestPathTimeBetweenClickAndMovementMax - features.RestPathTimeBetweenClickAndMovementMin
	features.RestPathTimeBetweenClickAndMovementSum = SumUint64(col.TimeBetweenClickAndMovement)
	features.RestPathTimeBetweenClickAndMovementMean = Mean(ConvertUint64ToFloat64(col.TimeBetweenClickAndMovement), float64(features.RestPathTimeBetweenClickAndMovementSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathTimeBetweenClickAndMovementStdDev, features.RestPathTimeBetweenClickAndMovementSkew, _ = rnd.StatMoments(ConvertUint64ToFloat64(col.TimeBetweenClickAndMovement))
}

// calcRestPathTimeBetweeenMovementAndDownClickFeatures calculates the features of the TimeBetweeenMovementAndDownClick vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathTimeBetweeenMovementAndDownClickFeatures(col *collection) {
	features.RestPathTimeBetweeenMovementAndDownClickMax = MaxUint64(col.TimeBetweeenMovementAndDownClick)
	features.RestPathTimeBetweeenMovementAndDownClickMin = MinUint64(col.TimeBetweeenMovementAndDownClick)
	features.RestPathTimeBetweeenMovementAndDownClickDiff = features.RestPathTimeBetweeenMovementAndDownClickMax - features.RestPathTimeBetweeenMovementAndDownClickMin
	features.RestPathTimeBetweeenMovementAndDownClickSum = SumUint64(col.TimeBetweeenMovementAndDownClick)
	features.RestPathTimeBetweeenMovementAndDownClickMean = Mean(ConvertUint64ToFloat64(col.TimeBetweeenMovementAndDownClick), float64(features.RestPathTimeBetweeenMovementAndDownClickSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathTimeBetweeenMovementAndDownClickStdDev, features.RestPathTimeBetweeenMovementAndDownClickSkew, _ = rnd.StatMoments(ConvertUint64ToFloat64(col.TimeBetweeenMovementAndDownClick))
}

// calcRestPathAngleStartEndPointFeatures calculates the features of the AngleStartEndPoint vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathAngleStartEndPointFeatures(col *collection) {
	features.RestPathAngleStartEndPointMin, features.RestPathAngleStartEndPointMean, features.RestPathAngleStartEndPointMax, features.RestPathAngleStartEndPointStdDev = rnd.StatBasic(col.AngleStartEndPoint, true)
	features.RestPathAngleStartEndPointSum = SumFloat(col.AngleStartEndPoint)
	if len(col.AngleStartEndPoint) == 1 {
		features.RestPathAngleStartEndPointMin, features.RestPathAngleStartEndPointMean, features.RestPathAngleStartEndPointMax = col.AngleStartEndPoint[0], col.AngleStartEndPoint[0], col.AngleStartEndPoint[0]
	}
	features.RestPathAngleStartEndPointDiff = features.RestPathAngleStartEndPointMax - features.RestPathAngleStartEndPointMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathAngleStartEndPointSkew, _ = rnd.StatMoments(col.AngleStartEndPoint)
}

// calcRestPathMeanAccelerationFeatures calculates the features of the MeanAcceleration vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathMeanAccelerationFeatures(col *collection) {
	features.RestPathMeanAccelerationMin, features.RestPathMeanAccelerationMean, features.RestPathMeanAccelerationMax, features.RestPathMeanAccelerationStdDev = rnd.StatBasic(col.MeanAcceleration, true)
	features.RestPathMeanAccelerationSum = SumFloat(col.MeanAcceleration)
	if len(col.MeanAcceleration) == 1 {
		features.RestPathMeanAccelerationMin, features.RestPathMeanAccelerationMean, features.RestPathMeanAccelerationMax = col.MeanAcceleration[0], col.MeanAcceleration[0], col.MeanAcceleration[0]
	}
	features.RestPathMeanAccelerationDiff = features.RestPathMeanAccelerationMax - features.RestPathMeanAccelerationMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathMeanAccelerationSkew, _ = rnd.StatMoments(col.MeanAcceleration)
}

// calcRestPathMeanVelocityFeatures calculates the features of the MeanVelocity vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathMeanVelocityFeatures(col *collection) {
	features.RestPathMeanVelocityMin, features.RestPathMeanVelocityMean, features.RestPathMeanVelocityMax, features.RestPathMeanVelocityStdDev = rnd.StatBasic(col.MeanVelocity, true)
	features.RestPathMeanVelocitySum = SumFloat(col.MeanVelocity)
	if len(col.MeanVelocity) == 1 {
		features.RestPathMeanVelocityMin, features.RestPathMeanVelocityMean, features.RestPathMeanVelocityMax = col.MeanVelocity[0], col.MeanVelocity[0], col.MeanVelocity[0]
	}
	features.RestPathMeanVelocityDiff = features.RestPathMeanVelocityMax - features.RestPathMeanVelocityMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathMeanVelocitySkew, _ = rnd.StatMoments(col.MeanVelocity)
}

// calcRestPathDistanceStartEndPointFeatures calculates the features of the DistanceStartEndPoint vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathDistanceStartEndPointFeatures(col *collection) {
	features.RestPathDistanceStartEndPointMin, features.RestPathDistanceStartEndPointMean, features.RestPathDistanceStartEndPointMax, features.RestPathDistanceStartEndPointStdDev = rnd.StatBasic(col.DistanceStartEndPoint, true)
	features.RestPathDistanceStartEndPointSum = SumFloat(col.DistanceStartEndPoint)
	if len(col.DistanceStartEndPoint) == 1 {
		features.RestPathDistanceStartEndPointMin, features.RestPathDistanceStartEndPointMean, features.RestPathDistanceStartEndPointMax = col.DistanceStartEndPoint[0], col.DistanceStartEndPoint[0], col.DistanceStartEndPoint[0]
	}
	features.RestPathDistanceStartEndPointDiff = features.RestPathDistanceStartEndPointMax - features.RestPathDistanceStartEndPointMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathDistanceStartEndPointSkew, _ = rnd.StatMoments(col.DistanceStartEndPoint)
}

// calcRestPathDistanceSumFeatures calculates the features of the DistanceSum vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathDistanceSumFeatures(col *collection) {
	features.RestPathDistanceSumMin, features.RestPathDistanceSumMean, features.RestPathDistanceSumMax, features.RestPathDistanceSumStdDev = rnd.StatBasic(col.DistanceSum, true)
	features.RestPathDistanceSumSum = SumFloat(col.DistanceSum)
	if len(col.DistanceSum) == 1 {
		features.RestPathDistanceSumMin, features.RestPathDistanceSumMean, features.RestPathDistanceSumMax = col.DistanceSum[0], col.DistanceSum[0], col.DistanceSum[0]
	}
	features.RestPathDistanceSumDiff = features.RestPathDistanceSumMax - features.RestPathDistanceSumMin
	defer func() {
		recover()
	}()
	_, _, _, _, _, features.RestPathDistanceSumSkew, _ = rnd.StatMoments(col.DistanceSum)
}

// calcRestPathNumberOfMovementPointsFeatures calculates the features of the NumberOfMovementPoints vector from the given path; Auto-generated
func (features *ProcessedFeatures) calcRestPathNumberOfMovementPointsFeatures(col *collection) {
	features.RestPathNumberOfMovementPointsMax = MaxUint16(col.NumberOfMovementPoints)
	features.RestPathNumberOfMovementPointsMin = MinUint16(col.NumberOfMovementPoints)
	features.RestPathNumberOfMovementPointsDiff = features.RestPathNumberOfMovementPointsMax - features.RestPathNumberOfMovementPointsMin
	features.RestPathNumberOfMovementPointsSum = SumUint16(col.NumberOfMovementPoints)
	features.RestPathNumberOfMovementPointsMean = Mean(ConvertUint16ToFloat64(col.NumberOfMovementPoints), float64(features.RestPathNumberOfMovementPointsSum))
	defer func() {
		recover()
	}()
	_, _, _, _, features.RestPathNumberOfMovementPointsStdDev, features.RestPathNumberOfMovementPointsSkew, _ = rnd.StatMoments(ConvertUint16ToFloat64(col.NumberOfMovementPoints))
}

//collection is a struct to collect all the paths; Auto-generated
type collection struct {
	PairwiseDistance                 []float64
	PairwiseVelocity                 []float64
	PairwiseAcceleration             []float64
	PairwiseAngle                    []float64
	AngleBetweenMovementAndStartEnd  []float64
	PairwiseAngularVelocity          []float64
	PairwiseDuration                 []uint64
	TimeBetweenClickAndRelease       []uint64
	BreakTimes                       []uint64
	MovementDuringClickDistance      []float64
	MovementDuringClickVelocity      []float64
	MovementDuringClickAcceleration  []float64
	MovementDuringClickAngle         []float64
	ScrollDX                         []float64
	ScrollDY                         []float64
	ScrollDZ                         []float64
	ScrollDM                         []uint8
	XPoints                          []float64
	YPoints                          []float64
	PairwiseXVelocity                []float64
	PairwiseYVelocity                []float64
	PairwiseXDistance                []float64
	PairwiseYDistance                []float64
	SumYDistance                     []float64
	MeanYVelocity                    []float64
	SumXDistance                     []float64
	MeanXVelocity                    []float64
	Straightness                     []float64
	NumberOfRightClicks              []uint8
	NumberOfMiddleClicks             []uint8
	NumberOfScrolls                  []uint8
	BreakTimeTotalTimeRatio          []float64
	NumberOfBreaks                   []uint16
	DurationOfPath                   []uint64
	TimeBetweenClickAndMovement      []uint64
	TimeBetweeenMovementAndDownClick []uint64
	AngleStartEndPoint               []float64
	MeanAcceleration                 []float64
	MeanVelocity                     []float64
	DistanceStartEndPoint            []float64
	DistanceSum                      []float64
	NumberOfMovementPoints           []uint16
}

// GenerateRestCalculation caluclates and stores the Features of the RestPath; Auto-generated
func (features *ProcessedFeatures) GenerateRestCalculation(paths []*PathFeatures) {
	collected := &collection{}
	for _, v := range paths {
		collected.PairwiseDistance = append(collected.PairwiseDistance, v.PairwiseDistance...)
		collected.PairwiseVelocity = append(collected.PairwiseVelocity, v.PairwiseVelocity...)
		collected.PairwiseAcceleration = append(collected.PairwiseAcceleration, v.PairwiseAcceleration...)
		collected.PairwiseAngle = append(collected.PairwiseAngle, v.PairwiseAngle...)
		collected.AngleBetweenMovementAndStartEnd = append(collected.AngleBetweenMovementAndStartEnd, v.AngleBetweenMovementAndStartEnd...)
		collected.PairwiseAngularVelocity = append(collected.PairwiseAngularVelocity, v.PairwiseAngularVelocity...)
		collected.PairwiseDuration = append(collected.PairwiseDuration, v.PairwiseDuration...)
		collected.TimeBetweenClickAndRelease = append(collected.TimeBetweenClickAndRelease, v.TimeBetweenClickAndRelease...)
		collected.BreakTimes = append(collected.BreakTimes, v.BreakTimes...)
		collected.MovementDuringClickDistance = append(collected.MovementDuringClickDistance, v.MovementDuringClickDistance...)
		collected.MovementDuringClickVelocity = append(collected.MovementDuringClickVelocity, v.MovementDuringClickVelocity...)
		collected.MovementDuringClickAcceleration = append(collected.MovementDuringClickAcceleration, v.MovementDuringClickAcceleration...)
		collected.MovementDuringClickAngle = append(collected.MovementDuringClickAngle, v.MovementDuringClickAngle...)
		collected.ScrollDX = append(collected.ScrollDX, v.ScrollDX...)
		collected.ScrollDY = append(collected.ScrollDY, v.ScrollDY...)
		collected.ScrollDZ = append(collected.ScrollDZ, v.ScrollDZ...)
		collected.ScrollDM = append(collected.ScrollDM, v.ScrollDM...)
		collected.XPoints = append(collected.XPoints, v.XPoints...)
		collected.YPoints = append(collected.YPoints, v.YPoints...)
		collected.PairwiseXVelocity = append(collected.PairwiseXVelocity, v.PairwiseXVelocity...)
		collected.PairwiseYVelocity = append(collected.PairwiseYVelocity, v.PairwiseYVelocity...)
		collected.PairwiseXDistance = append(collected.PairwiseXDistance, v.PairwiseXDistance...)
		collected.PairwiseYDistance = append(collected.PairwiseYDistance, v.PairwiseYDistance...)
		collected.SumYDistance = append(collected.SumYDistance, v.SumYDistance)
		collected.MeanYVelocity = append(collected.MeanYVelocity, v.MeanYVelocity)
		collected.SumXDistance = append(collected.SumXDistance, v.SumXDistance)
		collected.MeanXVelocity = append(collected.MeanXVelocity, v.MeanXVelocity)
		collected.Straightness = append(collected.Straightness, v.Straightness)
		collected.NumberOfRightClicks = append(collected.NumberOfRightClicks, v.NumberOfRightClicks)
		collected.NumberOfMiddleClicks = append(collected.NumberOfMiddleClicks, v.NumberOfMiddleClicks)
		collected.NumberOfScrolls = append(collected.NumberOfScrolls, v.NumberOfScrolls)
		collected.BreakTimeTotalTimeRatio = append(collected.BreakTimeTotalTimeRatio, v.BreakTimeTotalTimeRatio)
		collected.NumberOfBreaks = append(collected.NumberOfBreaks, v.NumberOfBreaks)
		collected.DurationOfPath = append(collected.DurationOfPath, v.DurationOfPath)
		collected.TimeBetweenClickAndMovement = append(collected.TimeBetweenClickAndMovement, v.TimeBetweenClickAndMovement)
		collected.TimeBetweeenMovementAndDownClick = append(collected.TimeBetweeenMovementAndDownClick, v.TimeBetweeenMovementAndDownClick)
		collected.AngleStartEndPoint = append(collected.AngleStartEndPoint, v.AngleStartEndPoint)
		collected.MeanAcceleration = append(collected.MeanAcceleration, v.MeanAcceleration)
		collected.MeanVelocity = append(collected.MeanVelocity, v.MeanVelocity)
		collected.DistanceStartEndPoint = append(collected.DistanceStartEndPoint, v.DistanceStartEndPoint)
		collected.DistanceSum = append(collected.DistanceSum, v.DistanceSum)
		collected.NumberOfMovementPoints = append(collected.NumberOfMovementPoints, v.NumberOfMovementPoints)
	}
	features.calcRestPathPairwiseDistanceFeatures(collected)
	features.calcRestPathPairwiseVelocityFeatures(collected)
	features.calcRestPathPairwiseAccelerationFeatures(collected)
	features.calcRestPathPairwiseAngleFeatures(collected)
	features.calcRestPathAngleBetweenMovementAndStartEndFeatures(collected)
	features.calcRestPathPairwiseAngularVelocityFeatures(collected)
	features.calcRestPathPairwiseDurationFeatures(collected)
	features.calcRestPathTimeBetweenClickAndReleaseFeatures(collected)
	features.calcRestPathBreakTimesFeatures(collected)
	features.calcRestPathMovementDuringClickDistanceFeatures(collected)
	features.calcRestPathMovementDuringClickVelocityFeatures(collected)
	features.calcRestPathMovementDuringClickAccelerationFeatures(collected)
	features.calcRestPathMovementDuringClickAngleFeatures(collected)
	features.calcRestPathScrollDXFeatures(collected)
	features.calcRestPathScrollDYFeatures(collected)
	features.calcRestPathScrollDZFeatures(collected)
	features.calcRestPathScrollDMFeatures(collected)
	features.calcRestPathXPointsFeatures(collected)
	features.calcRestPathYPointsFeatures(collected)
	features.calcRestPathPairwiseXVelocityFeatures(collected)
	features.calcRestPathPairwiseYVelocityFeatures(collected)
	features.calcRestPathPairwiseXDistanceFeatures(collected)
	features.calcRestPathPairwiseYDistanceFeatures(collected)
	features.calcRestPathSumYDistanceFeatures(collected)
	features.calcRestPathMeanYVelocityFeatures(collected)
	features.calcRestPathSumXDistanceFeatures(collected)
	features.calcRestPathMeanXVelocityFeatures(collected)
	features.calcRestPathStraightnessFeatures(collected)
	features.calcRestPathNumberOfRightClicksFeatures(collected)
	features.calcRestPathNumberOfMiddleClicksFeatures(collected)
	features.calcRestPathNumberOfScrollsFeatures(collected)
	features.calcRestPathBreakTimeTotalTimeRatioFeatures(collected)
	features.calcRestPathNumberOfBreaksFeatures(collected)
	features.calcRestPathDurationOfPathFeatures(collected)
	features.calcRestPathTimeBetweenClickAndMovementFeatures(collected)
	features.calcRestPathTimeBetweeenMovementAndDownClickFeatures(collected)
	features.calcRestPathAngleStartEndPointFeatures(collected)
	features.calcRestPathMeanAccelerationFeatures(collected)
	features.calcRestPathMeanVelocityFeatures(collected)
	features.calcRestPathDistanceStartEndPointFeatures(collected)
	features.calcRestPathDistanceSumFeatures(collected)
	features.calcRestPathNumberOfMovementPointsFeatures(collected)
}
