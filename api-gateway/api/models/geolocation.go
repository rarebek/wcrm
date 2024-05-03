package models

type Geolocation struct {
	Id        int64   `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	OwnerId   string  `json:"owner_id"`
}

type GeolocationList struct {
	Geolocations []Geolocation `json:"geolocations"`
}

type CreateGeolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	OwnerId   string  `json:"owner_id"`
}

type UpdateGeolocation struct {
	Id        int64   `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	OwnerId   string  `json:"owner_id"`
}

// type CheckResponse struct {
// 	Check bool `json:"chack"`
// }
