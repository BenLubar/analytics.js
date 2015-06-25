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

// SetAnonymizeIP sets the anonimizeIp field. When present, the IP address of
// the sender will be anonymized.
func SetAnonymizeIP(anonymizeIP bool) {
	Set("anonmyizeIp", anonymizeIP)
}

// SetDataSource sets the dataSource field, which indicates the data source of
// the hit. Defaults to "web".
func SetDataSource(dataSource string) {
	Set("dataSource", dataSource)
}

// SetForceSSL sets the forceSSL field. By default, tracking beacons sent from
// https pages will be sent using https while beacons sent from http pages will
// be sent using http. Setting forceSSL to true will force http pages to also
// send all beacons using https. Defaults to false.
func SetForceSSL(forceSSL bool) {
	Set("forceSSL", forceSSL)
}

// Transport is the transport mechanism to use when sending hits.
type Transport string

// https://developers.google.com/analytics/devguides/collection/analyticsjs/field-reference#transport
const (
	Beacon Transport = "beacon"
	XHR    Transport = "xhr"
	Image  Transport = "image"
)

// SetTransport sets the transport field, which specifies the transport
// mechanism with which hits will be sent. The options are "beacon", "xhr", or
// "image". By default, analytics.js will try to figure out the best method
// based on the hit size and browser capabilities. If you specify "beacon" and
// the user's browser does not support the `navigator.sendBeacon` method, it
// will fall back to "image" or "xhr" depending on hit size.
func SetTransport(transport Transport) {
	Set("transport", string(transport))
}

// SetUserID sets the userId field. This is intended to be a known identifier
// for a user provided by the site owner/tracking library user. It may not
// itself be PII (personally identifiable information). The value should never
// be persisted in GA cookies or other Analytics provided storage.
func SetUserID(userID string) {
	Set("userId", userID)
}
