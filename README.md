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

#### func  SetAnonymizeIP

```go
func SetAnonymizeIP(anonymizeIP bool)
```
SetAnonymizeIP sets the anonimizeIp field. When present, the IP address of the
sender will be anonymized.

#### func  SetDataSource

```go
func SetDataSource(dataSource string)
```
SetDataSource sets the dataSource field, which indicates the data source of the
hit. Defaults to "web".

#### func  SetForceSSL

```go
func SetForceSSL(forceSSL bool)
```
SetForceSSL sets the forceSSL field. By default, tracking beacons sent from
https pages will be sent using https while beacons sent from http pages will be
sent using http. Setting forceSSL to true will force http pages to also send all
beacons using https. Defaults to false.

#### func  SetMulti

```go
func SetMulti(fieldObject map[string]interface{})
```
SetMulti calls set with an object containing multiple field/value pairs for bulk
updating.

#### func  SetTransport

```go
func SetTransport(transport Transport)
```
SetTransport sets the transport field, which specifies the transport mechanism
with which hits will be sent. The options are "beacon", "xhr", or "image". By
default, analytics.js will try to figure out the best method based on the hit
size and browser capabilities. If you specify "beacon" and the user's browser
does not support the `navigator.sendBeacon` method, it will fall back to "image"
or "xhr" depending on hit size.

#### func  SetUserID

```go
func SetUserID(userID string)
```
SetUserID sets the userId field. This is intended to be a known identifier for a
user provided by the site owner/tracking library user. It may not itself be PII
(personally identifiable information). The value should never be persisted in GA
cookies or other Analytics provided storage.

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

#### type Transport

```go
type Transport string
```

Transport is the transport mechanism to use when sending hits.

```go
const (
	Beacon Transport = "beacon"
	XHR    Transport = "xhr"
	Image  Transport = "image"
)
```
https://developers.google.com/analytics/devguides/collection/analyticsjs/field-reference#transport
