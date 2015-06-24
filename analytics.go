// Package analytics provides bindings to Google Analytics's analytics.js for
// gopherjs.
package analytics

import "github.com/gopherjs/gopherjs/js"

var ga = func() func(...interface{}) {
	// conversion of the code given by Google Analytics as of 2015-06-24
	js.Global.Set("GoogleAnalyticsObject", "ga")
	if js.Global.Get("ga") == js.Undefined {
		js.Global.Set("ga", js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
			if js.Global.Get("ga").Get("q") == js.Undefined {
				js.Global.Get("ga").Set("q", make(js.S, 0))
			}
			js.Global.Get("ga").Get("q").Call("push", arguments)
			return js.Undefined
		}))
	}
	js.Global.Get("ga").Set("l", js.Global.Get("Date").New().Call("valueOf"))
	script := js.Global.Get("document").Call("createElement", "script")
	anchor := js.Global.Get("document").Call("getElementsByTagName", "script").Index(0)
	script.Set("async", true)
	script.Set("src", "//www.google-analytics.com/analytics.js")
	anchor.Get("parentNode").Call("insertBefore", script, anchor)

	return func(args ...interface{}) {
		js.Global.Get("ga").Invoke(args...)
	}
}()

// Create a new default tracker object. The tracking ID starts with "UA-" and
// contains two numbers separated by a hyphen.
//
// You can find your tracking ID under "Tracking Code" in the "Tracking Info"
// menu at https://www.google.com/analytics/web/#management/Settings
func Create(trackingId string) {
	ga("create", trackingId, "auto")
}

// HitType is the type of hit to send with the Send function.
type HitType string

const (
	// https://developers.google.com/analytics/devguides/collection/analyticsjs/pages
	PageView HitType = "pageview"
	// https://developers.google.com/analytics/devguides/collection/analyticsjs/events
	Event HitType = "event"
	// https://developers.google.com/analytics/devguides/collection/analyticsjs/social-interactions
	Social HitType = "social"
	// https://developers.google.com/analytics/devguides/collection/analyticsjs/user-timings
	Timing HitType = "timing"
)

// Send a tracking beacon to Google's collection servers. The optional field
// object allows users to override one or more field values for this hit only.
func Send(hitType HitType, fieldObject map[string]interface{}) {
	if fieldObject == nil {
		ga("send", string(hitType))
	} else {
		ga("send", string(hitType), fieldObject)
	}
}

// Set the value currently associated with the given field.
func Set(fieldName string, value interface{}) {
	ga("set", fieldName, value)
}

// SetMulti calls set with an object containing multiple field/value pairs for
// bulk updating.
func SetMulti(fieldObject map[string]interface{}) {
	ga("set", fieldObject)
}
