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
	Time          Time        `json:"Time"`
	Canvas        string      `json:"Canvas"`
	Audio         string      `json:"Audio"`
	MouseCheckbox []MouseData `json:"MouseCheckbox"`
	MouseRest     []MouseData `json:"MouseRest"`
	Browser       BrowserInfo `json:"Browser"`
	Screen        ScreenInfo  `json:"Screen"`
	UserID        string      `json:"UserId"`
	WebGL         WebGLInfo   `json:"WebGL"`
}

// MouseData represents a Mouseposition at a specific moment
type MouseData struct {
	X    uint16 `json:"x"`
	Y    uint16 `json:"y"`
	Time uint64 `json:"time"`
}

// ScreenInfo contains information about the screen
type ScreenInfo struct {
	Height       uint16 `json:"height"`
	PixelDensity uint16 `json:"density"`
	Width        uint16 `json:"width"`
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
