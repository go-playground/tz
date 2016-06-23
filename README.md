Package tz
==========

![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/tz/branches/master/badge.svg)](https://semaphoreci.com/joeybloggs/tz)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/tz)](https://goreportcard.com/report/github.com/go-playground/tz)
[![GoDoc](https://godoc.org/github.com/go-playground/tz?status.svg)](https://godoc.org/github.com/go-playground/tz)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Package tz contains Timezone Country and Zone data generated from timezonedb.com


This library is nothing special, it contains alphabetically sorted Country and Zone info that can be used within your project.

#### Motivation
I got tired of rewriting this, or using a db, to store this information just so I can add to an HTML dropdown for selection. So in short here it is for use and hope it saves someone else some time as well.

#### NOTES
This is intended to be used along side Go's own time logic eg. `time.LoadLocation`

##### Example using with 
```go

countries := tz.GetCountries()

// display to user for selection
// once selected load the zones

country := tz.GetCountry(countryCode)

// display zone options to user with country.Zones
// once selected by user use it in your application

zone := country zone, selected by user, gathered from posted data, or retrieved from db ( stored on user )...

loc, _ := time.LoadLocation(zone)

// now that you have location can use with Go's time package which handles timezone offsets & Daylight savings times.

time.ParseInLocation(...)
time.Now().In(loc)

```
