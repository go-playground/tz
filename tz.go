package tz

// Zone contains a single Country's Zone information
type Zone struct {
	CountryCode string
	Name        string
}

// Country contains a single Country's information
type Country struct {
	Code  string
	Name  string
	Zones []Zone
}
