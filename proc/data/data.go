// Package data provides all the structs for the gathered data from the Webextension
package data

// BrowserInfo contains additional information about the browser
type BrowserInfo struct {
	CookieEnabled bool     `json:"cookieEnabled"`
	CPUAmount     uint8    `json:"cpuNo"`
	DNT           string   `json:"dnt"` // true, false or "" (not set)
	Languages     []string `json:"lang"`
	MIMETypes     []string `json:"mimeTypes"`
	Platform      string   `json:"platform"`
	Plugins       []string `json:"plugins"`
	TimeZone      int16    `json:"timezone"`
	UserAgent     string   `json:"userAgent"`
}

// Data represents the whole Datafile
type Data struct {
	Time          Time          `json:"Time"`
	Canvas        string        `json:"Canvas"`
	Audio         string        `json:"Audio"`
	MouseCheckbox []MouseData   `json:"MouseCheckbox"`
	MouseRest     []MouseData   `json:"MouseRest"`
	Browser       BrowserInfo   `json:"Browser"`
	Screen        ScreenInfo    `json:"Screen"`
	UserID        string        `json:"UserId"`
	WebGL         WebGLInfo     `json:"WebGL"`
	FramePosition FramePosition `json:"FramePosition"`
}

// FramePosition represents the Positions of the Frames
type FramePosition struct {
	Captcha Rectangle `json:"captcha"`
	Task    Rectangle `json:"task"`
}

// MouseData represents a Mouseposition at a specific moment
type MouseData struct {
	X    uint16 `json:"x"`
	Y    uint16 `json:"y"`
	Time uint64 `json:"time"`
}

// NormalizedMouseData represents the Mouseposition relative to the Screenresolution at a specific moment
type NormalizedMouseData struct {
	X    float64
	Y    float64
	Time uint64
}

// Rectangle represents a Rectangle. Position (x|y) is the top left corner of the Rectangle.
type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// ScreenInfo contains information about the screen and the document size
type ScreenInfo struct {
	Height       uint16 `json:"height"`
	PixelDensity uint16 `json:"density"`
	Width        uint16 `json:"width"`
	InnerWidth   uint16 `json:"innerw"`
	InnerHeight  uint16 `json:"innerh"`
	DeltaX       uint16 `json:"deltax"` // Difference between screenX and clientX
	DeltaY       uint16 `json:"deltay"` // Difference between screenY and clientY
}

// Time represents Start and End time of the ReCaptcha
type Time struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

// WebGLInfo represents the Information about the Graphic Card
type WebGLInfo struct {
	Extensions []string `json:"webglExtensions"`
	GcModel    string   `json:"gcModel"`
	GcVendor   string   `json:"gcVendor"`
}
