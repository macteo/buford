# Buford

Apple Push Notification (APN) Provider for Go 1.6 and HTTP/2.

This release is a nearly feature complete. Please see [releases](https://github.com/RobotsAndPencils/buford/releases) for updates.

[![GoDoc](https://godoc.org/github.com/RobotsAndPencils/buford?status.svg)](https://godoc.org/github.com/RobotsAndPencils/buford) [![Build Status](https://travis-ci.org/RobotsAndPencils/buford.svg?branch=ci)](https://travis-ci.org/RobotsAndPencils/buford)

### Documentation

Buford uses Apple's new HTTP/2 Notification API that was announced at WWDC 2015 and [released on December 17, 2015](https://developer.apple.com/news/?id=12172015b).

[API documentation](https://godoc.org/github.com/RobotsAndPencils/buford/) is available from GoDoc.

Also see Apple's [Local and Remote Notification Programming Guide](https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/Introduction.html), especially the sections on the JSON [payload](https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/TheNotificationPayload.html#//apple_ref/doc/uid/TP40008194-CH107-SW1) and the [Notification API](https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/APNsProviderAPI.html#//apple_ref/doc/uid/TP40008194-CH101-SW1). Also see [Safari Push Notifications](https://developer.apple.com/library/mac/documentation/NetworkingInternet/Conceptual/NotificationProgrammingGuideForWebsites/PushNotifications/PushNotifications.html#//apple_ref/doc/uid/TP40013225-CH3-SW12).

#### Terminology

**APN** Apple Push Notification

**Provider** The Buford library is used to create a _provider_ of push notifications.

**Service** Apple provides the push notification service that Buford communications with.

**Client** An `http.Client` provides an HTTP/2 client to communication with the APN Service.

**Notification** A payload sent to a device token with some headers.

**Device Token** An identifier for an application on a given device.

**Payload** The JSON sent to a device.

**Headers** HTTP/2 headers are used to for priority and expiration.

### Installation

To use this library you can install [Go 1.6 beta 1 binaries](https://groups.google.com/forum/#!topic/golang-nuts/24zV9JeBoEE) or [install Go from source](https://golang.org/doc/install/source).

Other than the standard library, Buford depends on the pkcs12 package, which can be retrieved or updated with:

```
go get -u golang.org/x/crypto/pkcs12
```

The API is not yet stable. Please use a tool like [Godep](https://github.com/tools/godep) to vendor Buford and its dependencies in your project.

### Examples

```go
package main

import (
	"log"

	"github.com/RobotsAndPencils/buford/certificate"
	"github.com/RobotsAndPencils/buford/payload"
	"github.com/RobotsAndPencils/buford/payload/badge"
	"github.com/RobotsAndPencils/buford/push"
)

func main() {
	// set these variables appropriately
	filename := "/path/to/certificate.p12"
	password := ""
	deviceToken := "c2732227a1d8021cfaf781d71fb2f908c61f5861079a00954a5453f1d0281433"

	cert, err := certificate.Load(filename, password)
	if err != nil {
		log.Fatal(err)
	}

	service := push.Service{
		Client: push.NewClient(cert),
		Host:   push.Development,
	}

	p := payload.APS{
		Alert: payload.Alert{Body: "Hello HTTP/2"},
		Badge: badge.New(42),
	}

	id, err = service.Push(deviceToken, nil, p)
	if err != nil {
		log.Fatal(err)
	}
}
```
#### Headers

You can specify an ID, expiration, priority, and other parameters via the Headers struct.

```go
headers := &push.Headers{
	ID:          "922D9F1F-B82E-B337-EDC9-DB4FC8527676",
	Expiration:  time.Now().Add(time.Hour),
	LowPriority: true,
}

id, err = service.Push(deviceToken, headers, p)
```

If no ID is specified APNS will generate and return a unique ID. When an expiration is specified, APNS will store and retry sending the notification until that time, otherwise APNS will not store or retry the notification. LowPriority should be used when sending a ContentAvailable payload.

#### Custom values

To add custom values to an APS payload, use the Map method as follows:

```go
p := payload.APS{
	Alert: payload.Alert{Body: "Message received from Bob"},
}
pm := p.Map()
pm["acme2"] = []string{"bang", "whiz"}

id, err = service.Push(deviceToken, nil, pm)
```

The Push method will use json.Marshal to serialize whatever you send it.

#### Resend the same payload

Use json.Marshal to serialize your payload once and then send it to multiple device tokens with PushBytes.

```go
b, err := json.Marshal(p)
if err != nil {
	log.Fatal(err)
}

id, err := service.PushBytes(deviceToken, nil, b)
```

Whether you use Push or PushBytes the underlying HTTP/2 connection to APNS will be reused.
