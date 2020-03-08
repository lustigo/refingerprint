// Package data provides all the structs for the gathered data from the Webextension
package data

// Time represents Start and End time of the ReCaptcha
type Time struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

// Data represents the whole Datafile
type Data struct {
	Time Time `json:"Time"`
}
