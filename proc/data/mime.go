package data

import (
	"reflect"
)

// mimetypes holds a bunch of important MIME-types
// Source https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
var mimetypes = map[string]string{
	"application/octet-stream":        "MIMEAppOctet",
	"application/pdf":                 "MIMEAppPDF",
	"application/x-google-chrome-pdf": "MIMEAppChromePDF",
	"text/plain":                      "MIMEPlain",
	"text/css":                        "MIMECSS",
	"text/html":                       "MIMEHTML",
	"text/javascript":                 "MIMEJS",
	"text/ecmacript":                  "MIMEJS",
	"application/javascript":          "MIMEJS",
	"application/ecmacript":           "MIMEJS",
	"image/apng":                      "MIMEAPNG",
	"image/bmp":                       "MIMEBMP",
	"image/gif":                       "MIMEGIF",
	"image/x-icon":                    "MIMEICO",
	"image/jpeg":                      "MIMEJPG",
	"image/png":                       "MIMEPNG",
	"image/svg+xml":                   "MIMESVG",
	"image/tiff":                      "MIMETIFF",
	"image/webp":                      "MIMEWEBP",
	"audio/wave":                      "MIMEWAV",
	"audio/wav":                       "MIMEWAV",
	"audio/x-wav":                     "MIMEWAV",
	"audio/x-pn-wav":                  "MIMEWAV",
	"audio/webm":                      "MIMEAudioWEBM",
	"video/webm":                      "MIMEVideoWEBM",
	"audio/ogg":                       "MIMEAudioOGG",
	"video/ogg":                       "MIMEVideoOGG",
	"application/ogg":                 "MIMEAppOGG",
	"multipart/form-data":             "MIMEForm",
	"multipart/byteranges":            "MIMEByte",
}

// ExtractMIMETypes extractes the MIME Types from the BrowserInfo and saves it
func (features *ProcessedFeatures) ExtractMIMETypes(binfo BrowserInfo) {
	f := reflect.ValueOf(features).Elem()

	for _, t := range binfo.MIMETypes {
		v, o := mimetypes[t]
		if o {
			field := f.FieldByName(v)
			field.SetUint(1)
		}
	}
}
