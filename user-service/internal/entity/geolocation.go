package entity

type Geolocation struct {
	Id        int64
	Latitude  string
	Longitude string
	OwnerId   string
}

type AllGeolocation struct {
	Geolocations []Geolocation
	Count        int
}
