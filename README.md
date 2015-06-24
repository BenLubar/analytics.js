# analytics
--
    import "github.com/BenLubar/analytics.js"

Package analytics provides bindings to Google Analytics's analytics.js for
gopherjs.

## Usage

#### func  Create

```go
func Create(trackingId string)
```
Create a new default tracker object. The tracking ID starts with "UA-" and
contains two numbers separated by a hyphen.

You can find your tracking ID under "Tracking Code" in the "Tracking Info" menu
at https://www.google.com/analytics/web/#management/Settings

#### func  Send

```go
func Send(hitType HitType, fieldObject map[string]interface{})
```
Send a tracking beacon to Google's collection servers. The optional field object
allows users to override one or more field values for this hit only.

#### func  Set

```go
func Set(fieldName string, value interface{})
```
Set the value currently associated with the given field.

#### func  SetMulti

```go
func SetMulti(fieldObject map[string]interface{})
```
SetMulti calls set with an object containing multiple field/value pairs for bulk
updating.

#### type HitType

```go
type HitType string
```

HitType is the type of hit to send with the Send function.

```go
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
```
