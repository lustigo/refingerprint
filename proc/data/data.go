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

// CellEvent represents a Click Event on a Cell in the Task Matrix
type CellEvent struct {
	ID       uint8  `json:"id"`
	Time     uint64 `json:"time"`
	Selected bool   `json:"selected"`
}

// ClickData represents a MouseClick Event
type ClickData struct {
	MouseData
	Key     MouseKey `json:"key"`
	Release bool     `json:"release"` // If the Button was released or pushed
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
	SessionID     string        `json:"SessionId"`
	WebGL         WebGLInfo     `json:"WebGL"`
	FramePosition FramePosition `json:"FramePosition"`
	MouseClicks   []ClickData   `json:"MouseClicks"`
	ScrollEvents  []ScrollData  `json:"ScrollEvents"`
	TaskEvents    []TaskData    `json:"TaskEvents"`
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

// MouseKey represents a Key on the Mouse
type MouseKey uint8

// MouseKeys
const (
	MouseKeyLEFT   = MouseKey(0)
	MouseKeyMIDDLE = MouseKey(1)
	MouseKeyRIGHT  = MouseKey(2)
)

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

// ScrollData represents a Wheel Event
type ScrollData struct {
	MouseData
	DeltaX    float64 `json:"deltaX"`
	DeltaY    float64 `json:"deltaY"`
	DeltaZ    float64 `json:"deltaZ"`
	DeltaMode uint8   `json:"deltaMode"`
}

// TaskType represents the Type of a ReCaptcha Task
type TaskType uint8

// TaskTypes
const (
	TaskTypeNO  = TaskType(0)
	TaskTypeSIM = TaskType(2)
	TaskTypeDYN = TaskType(3)
	TaskTypeOBJ = TaskType(4)
)

// TaskData represents the Data which is collected per Task
type TaskData struct {
	Candidate string      `json:"candidate"`
	Term      string      `json:"term"`
	Type      TaskType    `json:"type"`
	Failed    bool        `json:"failed"`
	Solved    bool        `json:"solved"`
	Events    []CellEvent `json:"events"`
	Selected  [][]bool    `json:"selected"`
	Images    []string    `json:"images"` // Data URIs of the used images
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
